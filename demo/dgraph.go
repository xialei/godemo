package demo

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"godemo/core"
	"log"
	"strconv"
	"strings"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
)

var (
	dgraph            = flag.String("d", "localhost:9080", "Dgraph Alpha address")
	company_map       = make(map[int64]Company)
	investor_map      = make(map[int64]Investor)
	event_map         = make(map[int32]InvEvent)
	company_fin_map   = make(map[int64]CompanyFin)
	investor_uids_map = make(map[int64]string)
)

func buildDgraph() {
	fmt.Println("== start to build dgraph")

	flag.Parse()
	conn, err := grpc.Dial(*dgraph, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	ctx := context.Background()
	err = client.Alter(ctx, &api.Operation{DropAll: true})
	if err != nil {
		log.Fatal(err)
	}

	op := &api.Operation{
		Schema:          `name: string @index(term) .`,
		RunInBackground: true,
	}
	err = client.Alter(ctx, op)
	if err != nil {
		fmt.Printf("alter schema err: %d\n", err)
	}

	mu := &api.Mutation{CommitNow: true}
	for _, v := range investor_map {
		pb, err1 := json.Marshal(v)
		if err1 != nil {
			fmt.Printf("commit mutation err: %d\n", err1)
		}
		mu.SetJson = pb
		resp, err := client.NewTxn().Mutate(ctx, mu)
		uid_map := resp.GetUids()
		for _, inv_uid := range uid_map {
			investor_uids_map[v.ID] = inv_uid
		}

		if err != nil {
			fmt.Printf("commit mutation err: %d\n", err)
		}
	}
	fmt.Printf("investor_uids: %d\n", len(investor_uids_map))

	for _, v := range company_fin_map {

		eventlist := v.InvEvents
		for eindx, event := range eventlist {
			investor_list := event.Investors

			for i, v := range investor_list {
				investor_list[i].uid = investor_uids_map[v.ID]
			}
			eventlist[eindx].Investors = investor_list
		}
		v.InvEvents = eventlist

		pb, err1 := json.Marshal(v)
		if err1 != nil {
			fmt.Printf("commit mutation err: %d\n", err1)
		}
		mu.SetJson = pb
		_, err := client.NewTxn().Mutate(ctx, mu)
		if err != nil {
			fmt.Printf("commit mutation err: %d\n", err)
		}
	}
}

func prepareData() {
	dbConn, err := core.GetDBConnection()

	sql := `select par_code, par_name from base_pub_par where par_sys_code=100022`

	rows, err := dbConn.Query(sql)

	if err != nil {
		fmt.Printf("query gp data err: %d\n", err)
	}

	round_map := make(map[int64]string)
	for rows.Next() {
		var code int64
		var name string
		err := rows.Scan(&code, &name)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
		}
		round_map[code] = name
	}

	fmt.Println("1. start to get investors")

	sql = `select gp_id, gp_name_cn as gp_name, initial_cn as gp_brand from pe_gp_basic_info where gp_id in (select distinct b.gp_id from pe_company_financing  a 
	left join pe_company_financing_investor b on a.invest_id = b.invest_id
	where a.financing_round not in (0, 8, 13, 18, 28, 30, 32, 34, 36, 41, 47) and b.gp_id is not null )`

	rows, err = dbConn.Query(sql)

	if err != nil {
		fmt.Printf("query gp data err: %d\n", err)
	}

	for rows.Next() {
		var investor Investor
		err := rows.Scan(&investor.ID, &investor.Name, &investor.Brand)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
		}
		investor_map[investor.ID] = investor
	}
	fmt.Printf("investor_map: %d\n", len(investor_map))

	fmt.Println("2. start to get companies")

	sql = `select a.company_id, a.company_name, b.name as sector, 
	case when c.is_preipo = 1 then 2 when c.is_ipo = 0 then 0 when c.is_ipo = 1 then 1  else 0 END as ipo, 
	a.location_province as area from pe_company_basic_info a 
	left join base_industry as b on a.sector=b.code 
	left join pe_company_filing_status as c on a.company_id = c.entity_id 
	where a.company_id in (select distinct a.company_id from pe_company_financing  a 
		left join pe_company_financing_investor b on a.invest_id = b.invest_id
		where a.financing_round not in (0, 8, 13, 18, 28, 30, 32, 34, 36, 41, 47) and b.gp_id is not null ) 
		and a.location_country='CN' and a.location_province !=''`

	rows, err = dbConn.Query(sql)

	if err != nil {
		fmt.Printf("query gp data err: %d\n", err)
	}

	for rows.Next() {
		var company Company
		err := rows.Scan(&company.ID, &company.Name, &company.Sector, &company.Ipo, &company.Area)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
		}
		company_map[company.ID] = company
	}
	fmt.Printf("company_map: %d\n", len(company_map))

	fmt.Println("3. start to get events")

	sql = `select a.invest_id, a.company_id, a.financing_Date, a.financing_round, a.financing_amount, 
	GROUP_CONCAT(b.gp_id) as investors from pe_company_financing  a 
	left join pe_company_financing_investor b on a.invest_id = b.invest_id
	left join pe_company_basic_info c on a.company_id = c.company_id and c.location_country='CN' and c.location_province !=''
	where a.financing_round not in (0, 8, 13, 18, 28, 30, 32, 34, 36, 41, 47) and b.gp_id is not null GROUP BY a.invest_id`

	rows, err = dbConn.Query(sql)

	if err != nil {
		fmt.Printf("query events data err: %d\n", err)
	}

	for rows.Next() {
		var invid int32
		var compid int64
		var invdate, invRound, amt, investors string
		err := rows.Scan(&invid, &compid, &invdate, &invRound, &amt, &investors)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
		}
		strList := strings.Split(investors, ",")
		var investorList []Investor
		for _, item := range strList {
			val, _ := strconv.ParseInt(item, 10, 64)
			investorList = append(investorList, investor_map[int64(val)])
		}
		newRound, _ := strconv.ParseInt(invRound, 10, 32)
		e := InvEvent{invid, compid, invdate, round_map[newRound], amt, investorList}
		event_map[invid] = e

		var compFin, ok = company_fin_map[compid]
		if !ok {
			var comp = company_map[compid]
			var ev = []InvEvent{e}
			compFin := CompanyFin{comp.ID, comp.Name, comp.Sector, comp.Ipo, comp.Area, ev}
			company_fin_map[compid] = compFin
		} else {
			compFin.InvEvents = append(compFin.InvEvents, e)
			company_fin_map[compid] = compFin
		}

	}

	fmt.Printf("companyfin_map: %d\n", len(company_fin_map))

}

func TestDgraph() {
	prepareData()
	buildDgraph()
}
