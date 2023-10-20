package usecase

import (
	"merchant/models"
)

func (uc *usecase) WriteRedis(req models.RedisReq) error {
	err := uc.redis.WriteRedis(req.Key,req.Data,req.Exp)
	if err != nil {
		return err
	}

	return nil
}

func (uc *usecase) ReadRedis(req models.RedisReq) (string, error) {
	res, err := uc.redis.ReadRedis(req.Key)
	if err != nil {
		return res, err
	}

	return res, nil
}
