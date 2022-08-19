package predis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisDb struct {
	IsInit bool `json:"isInit"`
}

var GlobalRedis = &RedisDb{
	IsInit: false,
}

const RedisNil = "redis: nil"

// client redis单机客户端
var client *redis.Client

// redisClusterClient redis集群客户端
var redisClusterClient *redis.ClusterClient

// cluster 是否集群
var cluster bool

var ctx = context.Background()

type RedisCfg struct {
	Host     string   `json:"host"`
	Password string   `json:"password"`
	DB       int      `json:"db"`
	Cluster  bool     `json:"cluster"`
	Hosts    []string `json:"hosts"`
	PoolSize int      `json:"poolSize"`
}

// Setup 初始化
func (r *RedisDb) Setup(redisCfg *RedisCfg) error {
	cluster = redisCfg.Cluster
	//判断是否为集群配置
	if redisCfg.Cluster {
		return r.SetupCluster(redisCfg)
	} else {
		return r.SetupClient(redisCfg)
	}
}

// SetupCluster 初始化一个redis集群
func (r *RedisDb) SetupCluster(redisCfg *RedisCfg) error {
	//ClusterClient是一个Redis集群客户机，表示一个由0个或多个底层连接组成的池。它对于多个goroutine的并发使用是安全的。
	opt := &redis.ClusterOptions{
		Password: redisCfg.Password,
		Addrs:    redisCfg.Hosts,
	}
	if redisCfg.PoolSize > 0 {
		opt.PoolSize = redisCfg.PoolSize
	}
	redisClusterClient = redis.NewClusterClient(opt)

	//Ping
	_, err := redisClusterClient.Ping(ctx).Result()
	if err != nil {
		return err
	}

	r.IsInit = true
	return nil
}

// SetupClient 初始化一个redis客户端
func (r *RedisDb) SetupClient(redisCfg *RedisCfg) error {
	//Redis客户端，由零个或多个基础连接组成的池。它对于多个goroutine的并发使用是安全的。
	//更多参数参考Options结构体
	opt := &redis.Options{
		Addr:     redisCfg.Host,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	}
	if redisCfg.PoolSize > 0 {
		opt.PoolSize = redisCfg.PoolSize
	}
	client = redis.NewClient(opt)
	//Ping
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return err
	}

	r.IsInit = true
	return nil
}

// Client 获取客户端
func (r *RedisDb) Client() interface{} {
	if cluster {
		return redisClusterClient
	} else {
		return client
	}
}

// Ping 检测链接是否正常
// 一段时间后redis操作前，需要调用一下Ping(), 如果已经发生断开，会自动先触发重连，避免后继操作出错
func (r *RedisDb) Ping() *redis.StatusCmd {
	var result *redis.StatusCmd
	if cluster {
		result = redisClusterClient.Ping(ctx)
		// 重连一次
		if result.Err() != nil {
			result = redisClusterClient.Ping(ctx)
		}
	} else {
		result = client.Ping(ctx)
		// 重连一次
		if result.Err() != nil {
			result = client.Ping(ctx)
		}
	}

	return result
}

// Subscribe 订阅
func (r *RedisDb) Subscribe(channels ...string) *redis.PubSub {
	if cluster {
		return redisClusterClient.Subscribe(ctx, channels...)
	} else {
		return client.Subscribe(ctx, channels...)
	}
}

// Publish 发布
func (r *RedisDb) Publish(channel string, message interface{}) *redis.IntCmd {
	if cluster {
		return redisClusterClient.Publish(ctx, channel, message)
	} else {
		return client.Publish(ctx, channel, message)
	}
}

// Close 关闭
func (r *RedisDb) Close() error {
	if cluster {
		return redisClusterClient.Close()
	} else {
		return client.Close()
	}
}
