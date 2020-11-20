package gql

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "用户字段",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.ID,
			Description: "ID值",
		},
		"username": &graphql.Field{
			Type:        graphql.String,
			Description: "用户名",
		},
		"password": &graphql.Field{
			Type:        graphql.String,
			Description: "密码",
		},
		"info": &graphql.Field{
			Type:        userInfo,
			Description: "关联用户信息",
			Resolve:     info, // 级联函数调用
		},
	},
})

var userInfo = graphql.NewObject(graphql.ObjectConfig{
	Name:        "UserInfo",
	Description: "用户信息字段",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.ID,
			Description: "ID值",
		},
		"age": &graphql.Field{
			Type:        graphql.Int,
			Description: "年龄",
		},
		"firstName": &graphql.Field{
			Type:        graphql.String,
			Description: "名",
		},
		"lastName": &graphql.Field{
			Type:        graphql.String,
			Description: "姓",
		},
	},
})

// lastNameEnum 姓-枚举类型
var lastNameEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "LastNameEnum",
	Description: "名选择",
	Values: graphql.EnumValueConfigMap{
		"OK": &graphql.EnumValueConfig{
			Value:       graphql.String,
			Description: "好的",
		},
		"SOLO": &graphql.EnumValueConfig{
			Value:       graphql.String,
			Description: "solo",
		},
	},
})

// createInput  用户创建修改-Mutation提交输入类型
var createInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserCreateInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"username": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "用户名",
		},
		"password": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "密码",
		},
		"info": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(infoCreateInput),
			Description: "关联用户信息",
		},
	},
})

// infoCreateInput 部分用户信息创建修改-Mutation提交输入类型
var infoCreateInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserInfoCreateInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"age": &graphql.InputObjectFieldConfig{
			Type:        graphql.Int,
			Description: "年龄",
		},
		"firstName": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: "名",
		},
		"lastName": &graphql.InputObjectFieldConfig{
			Type:        lastNameEnum,
			Description: "姓",
		},
	},
})

//List 用户列表查询
func List() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(queryType),
		Description: "用户列表查询",
		Args: graphql.FieldConfigArgument{
			"page": &graphql.ArgumentConfig{
				Type:         graphql.Int,
				DefaultValue: 1,
			},
			"query": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			log.Printf("List query Args %v \n", params.Args)
			log.Printf("List query VariableValues %v \n", params.Info.VariableValues)

			log.Printf("FieldASTs %v \n", SelectionFieldNames(params.Info.FieldASTs))
			return testUserList, nil
		},
	}
}

//ByID 根据ID用户查询
func ByID() *graphql.Field {
	return &graphql.Field{
		Type:        queryType,
		Description: "根据ID用户查询",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			// 查找对应的id，SQL...
			userID := params.Args["id"]
			for _, v := range testUserList {
				if v["id"] == userID {
					return v, nil
				}
			}
			return nil, nil
		},
	}
}

// Create 创建新用户
func Create() *graphql.Field {
	return &graphql.Field{
		Type:        queryType,
		Description: "创建新用户",
		Args: graphql.FieldConfigArgument{
			"arg": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(createInput),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			arg := params.Args["arg"].(map[string]interface{})
			// 随机种子
			rand.Seed(time.Now().UnixNano())
			id := strconv.Itoa(rand.Intn(1000))
			// 用户添加
			user := map[string]interface{}{
				"id":       "user-" + id,
				"username": arg["username"],
				"password": arg["password"],
				"info":     "info-" + id,
			}
			testUserList = append(testUserList, user)
			// 用户信息添加
			info := arg["info"].(map[string]interface{})
			userInfo := map[string]interface{}{
				"id":        "info-" + id,
				"age":       info["age"],
				"firstName": info["firstName"],
				"lastName":  info["lastName"],
			}
			testUserInfoList = append(testUserInfoList, userInfo)
			return testUserList[len(testUserList)-1], nil
		},
	}
}

// Delete 删除用户
func Delete() *graphql.Field {
	return &graphql.Field{
		Type:        queryType,
		Description: "删除用户",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id := params.Args["id"].(string)
			user := make(map[string]interface{})
			for i, p := range testUserList {
				if id == p["id"] {
					user = p
					testUserList = append(testUserList[:i], testUserList[i+1:]...)
					testUserInfoList = append(testUserInfoList[:i], testUserInfoList[i+1:]...)
				}
			}
			return user, nil
		},
	}
}

// Update 更新用户
func Update() *graphql.Field {
	return &graphql.Field{
		Type:        queryType,
		Description: "更新用户",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"arg": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(createInput),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id := params.Args["id"].(string)
			arg := params.Args["arg"].(map[string]interface{})

			user := make(map[string]interface{})
			for _, p := range testUserList {
				if id == p["id"] {
					user = p
					// 随机种子
					rand.Seed(time.Now().UnixNano())
					rid := strconv.Itoa(rand.Intn(1000))
					user["password"] = arg["password"].(string) + rid
					user["username"] = arg["username"].(string) + rid
					break
				}
			}

			return user, nil
		},
	}
}

// info 级联的信息函数处理
func info(p graphql.ResolveParams) (interface{}, error) {
	log.Printf("info级联的信息函数处理来自上层Source %v \n", p.Source)
	// 查找对应的id，SQL...
	infoID := p.Source.(map[string]interface{})["info"]
	for _, v := range testUserInfoList {
		if v["id"] == infoID {
			return v, nil
		}
	}
	return nil, nil
}
