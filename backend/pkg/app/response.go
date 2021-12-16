package app

import (
	"github.com/csDeng/blog/pkg/e"
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
		"api":  "v1",
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})
}
