package zdpgo_elastic_search

/*
@Time : 2022/5/12 23:11
@Author : 张大鹏
@File : elastic_search
@Software: Goland2021.3.1
@Description: ES核心相关
*/
import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/zhangdapeng520/zdpgo_json"
	"github.com/zhangdapeng520/zdpgo_log"
	"log"
	"os"
	"time"
)

// ElasticSearch ES核心对象
type ElasticSearch struct {
	Config *Config          // 配置对象
	Log    *zdpgo_log.Log   // 日志对象
	Json   *zdpgo_json.Json // Json对象
	Client *elastic.Client  // ES客户端对象
}

func New() *ElasticSearch {
	return NewWithConfig(Config{})
}

func NewWithConfig(config Config) *ElasticSearch {
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
	if config.Host == "" {
		config.Host = "127.0.0.1"
	}
	if config.Port == 0 {
		config.Port = 9200
	}
	if config.Username == "" {
		config.Username = "elastic"
	}
	if config.Password == "" {
		config.Password = "elastic"
	}
	if config.HealthInterval == 0 {
		config.HealthInterval = 3
	}
	if config.RetryTimeNum == 0 {
		config.RetryTimeNum = 3
	}
	// 创建Client, 连接ES
	setZip := !config.IsCloseZip
	client, err := elastic.NewClient(
		// 关闭sniff模式，否则无法连接Docker中的ES
		elastic.SetSniff(config.IsOpenSniff),
		// elasticsearch 服务地址，多个服务地址使用逗号分隔
		elastic.SetURL(fmt.Sprintf("http://%s:%d", config.Host, config.Port)),
		// 基于http base auth验证机制的账号和密码
		elastic.SetBasicAuth(config.Username, config.Password),
		// 启用gzip压缩
		elastic.SetGzip(setZip),
		// 设置监控检查时间间隔
		elastic.SetHealthcheckInterval(time.Duration(config.HealthInterval)*time.Second),
		// 设置请求失败最大重试次数
		elastic.SetMaxRetries(config.RetryTimeNum),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	if err != nil {
		e.Log.Error("连接ES失败", "error", err, "config", config)
	} else {
		e.Log.Debug("连接ES成功")
	}
	e.Client = client

	// 配置对象
	e.Config = &config

	// 其他对象
	e.Json = zdpgo_json.New()

	return e
}

// Version 获取ES的版本号
func (e *ElasticSearch) Version() string {
	address := fmt.Sprintf("http://%s:%d", e.Config.Host, e.Config.Port)
	version, err := e.Client.ElasticsearchVersion(address)
	if err != nil {
		e.Log.Error("获取ES版本号失败", "error", err, "address", address)
	}
	return version
}
