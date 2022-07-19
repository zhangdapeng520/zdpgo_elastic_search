package main

import (
	"context"
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

	// 创建创建一条微博
	msg1 := Weibo{User: "olivere", Message: "打酱油的一天", Retweets: 0}

	// 使用client创建一个新的文档
	put1, err := client.Index().
		Index("weibo"). // 设置索引名称
		Id("1").        // 设置文档id
		BodyJson(msg1). // 指定前面声明的微博内容
		Do(ctx)         // 执行请求，需要传入一个上下文对象
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("文档Id %s, 索引名 %s\n", put1.Id, put1.Index)
}
