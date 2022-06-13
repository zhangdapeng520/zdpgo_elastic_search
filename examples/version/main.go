package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_elasticsearch"
)

func main() {
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
}
