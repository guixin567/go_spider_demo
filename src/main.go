package main

import (
	"zhengxin2/src/engine"
	"zhengxin2/src/model"
	"zhengxin2/src/parser"
)

func main() {
	engine.Run(model.Request{
		Url:           "https://www.zhenai.com/zhenghun",
		ParseFunction: parser.CityListParse,
	})
}
