package demo

import "database/sql"

type EntityBasic struct {
	EntityID  int64           `db:"entity_id"`
	Name      string          `db:"entity_name"`
	LegalRep  sql.NullString  `db:"legal_rep_name"`
	LegalType sql.NullInt32   `db:"legal_rep_type"`
	Province  sql.NullString  `db:"reg_province"`
	Cap       sql.NullFloat64 `db:"rec_cap"`
	Found     sql.NullString  `db:"found_date"`
	Type      sql.NullString  `db:"entity_kind"`
	Logo      sql.NullString  `db:"logo"`
}

type Company struct {
	ID     int64  `db:"company_id"`
	Name   string `db:"company_name"` // 融资过A轮的实体名
	Sector string `db:"sector"`       // 行业
	Ipo    int32  `db:"ipo"`          // not listed - 0, IPO-1, PreIPO-2
	Area   string `db:"area"`         // 所在省市
}

type Investor struct {
	// company_id, company_name, invest_id, financing_Date, financing_round, financing_amount
	ID    int64          `db:"gp_id"`
	Name  string         `db:"gp_name"`
	Brand sql.NullString `db:"gp_brand"`
}

type CompanyFin struct {
	ID        int64
	Name      string
	Sector    string
	Ipo       int32
	Area      string
	InvEvents []InvEvent
}

type InvEvent struct {
	InvID     int32
	CompID    int64
	InvDate   string
	Round     string
	Amt       string
	Investors []Investor
}
