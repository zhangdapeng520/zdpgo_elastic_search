package zdpgo_elastic_search

import "context"

/*
@Time : 2022/5/16 15:20
@Author : 张大鹏
@File : is_bool.go
@Software: Goland2021.3.1
@Description: is类型的判断方法
*/

func (e *ElasticSearch) IsExistsIndex(indexName string) bool {
	// 检测下weibo索引是否存在
	exists, err := e.Client.IndexExists(indexName).Do(context.Background())
	if err != nil {
		e.Log.Error("检测索引是否存在失败", "error", err, "indexName", indexName)
	}
	return exists
}
