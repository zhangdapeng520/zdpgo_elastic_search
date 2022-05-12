package zdpgo_elastic_search

import "testing"

/*
@Time : 2022/5/12 23:33
@Author : 张大鹏
@File : elastic_search_test
@Software: Goland2021.3.1
@Description: 核心对象相关测试
*/
func getElasticSearch() *ElasticSearch {
	return NewWithConfig(Config{
		Debug:    true,
		Host:     "192.168.32.128",
		Port:     9200,
		Username: "elastic",
		Password: "elastic",
	})
}

func TestElasticSearch_NewWithConfig(t *testing.T) {
	e := getElasticSearch()
	if e == nil {
		panic("获取ES对象失败")
	}
}