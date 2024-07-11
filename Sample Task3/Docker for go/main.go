package main

import (
    "errors"
    "fmt"
    "time"

    "github.com/redis/go-redis/v9"

    "gofr.dev/pkg/gofr"
)

func main() {
    // Initialize GoFr object
    app := gofr.New()

    app.GET("/redis", func(ctx *gofr.Context) (interface{}, error) {
        // Get the value using the Redis instance

        fmt.Println("setting greeting")
        redisStatus := ctx.Redis.Set(ctx.Context, "greeting", "welcome to redis", time.Hour)
        if redisStatus.Err() != nil {
            return nil, redisStatus.Err()
        }

        val, err := ctx.Redis.Get(ctx.Context, "greeting").Result()
        if err != nil && !errors.Is(err, redis.Nil) {
            // If the key is not found, we are not considering this an error and returning ""
            return nil, err
        }

        return val, nil
    })

    // Run the application

    app.Run()
}
