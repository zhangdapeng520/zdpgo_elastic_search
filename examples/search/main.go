package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_elasticsearch"
)

func main() {
	// 创建ES对象
	e, err := zdpgo_elasticsearch.NewWithConfig(&zdpgo_elasticsearch.Config{
		Debug:     true,
		Addresses: []string{"https://localhost:9200"},
		Username:  "elastic",
		Password:  "2ANkC+YoJab*7NnK2fgN",
		CertPath:  "http_ca.crt",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(e.Version())

	// 搜索数据
	request := &zdpgo_elasticsearch.Request{
		QueryData: zdpgo_elasticsearch.QueryData{
			Query: zdpgo_elasticsearch.Query{
				Match: zdpgo_elasticsearch.Match{
					Username: "elastic",
				},
			},
		},
		Index: "test",
	}
	err = e.Search(request)
	if err != nil {
		panic(err)
	}
}
