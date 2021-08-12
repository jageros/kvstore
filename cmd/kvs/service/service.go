package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jageros/attribute"
	"net/http"
	"strsvc/internal/utils"
)

const DBID = "str_data"

func getValue(c *gin.Context) {
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

func getValues(c *gin.Context) {
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

func saveValues(c *gin.Context) {
	args := map[string]interface{}{}
	err := c.BindJSON(&args)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error()})
		return
	}
	attr := attribute.NewAttrMgr("data", DBID)
	attr.Load(true)
	for key, val := range args {
		attr.Set(key, val)
	}
	err = attr.Save(false)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -2, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "successful"})
}

func RegisterHandle(r *gin.RouterGroup) {
	r.GET("/value", getValue)
	r.GET("/values", getValues)
	r.POST("/values", saveValues)
}
