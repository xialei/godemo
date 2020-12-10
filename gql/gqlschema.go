package gql

import (
	"github.com/graphql-go/graphql"
)

var schemaQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        graphql.DirectiveLocationQuery,
	Description: "查询函数",
	Fields: graphql.Fields{
		"hello": &graphql.Field{
			Type:        graphql.String, // 返回类型
			Description: "输出 world",     // 解释说明
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 根据查询处理函数方法进行返回对应类型的数据值
				return "world", nil
			},
		},
		"echo": &graphql.Field{
			Type:        graphql.String, // 返回类型
			Description: "参数直接输出",       // 解释说明
			Args: graphql.FieldConfigArgument{ // 参数接收
				"toEcho": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// 根据查询处理函数方法进行返回对应类型的数据值
				return "world", nil
			},
		},
		// 用户列表查询
		"users": List(),
		// 根据ID用户查询
		"user": ByID(),
	},
})

// schemaMutation 提交函数路由
var schemaMutation = graphql.NewObject(graphql.ObjectConfig{
	Name:        graphql.DirectiveLocationMutation,
	Description: "提交函数",
	Fields: graphql.Fields{
		// Create 创建新用户
		"userCreate": Create(),
		// Delete 删除用户
		"userDelete": Delete(),
		// Update 更新用户
		"userUpdate": Update(),
	},
})

// schema 解析文件对象
var schema graphql.Schema

// init 初始化，赋予schema解析文件对象
func init() {
	schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    schemaQuery,
		Mutation: schemaMutation,
	})
}

// ExecuteQuery GraphQL查询器
func ExecuteQuery(params *graphql.Params) *graphql.Result {
	params.Schema = schema
	return graphql.Do(*params)
}
