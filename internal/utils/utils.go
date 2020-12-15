package utils

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/url"
)

func CatchPanic(f func()) (err interface{}) {
	defer func() {
		err = recover()
		if err != nil {
			log.Printf("panic: %v", err)
		}
	}()

	f()
	return
}

func DecodeUrlVal(c *gin.Context, key string) string {
	val, err := url.ParseQuery(key + "=" + c.Query(key))
	if err == nil {
		return val[key][0]
	}
	return ""
}
