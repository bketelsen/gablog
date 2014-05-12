package main

import (
	"encoding/json"
	"github.com/hoisie/web"
	eject "github.com/mattn/go-eject"
)

type result struct {
	Error interface{} `json:"error"`
}

func main() {
	web.Get("/eject", func(ctx *web.Context) {
		ctx.ContentType("application/json")
		if err := eject.Eject(); err != nil {
			json.NewEncoder(ctx).Encode(&result{err.Error()})
		} else {
			json.NewEncoder(ctx).Encode(&result{nil})
		}
	})
	web.Run(":8080")
}
