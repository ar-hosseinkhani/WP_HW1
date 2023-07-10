package main

import (
	api "WP/auth/api/proto"
	"WP/auth/models"
	"WP/pkg/redis"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type AuthHandler struct {
	api.UnimplementedAuthServer
	RedisDB *redis.Redis
}

func (a *AuthHandler) ReqPq(ctx context.Context, req *api.RequestReqPQ) (*api.ResponseReqPQ, error) {
	serverNonce := RandString(20)
	redisKey := GetSHA1EncodedString(req.Nonce + serverNonce)
	p, g := GeneratePG()
	pgHolder := &models.PGHolder{
		G: g,
		P: p,
	}

	log.Printf("P: %d", p)
	log.Printf("G: %d", g)

	err := SetRedis(ctx, a.RedisDB, redisKey, pgHolder, time.Minute*20)
	if err != nil {
		log.Printf("insert to redis failed: %s", err.Error())
		return nil, err
	}

	return &api.ResponseReqPQ{
		Nonce:       req.Nonce,
		ServerNonce: serverNonce,
		P:           p,
		G:           g,
		MessageId:   req.MessageId + 1,
	}, nil
}

func (a *AuthHandler) Req_DHParams(ctx context.Context, req *api.RequestReqDHParams) (*api.ResponseReqDHParams, error) {
	//b := rand.Intn(BBound)
	var b int64 = 15

	redisKey := GetSHA1EncodedString(req.Nonce + req.ServerNonce)
	value, err := a.RedisDB.Get(ctx, redisKey).Result()
	if err != nil {
		log.Printf("read from redis failed: %s", err)
		return nil, errors.New("you should call ReqPQ first")
	}

	var pgHolder *models.PGHolder
	err = json.Unmarshal([]byte(value), &pgHolder)
	if err != nil {
		log.Printf("json unmarshal failed: %s", err)
		return nil, errors.New("you should call ReqPQ first")
	}

	publicKey := &models.KeyHolder{
		Key: CalculatePublic(pgHolder.G, pgHolder.P, b),
	}
	privateKey := &models.KeyHolder{
		Key: CalculatePrivate(pgHolder.P, req.A, b),
	}

	log.Printf("private: %d", privateKey.Key)
	log.Printf("public: %d", publicKey.Key)

	err = SetRedis(ctx, a.RedisDB, redisKey, privateKey, time.Minute*20)
	if err != nil {
		log.Printf("insert to redis failed: %s", err.Error())
	}

	return &api.ResponseReqDHParams{
		B:           publicKey.Key,
		Nonce:       req.Nonce,
		ServerNonce: req.ServerNonce,
		MessageId:   req.MessageId + 1,
	}, nil
}

func main() {
	handler := &AuthHandler{}
	handler.RedisDB = redis.NewRedisWithOption(redis.Option{
		Host:       "redis",
		Port:       "6379",
		PoolSize:   10,
		DB:         0,
		Pass:       "",
		MaxRetries: 1,
	})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5052))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterAuthServer(s, handler)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
