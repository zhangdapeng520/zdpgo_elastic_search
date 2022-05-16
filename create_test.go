package zdpgo_elastic_search

import (
	"fmt"
	"testing"
)

/*
@Time : 2022/5/12 23:46
@Author : 张大鹏
@File : create_test
@Software: Goland2021.3.1
@Description: create创建相关的测试
*/

func TestElasticSearch_CreateIndex(t *testing.T) {
	e := getElasticSearch()
	flag := e.CreateIndex("test", Index{
		Mappings: Mappings{
			Properties: Properties{
				User: Property{
					Type: "keyword",
				},
				Message: Property{
					Type: "text",
				},
				Image: Property{
					Type: "keyword",
				},
				Created: Property{
					Type: "date",
				},
				Tags: Property{
					Type: "keyword",
				},
				Location: Property{
					Type: "geo_point",
				},
				SuggestField: Property{
					Type: "completion",
				},
			},
		},
	})
	if flag {
		fmt.Println("创建索引成功")
	} else {
		fmt.Println("创建索引失败")
	}
}
