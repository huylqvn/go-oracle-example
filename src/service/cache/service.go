package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheService interface {
	Ping() error
	Put(key string, value string) error
	Get(key string) (string, error)
	GetClusterSlots() ([]redis.ClusterSlot, error)
	GetCluster() *redis.ClusterClient
}

type service struct {
	cmd     *redis.Client
	cluster *redis.ClusterClient
}

func NewService(hostPort string, password string) CacheService {
	if hostPort == ":" {
		return nil
	}
	cmd := redis.NewClient(&redis.Options{
		Addr:     hostPort,
		Password: password,
	})

	slots, err := cmd.ClusterSlots(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	var addrs []string
	for _, slot := range slots {
		for _, node := range slot.Nodes {
			addrs = append(addrs, node.Addr)
		}
	}

	cluster := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        addrs,
		Password:     password,
		PoolSize:     50,
		PoolTimeout:  30 * time.Second,
		MinIdleConns: 2,
		IdleTimeout:  300 * time.Second,
	})

	return &service{
		cmd:     cmd,
		cluster: cluster,
	}
}

func (s *service) Ping() error {
	_, err := s.cluster.Ping(context.Background()).Result()
	if err == nil {
		fmt.Println("Redis cluster is available")
	} else {
		fmt.Println("Can't connect redis cluster", err)
	}
	return err
}

func (s *service) Put(key string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := s.cluster.Set(ctx, key, value, 5*time.Minute).Err()
	return err
}

func (s *service) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	r, err := s.cluster.Get(ctx, key).Result()
	return r, err
}

func (s *service) GetClusterSlots() ([]redis.ClusterSlot, error) {
	slot, err := s.cmd.ClusterSlots(context.Background()).Result()
	return slot, err
}

func (s *service) GetCluster() *redis.ClusterClient {
	return s.cluster
}
