package context

import "github.com/farseer-go/redis"

type RedisContext struct {
	Redis redis.IClient `inject:"default"` // 使用farseer.yaml的Redis.default配置节点，并自动注入
}
