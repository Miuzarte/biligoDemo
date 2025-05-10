package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Miuzarte/biligo"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var err error
	var dd biligo.DynamicDetail
	var da biligo.DynamicAll
	var dau biligo.DynamicAllUpdate

	for { // 无限尝试获取 baseline
		da, err = biligo.FetchDynamicAll()
		if err != nil {
			fmt.Println("failed to fetch dynamic all: ", err)
			select {
			case <-time.After(time.Second * 10):
				continue
			case <-ctx.Done():
				return
			}
		}
		if da.UpdateBaseline == "" {
			fmt.Println("update baseline is empty")
			select {
			case <-time.After(time.Second * 10):
				continue
			case <-ctx.Done():
				return
			}
		}
		break
	}

	for { // 拉取更新
		dau, err = biligo.FetchDynamicAllUpdate(da.UpdateBaseline)
		if err != nil {
			fmt.Println("failed to fetch dynamic all update: ", err)
			goto FAILED
		}
		if dau.UpdateNum == 0 {
			goto WAIT
		}

		for range 3 { // 尝试 3 次
			da, err = biligo.FetchDynamicAll()
			if err == nil {
				break
			}
			<-time.After(time.Second * 10)
		}
		if err != nil {
			fmt.Println("failed to fetch dynamic all: ", err)
			goto FAILED
		}
		if len(da.Items) == 0 {
			fmt.Println("dynamic items is empty")
			goto FAILED
		}

		for i := range da.Items {
			if i >= dau.UpdateNum {
				break // 更新了几条就拉几条
			}

			dd = da.Items[i]
			fmt.Printf(
				"new dynamic %s from %d, type: %s\n",
				dd.IdStr, dd.Modules.Author.Mid, dd.Type,
			)
		}

	WAIT:
		select {
		case <-time.After(time.Second):
			continue
		case <-ctx.Done():
			return
		}

	FAILED:
		select {
		case <-time.After(time.Second * 10):
			continue
		case <-ctx.Done():
			return
		}
	}
}
