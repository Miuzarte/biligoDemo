package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Miuzarte/biligo"
	"github.com/mdp/qrterminal/v3"
)

func main() {
	qrcodeUrl, it, err := biligo.Login()
	if err != nil {
		panic(err)
	}

	qrterminal.Generate(qrcodeUrl, qrterminal.L, os.Stdout)

	for code, err := range it {
		if err != nil {
			panic(err)
		}
		switch code {
		case biligo.LOGIN_CODE_STATE_SUCCESS:
			fmt.Println("登录成功")
			id := biligo.ExportIdentity()
			_ = json.NewEncoder(os.Stdout).Encode(id)
			return

		case biligo.LOGIN_CODE_STATE_EXPIRED:
			fmt.Println("二维码已失效")
			return

		case biligo.LOGIN_CODE_STATE_SCANED:
			fmt.Println("已扫码")
		case biligo.LOGIN_CODE_STATE_UNSCANED:
			fmt.Println("未扫码")
		}
	}
}
