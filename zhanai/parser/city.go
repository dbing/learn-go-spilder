package parser

import (
	"imooc/learngo/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(cityRe)
	matches := compile.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseUserProfile(c, string(m[2]))
			},
		})
	}
	return result
}
