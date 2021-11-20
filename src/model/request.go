package model

// Request Url用于fetch 解析成字节切片
//ParseFunction 用于将字节切片 解析成 ParseResult
type Request struct {
	Url           string
	ParseFunction func([]byte) ParseResult
}

// ParseResult 建字节解析成 Requests和Items
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParseFunction([]byte) ParseResult {
	return ParseResult{}
}
