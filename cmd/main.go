package main

import (
	"github.com/earelin/pega/pkg"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	e := gin.Default()
	pkg.ApplicationConfig(e)
	err := e.Run()
	if err != nil {
		return
	}
}
