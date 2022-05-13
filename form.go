package zdpgo_elastic_search

/*
@Time : 2022/5/12 23:56
@Author : 张大鹏
@File : form
@Software: Goland2021.3.1
@Description: form表单相关的静态结构体
*/

// Index 索引对象
type Index struct {
	Mappings Mappings `json:"mappings"`
}

type Mappings struct {
	Properties Properties `json:"properties"`
}

type Properties struct {
	User         Property `json:"user"`
	Message      Property `json:"message"`
	Image        Property `json:"image"`
	Created      Property `json:"created"`
	Tags         Property `json:"tags"`
	Location     Property `json:"location"`
	SuggestField Property `json:"suggest_field"`
}

type Property struct {
	Type string `json:"type"`
}
