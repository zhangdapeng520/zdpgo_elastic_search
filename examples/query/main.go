package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

// 索引mapping定义，这里仿微博消息结构定义
const mapping = `
{
  "mappings": {
    "properties": {
      "user": {
        "type": "keyword"
      },
      "message": {
        "type": "text"
      },
      "image": {
        "type": "keyword"
      },
      "created": {
        "type": "date"
      },
      "tags": {
        "type": "keyword"
      },
      "location": {
        "type": "geo_point"
      },
      "suggest_field": {
        "type": "completion"
      }
    }
  }
}`

type Weibo struct {
	User     string                `json:"user"`               // 用户
	Message  string                `json:"message"`            // 微博内容
	Retweets int                   `json:"retweets"`           // 转发数
	Image    string                `json:"image,omitempty"`    // 图片
	Created  time.Time             `json:"created,omitempty"`  // 创建时间
	Tags     []string              `json:"tags,omitempty"`     // 标签
	Location string                `json:"location,omitempty"` //位置
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

func main() {
	// 创建client
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200", "http://127.0.0.1:9201"),
		elastic.SetBasicAuth("user", "secret"))
	if err != nil {
		// Handle error
		fmt.Printf("连接失败: %v\n", err)
	} else {
		fmt.Println("连接成功")
	}

	// 执行ES请求需要提供一个上下文对象
	ctx := context.Background()

	// 根据id查询文档
	get1, err := client.Get().
		Index("weibo"). // 指定索引名
		Id("1").        // 设置文档id
		Do(ctx)         // 执行请求
	if err != nil {
		// Handle error
		panic(err)
	}
	if get1.Found {
		fmt.Printf("文档id=%s 版本号=%d 索引名=%s\n", get1.Id, get1.Version, get1.Index)
	}

	//# 手动将文档内容转换成go struct对象
	msg2 := Weibo{}
	// 提取文档内容，原始类型是json数据
	data, _ := get1.Source.MarshalJSON()
	// 将json转成struct结果
	json.Unmarshal(data, &msg2)
	// 打印结果
	fmt.Println(msg2.Message)
}
