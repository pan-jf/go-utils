package predis

import (
	"time"
)

// Del 删除一个key
func (r *RedisDb) Del(key string) (int64, error) {
	if cluster {
		return redisClusterClient.Del(ctx, key).Result()
	} else {
		return client.Del(ctx, key).Result()
	}
}

// DelKeys 删除多个key
func (r *RedisDb) DelKeys(keys ...string) error {
	if cluster {
		return redisClusterClient.Del(ctx, keys...).Err()
	} else {
		return client.Del(ctx, keys...).Err()
	}
}

// Exists 键是否存在
func (r *RedisDb) Exists(key string) (int64, error) {
	if cluster {
		return redisClusterClient.Exists(ctx, key).Result()
	} else {
		return client.Exists(ctx, key).Result()
	}
}

// Expire 设置过期时间
func (r *RedisDb) Expire(key string, expiration time.Duration) error {
	if cluster {
		return redisClusterClient.Expire(ctx, key, expiration).Err()
	} else {
		return client.Expire(ctx, key, expiration).Err()
	}
}

// Set set操作
func (r *RedisDb) Set(key string, value interface{}, expiration time.Duration) error {
	if cluster {
		return redisClusterClient.Set(ctx, key, value, expiration).Err()
	} else {
		return client.Set(ctx, key, value, expiration).Err()
	}
}

// SetNX 带过期时间的set操作
func (r *RedisDb) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	if cluster {
		return redisClusterClient.SetNX(ctx, key, value, expiration).Result()
	} else {
		return client.SetNX(ctx, key, value, expiration).Result()
	}
}

// Get 获取一个字符串
func (r *RedisDb) Get(key string) (string, error) {
	if cluster {
		return redisClusterClient.Get(ctx, key).Result()
	} else {
		return client.Get(ctx, key).Result()
	}
}

// GetInt64 获取一个64位整型
func (r *RedisDb) GetInt64(key string) (int64, error) {
	if cluster {
		return redisClusterClient.Get(ctx, key).Int64()
	} else {
		return client.Get(ctx, key).Int64()
	}
}

// Incr 自增1
func (r *RedisDb) Incr(key string) error {
	if cluster {
		return redisClusterClient.Incr(ctx, key).Err()
	} else {
		return client.Incr(ctx, key).Err()
	}
}

// Decr 自减1
func (r *RedisDb) Decr(key string) error {
	if cluster {
		return redisClusterClient.Decr(ctx, key).Err()
	} else {
		return client.Decr(ctx, key).Err()
	}
}

// IncrBy 自增特定数
func (r *RedisDb) IncrBy(key string, value int64) error {
	if cluster {
		return redisClusterClient.IncrBy(ctx, key, value).Err()
	} else {
		return client.IncrBy(ctx, key, value).Err()
	}
}

// DecrBy 自减特定数
func (r *RedisDb) DecrBy(key string, value int64) error {
	if cluster {
		return redisClusterClient.DecrBy(ctx, key, value).Err()
	} else {
		return client.DecrBy(ctx, key, value).Err()
	}
}

// Keys 模糊查询，尽量别用吧
func (r *RedisDb) Keys(pattern string) ([]string, error) {
	if cluster {
		return redisClusterClient.Keys(ctx, pattern).Result()
	} else {
		return client.Keys(ctx, pattern).Result()
	}
}
