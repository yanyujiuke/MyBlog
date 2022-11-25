package v1

import (
	"github.com/gin-gonic/gin"
	"myblog/model"
	"myblog/utils/errmsg"
	"net/http"
	"strconv"
)

func GetProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	pro, code := model.GetProfile(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    pro,
	})
}

func UpdateProfile(c *gin.Context) {
	var pro model.Profile
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&pro)

	code := model.UpdateProfile(id, &pro)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
