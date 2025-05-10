package main

import (
	"encoding/base64"
	"fmt"
	"strings"
	"text/template"

	"github.com/Miuzarte/biligo"
)

func main() {
	const myTemplate = //
	`{{define "SpaceCard"}}{{ToBase64 .Card.Face}}
{{.Card.Name}} (LV{{.Card.LevelInfo.CurrentLevel}})
space.bilibili.com/{{.Card.Mid}}{{end}}`

	sc, err := biligo.FetchSpaceCard("59442895")
	if err != nil {
		panic(err)
	}
	fmt.Println("before:")
	fmt.Println(sc.DoTemplate())
	/*
		before:
		[CQ:image,file=https://i1.hdslb.com/bfs/face/2cb14ef00fc794fa9a9fed9c403113a83a63f345.jpg]
		謬紗特（LV6）
		签名：方舟B服：謬紗特 全图鉴
		space.bilibili.com/59442895
	*/

	biligo.SetTemplateFor[biligo.VideoInfo](biligo.TemplateConfig{
		Template: myTemplate,
		Funcs: template.FuncMap{
			"ToBase64": func(s string) string {
				sb := strings.Builder{}
				base64.NewEncoder(base64.StdEncoding, &sb).Write([]byte(s))
				return sb.String()
			},
		},
	})
	fmt.Println("after:")
	fmt.Println(sc.DoTemplate())
	/*
		after:
		aHR0cHM6Ly9pMS5oZHNsYi5jb20vYmZzL2ZhY2UvMmNiMTRlZjAwZmM3OTRmYTlhOWZlZDljNDAzMTEzYTgzYTYzZjM0NS5q
		謬紗特 (LV6)
		space.bilibili.com/59442895
	*/
}
