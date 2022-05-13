package zdpgo_elastic_search

import "testing"

/*
@Time : 2022/5/12 23:46
@Author : 张大鹏
@File : create_test
@Software: Goland2021.3.1
@Description: create创建相关的测试
*/

func TestElasticSearch_Create(t *testing.T) {
	e := getElasticSearch()
	e.Create()
}

func TestElasticSearch_CreateIndex(t *testing.T) {
	e := getElasticSearch()
	e.CreateIndex("")
}
