package main

import (
	"imooc/learngo/crawler/engine"
	"imooc/learngo/crawler/scheduler"
	"imooc/learngo/crawler/zhanai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		//Scheduler:   &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
