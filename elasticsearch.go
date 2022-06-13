package zdpgo_elasticsearch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/zhangdapeng520/zdpgo_elasticsearch/elasticsearch"
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

	// Check response status
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
