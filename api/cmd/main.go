package main

import (
	"context"

	"github.com/mori-hisayuki/go-study/api/internal/presenter"
)

// @title			ユーザー管理サービスAPI
// @version		0.1.0
// @description	ユーザー管理サービスAPIの機能
// @host			localhost:8080
func main() {
	server := presenter.NewServer()
	if err := server.Run(context.Background()); err != nil {
		panic(err)
	}
}
