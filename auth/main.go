package main

import (
	api "WP/auth/api/proto"
	"WP/auth/models"
	"WP/pkg/redis"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const BBound = 10

type AuthHandler struct {
	api.UnimplementedAuthServer
	RedisDB *redis.Redis
}

func (a *AuthHandler) ReqPq(ctx context.Context, req *api.RequestReqPQ) (*api.ResponseReqPQ, error) {
	serverNonce := RandString(20)
	redisKey := GetSHA1EncodedString(req.Nonce + serverNonce)
	//p, g := GeneratePG()
	var p, g int64 = 23, 5
	pgHolder := &models.PGHolder{
		G: g,
		P: p,
	}
	err := SetRedis(ctx, a.RedisDB, redisKey, pgHolder, time.Minute*20)
	if err != nil {
		log.Printf("insert to redis failed: %s", err.Error())
	}

	return &api.ResponseReqPQ{
		Nonce:       req.Nonce,
		ServerNonce: serverNonce,
		P:           p,
		G:           g,
		MessageId:   100, // TODO
	}, nil
}

func (a *AuthHandler) Req_DHParams(ctx context.Context, req *api.RequestReqDHParams) (*api.ResponseReqDHParams, error) {
	//b := rand.Intn(BBound)
	var b int64 = 15

	redisKey := GetSHA1EncodedString(req.Nonce + req.ServerNonce)
	value, err := a.RedisDB.Get(ctx, redisKey).Result()
	if err != nil {
		log.Printf("read from redis failed: %s", err)
		return nil, err
	}
	var pgHolder *models.PGHolder
	err = json.Unmarshal([]byte(value), &pgHolder)
	if err != nil {
		log.Printf("json unmarshal failed: %s", err)
		return nil, err
	}

	B := CalculatePublic(pgHolder.G, pgHolder.P, b)
	publicKey := &models.CommonKeyHolder{
		Key: B,
	}

	privateKey := CalculatePrivate(pgHolder.P, req.A, b)
	log.Printf("private: %d", privateKey)

	err = SetRedis(ctx, a.RedisDB, redisKey, publicKey, time.Minute*20)
	if err != nil {
		log.Printf("insert to redis failed: %s", err.Error())
	}

	return &api.ResponseReqDHParams{
		B:           int64(int32(B)),
		Nonce:       req.Nonce,
		ServerNonce: req.ServerNonce,
		MessageId:   100, // TODO
	}, nil
}

func main() {
	handler := &AuthHandler{}
	handler.RedisDB = redis.NewRedisWithOption(redis.Option{
		Host:       "localhost",
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
