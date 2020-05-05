package main

import (
	"imooc/learngo/crawler/engine"
	"imooc/learngo/crawler/zhanai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
