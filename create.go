package zdpgo_elastic_search

import (
	"context"
	"fmt"
)

/*
@Time : 2022/5/12 23:39
@Author : 张大鹏
@File : create
@Software: Goland2021.3.1
@Description: create创建相关
*/

//创建
func (e *ElasticSearch) Create() {
	//使用字符串
	e2 := `{"first_name":"John","last_name":"Smith","age":25,"about":"I love to go rock climbing","interests":["sports","music"]}`
	put2, err := e.Client.Index().
		Index("megacorp").
		Type("employee").
		Id("2").
		BodyJson(e2).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put2.Id, put2.Index, put2.Type)
}

// CreateIndex 创建索引
func (e *ElasticSearch) CreateIndex(indexName string, index Index) bool {

	// 执行ES请求需要提供一个上下文对象
	ctx := context.Background()

	// 索引mapping定义，这里仿微博消息结构定义
	indexJson, err := e.Json.Dumps(index)
	if err != nil {
		e.Log.Error("序列化索引对象失败", "error", err, "index", index)
		return false
	}
	// 创建索引
	_, err = e.Client.CreateIndex(indexName).BodyString(indexJson).Do(ctx)
	if err != nil {
		e.Log.Error("创建索引失败", "error", err)
		return false
	}

	return true
}
