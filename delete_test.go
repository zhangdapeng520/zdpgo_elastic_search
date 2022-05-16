package zdpgo_elastic_search

import (
	"fmt"
	"testing"
)

/*
@Time : 2022/5/16 15:00
@Author : 张大鹏
@File : delete_test.go
@Software: Goland2021.3.1
@Description: delete删除相关的方法测试
*/

func TestElasticSearch_DeleteIndex(t *testing.T) {
	e := getElasticSearch()
	flag := e.DeleteIndex("test")
	if flag {
		fmt.Println("删除索引成功")
	} else {
		fmt.Println("删除索引失败")
	}
}
