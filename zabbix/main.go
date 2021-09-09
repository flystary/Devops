package main

type Interface struct {
	Type  int    `json:"type"`
	Main  int    `json:"main"`
	Useip int    `json:"useip"`
	Ip    string `json:"ip"`
	Dns   string `json:"dns"`
	Port  string `json:"port"`
}

// Groups 数组
type Groups struct {
	Groupid string `json:"groupid"`
}

// Template 模板数组
type Template struct {
	Templateid string `json:"templateid"`
}

type Params struct {
	Host       string        `json:"host"`
	Interfaces []interface{} `json:"interfaces"`
	Groups     []interface{} `json:"groups"`
	Templates  []interface{} `json:"templates"`
}

type Data struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
	Auth    string `json:"auth"`
	Id      int    `json:"id"`
}
