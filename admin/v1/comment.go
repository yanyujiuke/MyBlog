package v1

import (
	"github.com/gin-gonic/gin"
	"myblog/model"
	"myblog/utils/errmsg"
	"net/http"
	"strconv"
)

func GetComments(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum <= 0 {
		pageNum = 1
	}

	data, total, code := model.GetComments(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"total":   total,
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteComment(id)

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg,
	})
}

func CheckComment(c *gin.Context) {
	var comment model.Comment
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&comment)

	code := model.CheckComment(id, &comment, true)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func UncheckComment(c *gin.Context) {
	var comment model.Comment
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&comment)

	code := model.CheckComment(id, &comment, false)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
