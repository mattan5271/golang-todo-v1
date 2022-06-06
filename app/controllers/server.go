package controllers

import (
	"net/http"
	"todo_app/config"
)

func StartMainServer() error {
	http.HandleFunc("/", top)

	// 第2引数にnilを渡すことで、デフォルト設定(ページが存在しない場合に404を返す)を使用する
	return http.ListenAndServe(":"+config.Config.Port, nil)
}