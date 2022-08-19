package predis

func (r *RedisDb) LPop(key string) (string, error) {
	if cluster {
		return redisClusterClient.LPop(ctx, key).Result()
	} else {
		return client.LPop(ctx, key).Result()
	}
}

func (r *RedisDb) LPush(key string, values ...interface{}) error {
	if cluster {
		return redisClusterClient.LPush(ctx, key, values).Err()
	} else {
		return client.LPush(ctx, key, values).Err()
	}
}

func (r *RedisDb) RPop(key string) (string, error) {
	if cluster {
		return redisClusterClient.RPop(ctx, key).Result()
	} else {
		return client.RPop(ctx, key).Result()
	}
}
func (r *RedisDb) RPush(key string, values ...interface{}) error {
	if cluster {
		return redisClusterClient.RPush(ctx, key, values...).Err()
	} else {
		return client.RPush(ctx, key, values...).Err()
	}
}

func (r *RedisDb) LRem(key string, count int64, values interface{}) (int64, error) {
	if cluster {
		return redisClusterClient.LRem(ctx, key, count, values).Result()
	} else {
		return client.LRem(ctx, key, count, values).Result()
	}
}

func (r *RedisDb) LLen(key string) (int64, error) {
	if cluster {
		return redisClusterClient.LLen(ctx, key).Result()
	} else {
		return client.LLen(ctx, key).Result()
	}
}

func (r *RedisDb) LRange(key string, start, stop int64) ([]string, error) {
	if cluster {
		return redisClusterClient.LRange(ctx, key, start, stop).Result()
	} else {
		return client.LRange(ctx, key, start, stop).Result()
	}
}

func (r *RedisDb) LIndex(key string, index int64) (string, error) {
	if cluster {
		return redisClusterClient.LIndex(ctx, key, index).Result()
	} else {
		return client.LIndex(ctx, key, index).Result()
	}
}

func (r *RedisDb) LTrim(key string, start, stop int64) (string, error) {
	if cluster {
		return redisClusterClient.LTrim(ctx, key, start, stop).Result()
	} else {
		return client.LTrim(ctx, key, start, stop).Result()
	}
}
