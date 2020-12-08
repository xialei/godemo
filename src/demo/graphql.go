package demo

import (
	"encoding/json"
	"gql"
	"log"
	"net/http"
	"text/template"

	"github.com/graphql-go/graphql"
)

// StartGraphQL start server
func StartGraphQL() (string, error) {
	// home page
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-Type", "text/html; charset=utf-8")
		res.WriteHeader(http.StatusOK)
		// template
		t, err := template.ParseFiles("../index.html")
		if err != nil {
			log.Printf("res:%v \n err:%v \n", res, err)
			return
		}
		err = t.Execute(res, "GraphiQL")
	})

	// request entry
	http.HandleFunc("/graphql", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-Type", "application/json; charset=utf-8")
		res.WriteHeader(http.StatusOK)
		opts := gql.ParseRequestOptions(req)
		// query
		result := gql.ExecuteQuery(&graphql.Params{
			RequestString:  opts.Query,
			VariableValues: opts.Variables,
			OperationName:  opts.OperationName,
			Context:        req.Context(),
		})
		if len(result.Errors) > 0 {
			log.Printf("errors: %v", result.Errors)
		}
		// map to json
		buff, _ := json.Marshal(result)
		_, _ = res.Write(buff)
	})

	log.Printf("Server is running on 127.0.0.1:8080 \n")
	_ = http.ListenAndServe("127.0.0.1:8080", nil)

	return "exit", nil

}
