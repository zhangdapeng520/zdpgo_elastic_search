package zdpgo_elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/zhangdapeng520/zdpgo_elasticsearch/elasticsearch"
	"github.com/zhangdapeng520/zdpgo_elasticsearch/elasticsearch/esapi"
	"github.com/zhangdapeng520/zdpgo_json"
	"github.com/zhangdapeng520/zdpgo_log"
)

// ElasticSearch ES核心对象
type ElasticSearch struct {
	Config *Config               // 配置对象
	Log    *zdpgo_log.Log        // 日志对象
	Json   *zdpgo_json.Json      // Json对象
	Client *elasticsearch.Client // ES客户端对象
}

func New() (*ElasticSearch, error) {
	return NewWithConfig(&Config{})
}

func NewWithConfig(config *Config) (*ElasticSearch, error) {
	e := &ElasticSearch{}

	// 日志
	if config.LogFilePath == "" {
		config.LogFilePath = "logs/zdpgo/zdpgo_elastic_search.log"
	}
	logConfig := zdpgo_log.Config{
		Debug:       config.Debug,
		OpenJsonLog: true,
		LogFilePath: config.LogFilePath,
	}
	if config.Debug {
		logConfig.IsShowConsole = true
	}
	e.Log = zdpgo_log.NewWithConfig(logConfig)

	// ES客户端
	if config.Addresses == nil || len(config.Addresses) == 0 {
		config.Addresses = []string{"https://127.0.0.1:9200"}
	}
	if config.Username == "" {
		config.Username = "elastic"
	}
	if config.Password == "" {
		config.Password = "elastic"
	}
	if config.CertPath == "" {
		return nil, fmt.Errorf("cert文件路径不能为空")
	}
	if config.HealthInterval == 0 {
		config.HealthInterval = 3
	}
	if config.RetryTimeNum == 0 {
		config.RetryTimeNum = 3
	}
	// 创建Client, 连接ES
	cert, err := ioutil.ReadFile("http_ca.crt")
	if err != nil {
		return nil, err
	}
	cfg := elasticsearch.Config{
		Addresses: config.Addresses,
		Username:  config.Username,
		Password:  config.Password,
		CACert:    cert,
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	e.Client = client

	// 配置对象
	e.Config = config

	// 其他对象
	e.Json = zdpgo_json.New()

	return e, nil
}

// Version 获取ES的版本号
func (e *ElasticSearch) Version() string {
	res, err := e.Client.Info()
	if err != nil {
		e.Log.Error("获取ES服务器信息失败", "error", err)
		return ""
	}
	defer res.Body.Close()

	// 检查响应状态
	if res.IsError() {
		e.Log.Error("获取ES服务器信息失败", "error", res.String())
		return ""
	}

	// 解析想要为一个map
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		e.Log.Error("解析相应数据失败", "error", err)
		return ""
	}

	defer func() {
		if r1 := recover(); r1 != nil {
			e.Log.Error("获取ES服务器版本失败", "error", r1)
		}
	}()
	return r["version"].(map[string]interface{})["number"].(string)
}

type Response struct {
	Status  string      `json:"status"`
	Result  interface{} `json:"result"`
	Version int         `json:"version"`
}

// 添加数据
func (e *ElasticSearch) Add(index, documentId string, jsonData interface{}) (*Response, error) {
	// 构建要添加的数据
	data, err := json.Marshal(jsonData)
	if err != nil {
		e.Log.Error("序列化JSON数据失败，请检查数据是否符合JSON规范", "error", err, "data", jsonData)
		return nil, err
	}

	// 构建请求对象
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: documentId,
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}

	// 发送请求
	res, err := req.Do(context.Background(), e.Client)
	if err != nil {
		e.Log.Error("发送ES请求失败", "error", err, "req", req)
		return nil, err
	}
	defer res.Body.Close()

	// 解析相应数据
	if res.IsError() {
		e.Log.Error("获取响应数据失败", "error", res.Status())
		return nil, errors.New("获取响应数据失败")
	} else {
		// 解析数据为一个MAP对象
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			e.Log.Error("解析数据为map对象失败", "error", err)
			return nil, err
		}
		return &Response{
			Status:  res.Status(),
			Result:  r["result"],
			Version: int(r["_version"].(float64)),
		}, nil
	}
}

type Request struct {
	QueryData QueryData `json:"query_data"`
	Index     string    `json:"index"`
}

type QueryData struct {
	Query Query `json:"query"`
}

type Query struct {
	Match Match `json:"match"`
}

type Match struct {
	Username string `json:"username"`
}

func (e *ElasticSearch) Search(request *Request) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request.QueryData); err != nil {
		e.Log.Error("解析查询数据失败", "error", err)
		return err
	}

	// 发送搜索请求
	es := e.Client
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(request.Index),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	// 解析响应数据
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	log.Println(strings.Repeat("=", 37))
	return nil
}
