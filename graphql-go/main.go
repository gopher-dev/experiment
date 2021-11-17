package main

import (
	"github.com/gopher-dev/experiment/graphql-go/resolver"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func main() {
	e := echo.New()
	h := resolver.Handler{}
	h.Initialize("./schema/", &resolver.Resolver{})
	h.ServeEcho(e)
	log.Fatal(e.Start(":8080"))
}
