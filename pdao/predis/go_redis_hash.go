package predis

// HSet 某个的hash表的字段设置
func (r *RedisDb) HSet(key string, field string, value interface{}) error {
	if cluster {
		return redisClusterClient.HSet(ctx, key, field, value).Err()
	} else {
		return client.HSet(ctx, key, field, value).Err()
	}
}

// HGet 某个的hash表的字段获取
func (r *RedisDb) HGet(key string, field string) (string, error) {
	if cluster {
		return redisClusterClient.HGet(ctx, key, field).Result()
	} else {
		return client.HGet(ctx, key, field).Result()
	}
}

// HIncrBy 增加某个的hash表的字段的值
func (r *RedisDb) HIncrBy(key, field string, incr int64) error {
	if cluster {
		return redisClusterClient.HIncrBy(ctx, key, field, incr).Err()
	} else {
		return client.HIncrBy(ctx, key, field, incr).Err()
	}
}

// HDel 某个的hash表的删除某些字段
func (r *RedisDb) HDel(key string, fields ...string) (int64, error) {
	if cluster {
		return redisClusterClient.HDel(ctx, key, fields...).Result()
	} else {
		return client.HDel(ctx, key, fields...).Result()
	}
}

// HKeys 某个hash表的key模糊查询
func (r *RedisDb) HKeys(key string) ([]string, error) {
	if cluster {
		return redisClusterClient.HKeys(ctx, key).Result()
	} else {
		return client.HKeys(ctx, key).Result()
	}
}

// HExists 某个的hash表里是否存在此字段
func (r *RedisDb) HExists(key string, field string) (bool, error) {
	if cluster {
		return redisClusterClient.HExists(ctx, key, field).Result()
	} else {
		return client.HExists(ctx, key, field).Result()
	}
}
