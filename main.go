package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

func main() {
	engine.SimpleEngine{}.Run(engine.Request{
		URL:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
