package predis

func (r *RedisDb) SAdd(key string, members ...interface{}) (int64, error) {
	if cluster {
		return redisClusterClient.SAdd(ctx, key, members).Result()
	} else {
		return client.SAdd(ctx, key, members).Result()
	}
}

func (r *RedisDb) SCard(key string) (int64, error) {
	if cluster {
		return redisClusterClient.SCard(ctx, key).Result()
	} else {
		return client.SCard(ctx, key).Result()
	}
}

func (r *RedisDb) SRem(key string, members ...interface{}) (int64, error) {
	if cluster {
		return redisClusterClient.SRem(ctx, key, members).Result()
	} else {
		return client.SRem(ctx, key, members).Result()
	}
}

func (r *RedisDb) SPop(key string, count int64) ([]string, error) {
	if cluster {
		return redisClusterClient.SPopN(ctx, key, count).Result()
	} else {
		return client.SPopN(ctx, key, count).Result()
	}
}

func (r *RedisDb) SIsMember(key string, member interface{}) (bool, error) {
	if cluster {
		return redisClusterClient.SIsMember(ctx, key, member).Result()
	} else {
		return client.SIsMember(ctx, key, member).Result()
	}
}

func (r *RedisDb) SMembers(key string) ([]string, error) {
	if cluster {
		return redisClusterClient.SMembers(ctx, key).Result()
	} else {
		return client.SMembers(ctx, key).Result()
	}
}
