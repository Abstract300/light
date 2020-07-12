package main

import "github.com/Necroforger/dgrouter/exrouter"

func init() {
	router = exrouter.New()

	// Help message
	router.On("help", func(ctx *exrouter.Context) {
		ctx.Reply("generic help message.")
	})

	//
}
