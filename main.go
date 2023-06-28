package main

import (
	"github.com/Natchalit/account/src/routex"
	"github.com/Natchalit/account/src/services/auth"
	"github.com/Natchalit/account/src/services/mainx"
)

func main() {

	mainx.Init()
	auth.Auth()
	routex.Route()
}
