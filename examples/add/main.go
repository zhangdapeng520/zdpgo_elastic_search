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

	// 添加数据
	data := map[string]string{
		"username": "elastic",
		"password": "2ANkC+Yo",
	}
	response, err := e.Add("test", "1", data)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
