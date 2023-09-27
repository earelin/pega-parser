package main

import (
	"github.com/earelin/pega/pkg"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := gin.Default()
	pkg.ApplicationConfig(e)
	err := e.Run()
	if err != nil {
		return
	}
}
