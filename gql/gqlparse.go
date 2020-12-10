package gql

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/graphql-go/graphql/language/ast"
)

// RequestOptions 查询参数结构体
type RequestOptions struct {
	Query         string                 `json:"query" url:"query" schema:"query"`
	Variables     map[string]interface{} `json:"variables" url:"variables" schema:"variables"`
	OperationName string                 `json:"operationName" url:"operationName" schema:"operationName"`
}

// parseQueryForm 解析get地址参数和post表单参数
func parseQueryForm(values url.Values) *RequestOptions {
	// 变量参数值
	var variables map[string]interface{}
	if variablesStr := values.Get("variables"); variablesStr != "" {
		variables = make(map[string]interface{})
		_ = json.Unmarshal([]byte(variablesStr), &variables)
	}
	return &RequestOptions{
		Query:         values.Get("query"),
		Variables:     variables,
		OperationName: values.Get("operationName"),
	}
}

// ParseRequestOptions 将请求参数解析为查询参数结构
func ParseRequestOptions(req *http.Request) *RequestOptions {
	// get request
	if req.Method == http.MethodGet {
		return parseQueryForm(req.URL.Query())
	}
	// post request
	switch req.Header.Get("Content-Type") {
	case "application/x-www-form-urlencoded":
		if err := req.ParseForm(); err != nil {
			return &RequestOptions{}
		}
		if reqOpt := parseQueryForm(req.PostForm); reqOpt != nil {
			return reqOpt
		}
		return &RequestOptions{}
	case "application/json":
		fallthrough
	default:
		var opts RequestOptions
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return &opts
		}
		// 反序列化成结构体
		_ = json.Unmarshal(body, &opts)
		return &opts
	}
}

// SelectionFieldNames 查询选择字段
// return 字符数组
func SelectionFieldNames(fieldASTs []*ast.Field) []string {
	fieldNames := make([]string, 0)
	for _, field := range fieldASTs {
		selections := field.SelectionSet.Selections
		for _, selection := range selections {
			fieldNames = append(fieldNames, selection.(*ast.Field).Name.Value)
		}
	}
	return fieldNames
}
