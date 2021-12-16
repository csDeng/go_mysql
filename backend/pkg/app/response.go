package app

import (
	"github.com/csDeng/go-gin-example/pkg/e"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

/*
appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)

*/
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})
}
