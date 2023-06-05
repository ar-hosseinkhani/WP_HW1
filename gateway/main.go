package main

import (
	api "WP/auth/api/proto"
	api2 "WP/biz/api/proto"
	"WP/gateway/api/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GatewayHandler struct {
	authClient  api.AuthClient
	bizClient   api2.BizClient
	httpHandler *http.Handler
}

func main() {
	handler := GatewayHandler{}

	authConn := getGRPCConn("127.0.0.1:5052")
	handler.authClient = api.NewAuthClient(authConn)

	bizConn := getGRPCConn("127.0.0.1:5062")
	handler.bizClient = api2.NewBizClient(bizConn)

	handler.httpHandler = http.NewHandler(handler.authClient, handler.bizClient)

	route := gin.Default()
	route.POST("/auth/req_pq", handler.httpHandler.ReqPQ)
	route.POST("/auth/req_dh_param", handler.httpHandler.ReqDHParam)

	route.POST("/biz/get_users", handler.httpHandler.GetUsers)
	route.POST("/biz/get_users_with_sql_injection", handler.httpHandler.GetUsersWithSQLInjection)

	err := route.Run(":8080")
	if err != nil {
		return
	}

}

func getGRPCConn(address string) *grpc.ClientConn {
	cc, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(fmt.Errorf("failed to connect to auth: %v", err))
	}
	return cc
}
