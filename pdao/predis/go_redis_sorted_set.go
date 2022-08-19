package predis

import "github.com/go-redis/redis/v8"

func (r *RedisDb) ZAdd(key string, members *redis.Z) (int64, error) {
	if cluster {
		return redisClusterClient.ZAdd(ctx, key, members).Result()
	} else {
		return client.ZAdd(ctx, key, members).Result()
	}
}

func (r *RedisDb) ZRem(key string, members ...interface{}) (int64, error) {
	if cluster {
		return redisClusterClient.ZRem(ctx, key, members).Result()
	} else {
		return client.ZRem(ctx, key, members).Result()
	}
}

func (r *RedisDb) ZScore(key string, member string) (float64, error) {
	if cluster {
		return redisClusterClient.ZScore(ctx, key, member).Result()
	} else {
		return client.ZScore(ctx, key, member).Result()
	}
}

func (r *RedisDb) ZRangeByScore(key string, opt *redis.ZRangeBy) ([]string, error) {
	if cluster {
		return redisClusterClient.ZRangeByScore(ctx, key, opt).Result()
	} else {
		return client.ZRangeByScore(ctx, key, opt).Result()
	}
}

func (r *RedisDb) ZRevRangeWithScores(key string, start, stop int64) ([]redis.Z, error) {
	if cluster {
		return redisClusterClient.ZRevRangeWithScores(ctx, key, start, stop).Result()
	} else {
		return client.ZRevRangeWithScores(ctx, key, start, stop).Result()
	}
}

func (r *RedisDb) ZRevRangeByScoreWithScores(key string, opt *redis.ZRangeBy) ([]redis.Z, error) {
	if cluster {
		return redisClusterClient.ZRevRangeByScoreWithScores(ctx, key, opt).Result()
	} else {
		return client.ZRevRangeByScoreWithScores(ctx, key, opt).Result()
	}
}

func (r *RedisDb) ZCard(key string) (int64, error) {
	if cluster {
		return redisClusterClient.ZCard(ctx, key).Result()
	} else {
		return client.ZCard(ctx, key).Result()
	}
}

func (r *RedisDb) ZRemRangeByRank(key string, start, stop int64) (int64, error) {
	if cluster {
		return redisClusterClient.ZRemRangeByRank(ctx, key, start, stop).Result()
	} else {
		return client.ZRemRangeByRank(ctx, key, start, stop).Result()
	}
}
