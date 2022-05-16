package zdpgo_elastic_search

import "context"

/*
@Time : 2022/5/16 14:58
@Author : 张大鹏
@File : delete.go
@Software: Goland2021.3.1
@Description: delete 删除相关方法
*/

// DeleteIndex 删除索引，支持同时删除多个
func (e *ElasticSearch) DeleteIndex(indexNameList ...string) bool {
	_, err := e.Client.DeleteIndex(indexNameList...).Do(context.Background())
	if err != nil {
		e.Log.Error("删除索引失败", "error", err, "indexNameList", indexNameList)
		return false
	}
	return true
}
