// Copyright 2024 Robert Cronin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var client *redis.Client

// Config holds the configuration for Redis
type Config struct {
	Host     string
	Port     int
	Password string
	DB       int
}

// New creates a new Redis client
func New(cfg Config) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// Ping the Redis server to check if the connection is alive
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := redisClient.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return redisClient, nil
}

// Init creates a new global Redis client
func Init() {
	// Get password from environment variable (injected by k8s or empty if on local)
	password := viper.GetString("REDIS_PASSWORD")
	cfg := Config{
		Host:     viper.GetString("redis.host"),
		Port:     viper.GetInt("redis.port"),
		Password: password,
		DB:       viper.GetInt("redis.db"),
	}

	redisClient, err := New(cfg)
	if err != nil {
		panic(err)
	}

	client = redisClient
}

// Implements fiber.Storage
type FiberClient struct {
	keyPrefix string
}

// Rewriting the storage implementation because fiber flushes the DB on every restart!
// https://github.com/gofiber/storage/blob/169c89147490500d9e74ec9234bf108f144b0f95/redis/redis.go#L108
var _ fiber.Storage = (*FiberClient)(nil)

// Close implements fiber.Storage::Close
func (f *FiberClient) Close() error {
	// no-op
	return nil
}

// Delete implements fiber.Storage::Delete
func (f *FiberClient) Delete(key string) error {
	return client.Del(context.Background(), f.keyPrefix+key).Err()
}

// Get implements fiber.Storage::Get
func (f *FiberClient) Get(key string) ([]byte, error) {
	data, err := client.Get(context.Background(), f.keyPrefix+key).Bytes()

	if err == redis.Nil {
		return nil, nil
	}

	return data, err
}

// Reset implements fiber.Storage::Reset
func (f *FiberClient) Reset() error {
	// clear all keys with the prefix
	keys, err := client.Keys(context.Background(), f.keyPrefix+"*").Result()
	if err != nil {
		return err
	}

	for _, key := range keys {
		if err := client.Del(context.Background(), f.keyPrefix+key).Err(); err != nil {
			return err
		}
	}

	return nil
}

// Set implements fiber.Storage::Set
func (f *FiberClient) Set(key string, val []byte, exp time.Duration) error {
	return client.Set(context.Background(), f.keyPrefix+key, val, exp).Err()
}

// NewFiberClient creates a new FiberClient
func NewFiberClient(keyPrefix string) *FiberClient {
	fc := &FiberClient{
		keyPrefix: keyPrefix,
	}

	return fc
}
