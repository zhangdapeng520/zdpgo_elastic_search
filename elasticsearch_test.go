package zdpgo_elasticsearch

import (
	"testing"
)

// 测试获取版本号
func TestElasticSearch_Version(t *testing.T) {
	e, err := NewWithConfig(&Config{
		Addresses: []string{"http://localhost:9200"},
		Username:  "elastic",
		Password:  "2ANkC+YoJab*7NnK2fgN",
		CertPath:  "http_ca.crt",
	})
	if err != nil {
		panic(err)
	}
	if e.Version() == "" {
		panic("version error")
	}
}
