package controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Testshutdown(ctx *gin.Context) {
	fmt.Println("testrestart start")
	time.Sleep(7 * time.Second)
	fmt.Println("testrestart ent")
}
