package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/redis/go-redis/v9"
	"healthcare-allopurinol/entity"
	"strconv"
	"time"
)

type RedisRepoItf interface {
	IsResetExist(email string) (bool, error)
	SetReset(email, token string) error
	GetResetValue(token string) (string, error)
	DeleteReset(token, email string) error
	IsVerifyExist(credId int64) (bool, error)
	SetVerify(credId int64, token string) error
	GetVerifyValue(token string) (int64, error)
	DeleteVerify(token string, credId int64) error
	SetShippingData(shipping *entity.Shipping) error
	GetShippingData(userId int64, pharmacyId int64) (*entity.Shipping, error)
	DeleteShippingData(userId int64, pharmacyId int64) error
}

type RedisRepoImpl struct {
	rdb *redis.Client
	ctx context.Context
}

func NewRedisRepo(rdb *redis.Client, ctx context.Context) *RedisRepoImpl {
	return &RedisRepoImpl{rdb: rdb, ctx: ctx}
}

func isExist(r *RedisRepoImpl, check string) (bool, error) {
	_, err := r.rdb.Get(r.ctx, check).Result()
	if errors.Is(err, redis.Nil) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func isExistValue(r *RedisRepoImpl, check string) (string, error) {
	val, err := r.rdb.Get(r.ctx, check).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (r *RedisRepoImpl) IsResetExist(email string) (bool, error) {
	return isExist(r, fmt.Sprintf("forgot-%s", email))
}

func (r *RedisRepoImpl) SetReset(email, token string) error {
	err := r.rdb.Set(r.ctx, fmt.Sprintf("forgot-%s", email), true, 5*time.Minute).Err()
	if err != nil {
		return err
	}

	err = r.rdb.Set(r.ctx, token, email, 5*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisRepoImpl) GetResetValue(token string) (string, error) {
	return isExistValue(r, token)
}

func (r *RedisRepoImpl) DeleteReset(token, email string) error {
	err := r.rdb.Del(r.ctx, fmt.Sprintf("forgot-%s", email)).Err()
	if err != nil {
		return err
	}

	err = r.rdb.Del(r.ctx, token).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisRepoImpl) IsVerifyExist(credId int64) (bool, error) {
	return isExist(r, fmt.Sprintf("verify-%d", credId))
}

func (r *RedisRepoImpl) SetVerify(credId int64, token string) error {
	err := r.rdb.Set(r.ctx, fmt.Sprintf("verify-%d", credId), true, 5*time.Minute).Err()
	if err != nil {
		return err
	}

	err = r.rdb.Set(r.ctx, token, credId, 5*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisRepoImpl) GetVerifyValue(token string) (int64, error) {
	id, err := isExistValue(r, token)
	if err != nil {
		return 0, err
	}

	credId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, err
	}

	return credId, nil
}

func (r *RedisRepoImpl) DeleteVerify(token string, credId int64) error {
	err := r.rdb.Del(r.ctx, fmt.Sprintf("verify-%d", credId)).Err()
	if err != nil {
		return err
	}

	err = r.rdb.Del(r.ctx, token).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisRepoImpl) SetShippingData(shipping *entity.Shipping) error {
	data, err := sonic.Marshal(shipping)
	if err != nil {
		return err
	}

	err = r.rdb.Set(r.ctx, fmt.Sprintf("cart-%d-%d", shipping.UserId, shipping.PharmacyId), data, 24*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisRepoImpl) GetShippingData(userId, pharmacyId int64) (*entity.Shipping, error) {
	data, err := r.rdb.Get(r.ctx, fmt.Sprintf("cart-%d-%d", userId, pharmacyId)).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var shipping *entity.Shipping
	err = sonic.Unmarshal([]byte(data), &shipping)
	if err != nil {
		return nil, err
	}

	return shipping, nil
}

func (r *RedisRepoImpl) DeleteShippingData(userId int64, pharmacyId int64) error {
	err := r.rdb.Del(r.ctx, fmt.Sprintf("cart-%d-%d", userId, pharmacyId)).Err()
	if err != nil {
		return err
	}

	return nil
}
