package main

import (
	"net/http"

	"github.com/savsgio/atreugo/v11"
)

type responseMSG struct {
	Message string `json:"message"`
}

func main() {
	config := atreugo.Config{
		Addr: "0.0.0.0:8000",
	}

	server := atreugo.New(config)

	server.GET("/get", func(ctx *atreugo.RequestCtx) error {
		m := responseMSG{"all good"}

		return ctx.JSONResponse(m, http.StatusOK)
	})

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
