package redis_db

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

func (uc *redisDB) WriteRedis(key string, val string, exp int) error {
	// store data using SET command
	err := uc.redis.Set(context.Background(), key, val, time.Duration(exp)*time.Second).Err()
	if err != nil {
		logrus.Errorf("Error on :%v, unable to SET data. error: %v", "WriteRedis DB Function", err)
		return err
	}
	logrus.Printf("set operation success on WriteRedis DB Function:",val)

	return nil
}

func (uc *redisDB) ReadRedis(key string) (string, error) {
	// get data
	res, err := uc.redis.Get(context.Background(), key).Result()
	if err != nil {
		logrus.Errorf("Error on :%v, unable to GET data. error: %v", "ReadRedis DB Function", err)
		return res, err
	}
	logrus.Printf("get operation success on ReadRedis DB Function. result:", res)

	return res, nil
}
