package server

import (
	"coconut/api"
	M "coconut/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "OK",
	})
}

func Search(ctx *gin.Context) {
	var data M.RequestForm

	err := ctx.ShouldBind(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "error parsing json",
		})

		return
	}

	if len(data.Keyword) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "the parameter is empty",
		})

		return
	}

	total, duration, msg, records, err := api.GetM3u8(data.Keyword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  msg,
		"data": gin.H{
			"time":    duration,
			"total":   total,
			"records": records,
		},
	})
}
