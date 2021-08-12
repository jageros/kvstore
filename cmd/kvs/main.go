package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"time"

	"kvstore/cmd/kvs/service"
	"kvstore/internal/conf"

	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jageros/attribute"
	"github.com/jageros/db"
	"github.com/jageros/group"
)

func main() {
	rand.Seed(time.Now().Unix())
	cfg := conf.Parse("config.yaml")

	g := group.Default()
	attribute.Initialize(g, func(opt *db.Option) {
		opt.Addr = cfg.Mongo.Addr
		opt.DBName = cfg.Mongo.DB
		opt.User = cfg.Mongo.User
		opt.Password = cfg.Mongo.Password
	})

	gin.SetMode(cfg.Model)
	engine := gin.Default()

	// gin解决跨域问题中间件
	//engine.Use(cors.Default())

	// 群组路由
	r := engine.Group("/api")

	// 注册 http handler
	service.RegisterHandle(r)

	// 创建服务
	svr := &http.Server{
		Addr:    cfg.ListenAddr,
		Handler: engine,
	}

	// 起服监听
	g.Go(func(_ context.Context) error {
		return svr.ListenAndServe()
	})

	// 优雅停服
	g.Go(func(ctx context.Context) error {
		<-ctx.Done()
		tCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return svr.Shutdown(tCtx)
	})

	// 阻塞等待所有goroutine停止，并返回停止的原因
	err := g.Wait()
	log.Printf("Stop With: %v", err)
}
