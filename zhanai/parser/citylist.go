package parser

import (
	"imooc/learngo/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" [^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(cityListRe)
	matchs := compile.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	//limit := 3
	for _, m := range matchs {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		//limit--
		//if limit == 0 {
		//	break
		//}
	}
	return result
}
