package main

import (
	api "WP/biz/api/proto"
	"WP/biz/repos"
	"WP/pkg/postgres"
	"WP/pkg/redis"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type BizHandler struct {
	api.UnimplementedBizServer
	RedisDB  *redis.Redis
	MasterPg *postgres.PGXDatabase
	UserRepo repos.UsersRepository
}

func (b *BizHandler) GetUsers(ctx context.Context, req *api.RequestGetUsers) (*api.ResponseGetUsers, error) {
	var users []*api.User

	if req.UserId != 0 {
		user, err := b.UserRepo.FindById(ctx, req.UserId)
		if err != nil {
			log.Printf("read from postgres failed: %s", err)
			return nil, err
		}
		users = append(users, UserToProto(user))
	} else {
		userObjects, err := b.UserRepo.FindByLimit(ctx, 100)
		if err != nil {
			log.Printf("read from postgres failed: %s", err)
			return nil, err
		}
		for _, userObj := range userObjects {
			users = append(users, UserToProto(userObj))
		}
	}

	return &api.ResponseGetUsers{
		Users:     users,
		MessageId: 100, // TODO
	}, nil
}

func (b *BizHandler) GetUsersWithSqlInjection(ctx context.Context, req *api.RequestGetUsersWithSQLInjection) (*api.ResponseGetUsers, error) {
	var users []*api.User

	if req.UserId != "" {
		user, err := b.UserRepo.FindByIdWithInjection(ctx, req.UserId)
		if err != nil {
			log.Printf("read from postgres failed: %s", err)
			return nil, err
		}
		users = append(users, UserToProto(user))
	} else {
		userObjects, err := b.UserRepo.FindByLimit(ctx, 100)
		if err != nil {
			log.Printf("read from postgres failed: %s", err)
			return nil, err
		}
		for _, userObj := range userObjects {
			users = append(users, UserToProto(userObj))
		}
	}

	return &api.ResponseGetUsers{
		Users:     users,
		MessageId: 100, // TODO
	}, nil
}

func main() {
	handler := &BizHandler{}
	handler.RedisDB = redis.NewRedisWithOption(redis.Option{
		Host:       "localhost",
		Port:       "6379",
		PoolSize:   10,
		DB:         0,
		Pass:       "",
		MaxRetries: 1,
	})
	handler.MasterPg = postgres.NewPGXPostgres(postgres.Option{
		Host: "localhost",
		Port: 5432,
		User: "postgres",
		Pass: "1234",
		Db:   "users",
	}, 1000)
	handler.UserRepo = repos.NewUsersRepository(handler.MasterPg)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5062))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterBizServer(s, handler)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("biz")
}
