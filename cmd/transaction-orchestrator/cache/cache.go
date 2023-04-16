package cache

import (
	"fmt"
	"time"

	"github.com/casperfj/bachelor/cmd/transaction-orchestrator/config"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type Cache struct {
	redisClient *redis.Client
}

func NewCache(conf *config.Configuration) (*Cache, error) {
	// Connect to redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port),
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})

	// Check connection
	_, err := redisClient.Ping(redisClient.Context()).Result()
	if err != nil {
		return nil, err
	}

	// Initialize cache
	cache := &Cache{
		redisClient: redisClient,
	}

	return cache, nil
}

func (c *Cache) BlockAccountIDs(accountIDs []uuid.UUID) error {
	var err error = nil
	var i int

	// Insert account IDs into cache
	for i = 0; i < len(accountIDs); i++ {
		setIfNotExist, newErr := c.redisClient.SetNX(c.redisClient.Context(), accountIDs[i].String(), time.Now().UTC().Unix(), 0).Result()
		if newErr != nil {
			err = newErr
			break
		}

		// Skip the rest of the account IDs if error or an account was already in the cache
		if !setIfNotExist {
			err = fmt.Errorf("account ID %s is already in the cache", accountIDs[i].String())
			break
		}
	}

	// Release account IDs that were successfully inserted into the cache if error or an account was already in the cache
	if err != nil {
		c.ReleaseAccountIDs(accountIDs[:i])
	}

	// Return error
	return err
}

func (c *Cache) ReleaseAccountIDs(accountIDs []uuid.UUID) error {
	var err error = nil

	// Delete account IDs from cache
	for _, accountID := range accountIDs {
		newErr := c.redisClient.Del(c.redisClient.Context(), accountID.String()).Err()
		if newErr != nil {
			err = newErr
		}
	}

	// Return error
	return err
}

func (c *Cache) IsInTransaction(accountIDs []uuid.UUID) (bool, error) {
	// Check if any of the account IDs are in the cache
	for _, accountID := range accountIDs {
		// Check if account ID is in cache
		_, err := c.redisClient.Get(c.redisClient.Context(), accountID.String()).Result()
		if err != nil {
			// Continue if account ID was not in the cache
			if err == redis.Nil {
				continue
			}

			return false, err
		}

		// Return true because an account ID was in the cache
		return true, nil
	}

	// Return false because no account ID was in the cache
	return false, nil
}
