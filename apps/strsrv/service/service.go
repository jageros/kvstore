package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strsvc/internal/attribute"
	"strsvc/internal/utils"
)

const DBID = "str_data"

func handle_StrGet(c *gin.Context) {
	key := utils.DecodeUrlVal(c, "key")
	if key == "" {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误！"})
		return
	}

	attr := attribute.NewAttrMgr("data", DBID)
	err := attr.Load(true)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error()})
		return
	}
	val := attr.Get(key)
	result := map[string]interface{}{"value": val}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "successful", "data": result})
}

func handle_StrGetAll(c *gin.Context) {
	attr := attribute.NewAttrMgr("data", DBID)
	err := attr.Load(true)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error()})
		return
	}
	result := map[string]interface{}{}
	attr.ForEachKey(func(key string) bool {
		val := attr.Get(key)
		result[key] = val
		return true
	})
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "successful", "data": result})
}

func handle_StrSave(c *gin.Context) {
	args := map[string]string{}
	err := c.BindJSON(&args)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error()})
		return
	}
	attr := attribute.NewAttrMgr("data", DBID)
	attr.Load(true)
	for key, val := range args {
		attr.SetStr(key, val)
	}
	err = attr.Save(false)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -2, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "successful"})
}

func RegisterHandle(r *gin.RouterGroup) {
	r.GET("/str_get", handle_StrGet)
	r.GET("/str_get_all", handle_StrGetAll)
	r.POST("/str_save", handle_StrSave)
}
