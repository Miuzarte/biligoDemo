package main

import (
	"fmt"

	"github.com/Miuzarte/biligo"
)

func main() {
	const keyword = "陈睿"

	// SearchFormatAuto 在 searchType 为空字符串时使用综合搜索,
	// 返回值可被断言为:
	// [*biligo.VideoInfo], [*biligo.Media]
	// [*biligo.LiveStatus], [*biligo.ArticleInfo],
	// [*biligo.SpaceCard], [*biligo.Error]
	// 为什么这么复杂, 因为还在用 gjson 来解析搜索结果,
	// 获取了对应的具体信息再返回解析好的结构体
	results, err := biligo.SearchFormatAuto("", keyword)
	// results, err := biligo.SearchFormatAll(keyword)
	if err != nil {
		panic(err)
	}

	for _, result := range results {
		switch result := result.(type) {
		case *biligo.VideoInfo:
			fmt.Println("VideoInfo:", result.Title)
		case *biligo.Media:
			fmt.Println("Media:", result.Title)
		case *biligo.LiveStatus:
			fmt.Println("LiveStatus:", result.Title)
		case *biligo.ArticleInfo:
			fmt.Println("ArticleInfo:", result.Title)
		case *biligo.SpaceCard:
			fmt.Println("SpaceCard:", result.Card.Name)

		case *biligo.Error:
			fmt.Println("Error:", result.Error())
		}
	}
}
