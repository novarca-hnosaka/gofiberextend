package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	gofiber_extend "github.com/novarca-hnosaka/gofiber_extend"
	"github.com/redis/go-redis/v9"
)

func Routes(ex *gofiber_extend.IFiberEx) {
	app := ex.App
	ctx := context.Background()

	app.Get("/", func(c *fiber.Ctx) error {
		ex.Log.Debug("test")
		return ex.Result(c, 200, map[string]interface{}{"status": "ok"})
	})

	api := app.Group("api/v1")
	api.Get("/test", func(c *fiber.Ctx) error {
		ex.Log.Debug("test")
		return ex.Result(c, 200, map[string]interface{}{"status": "ok"})
	})
	api.Post("/test", func(c *fiber.Ctx) error {
		ex.Log.Debug("test")
		if err := ex.Redis.Set(ctx, "test_key", "foo", 0).Err(); err != nil {
			return ex.ResultError(c, 500, err)
		}
		cmd := ex.Redis.Get(ctx, "test_key1")
		if err := cmd.Err(); err != nil {
			ex.Log.Debug(fmt.Sprintf("test_key1: %s", err))
		} else {
			rs, err := cmd.Result()
			if err == nil {
				ex.Log.Debug(fmt.Sprintf("test_key1: %s", rs))
			}
		}
		if err := ex.SetRedisJson("test_key", map[string]string{"bar": "foo"}, time.Duration(1*time.Hour)); err != nil {
			return ex.ResultError(c, 500, err)
		}
		rs := map[string]string{}
		if err := ex.GetRedisJson(&rs, "test_key"); err != nil {
			return ex.ResultError(c, 500, err)
		}
		return ex.Result(c, 200, map[string]interface{}{"status": "ok"}, rs)
	})

}

func main() {
	ex := gofiber_extend.New(gofiber_extend.IFiberExConfig{
		DevMode: gofiber_extend.Bool(true),
		UseDB:   true,
		DBConfig: &gofiber_extend.IDBConfig{
			Addr: "db:3306",
			User: "root",
			Pass: "qwerty",
		},
		UseRedis: true,
		RedisOptions: &redis.Options{
			Addr:         "redis:6379",
			Username:     "",
			Password:     "",
			DB:           0,
			MaxRetries:   5,
			PoolSize:     100,
			MinIdleConns: 10,
			MaxIdleConns: 100,
			TLSConfig:    nil,
		},
		// UseES: true,
		// ESConfig: &elasticsearch.Config{
		// 	Addresses: []string{"es:9200"},
		// },
	})
	app := ex.NewApp()
	Routes(ex)

	if err := app.Listen(":80"); err != nil {
		ex.Log.Fatal(err.Error())
	}
}
