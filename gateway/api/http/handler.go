package http

import (
	api "WP/auth/api/proto"
	api2 "WP/biz/api/proto"
	"WP/gateway/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Handler struct {
	authClient api.AuthClient
	bizClient  api2.BizClient
}

func NewHandler(authClient api.AuthClient, bizClient api2.BizClient) *Handler {
	return &Handler{
		authClient: authClient,
		bizClient:  bizClient,
	}
}

func (h *Handler) ReqPQ(c *gin.Context) {
	var req *models.RequestReqPQ
	logPrefix := "http handler: ReqPQ:"
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Printf(logPrefix+" ShouldBindJSON: %v", err)
		return
	}
	if req.Nonce == "" || req.MessageId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid inputs",
		})
		return
	}

	resp, err := h.authClient.ReqPq(c, &api.RequestReqPQ{Nonce: req.Nonce, MessageId: req.MessageId})
	if err != nil {
		log.Printf(logPrefix+" call rpc failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &models.ResponseReqPQ{
		Nonce:       resp.Nonce,
		MessageId:   resp.MessageId,
		P:           resp.P,
		G:           resp.G,
		ServerNonce: resp.ServerNonce,
	})
}

func (h *Handler) ReqDHParam(c *gin.Context) {
	var req *models.RequestReqDHParams
	logPrefix := "http handler: ReqDHParam:"
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Printf(logPrefix+" ShouldBindJSON: %v", err)
		return
	}
	if req.Nonce == "" || req.ServerNonce == "" || req.A == 0 || req.MessageId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid inputs",
		})
		return
	}
	resp, err := h.authClient.Req_DHParams(c, &api.RequestReqDHParams{Nonce: req.Nonce, MessageId: req.MessageId, ServerNonce: req.ServerNonce, A: req.A})
	if err != nil {
		log.Printf(logPrefix+" call rpc failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &models.ResponseReqDHParams{
		Nonce:       resp.Nonce,
		MessageId:   resp.MessageId,
		B:           resp.B,
		ServerNonce: resp.ServerNonce,
	})
}

func (h *Handler) GetUsers(c *gin.Context) {
	var req *models.RequestGetUsers
	logPrefix := "http handler: GetUsers:"
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Printf(logPrefix+" ShouldBindJSON: %v", err)
		return
	}

	if req.UserId < 0 || req.MessageId == 0 || req.AuthKey == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid inputs",
		})
	}

	resp, err := h.bizClient.GetUsers(c, &api2.RequestGetUsers{MessageId: req.MessageId, UserId: req.UserId, AuthKey: req.AuthKey})
	if err != nil {
		log.Printf(logPrefix+" call rpc failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var users []*models.User
	for _, userProto := range resp.Users {
		users = append(users, models.UserProtoToModel(userProto))
	}
	c.JSON(http.StatusOK, &models.ResponseGetUsers{
		MessageId: resp.MessageId,
		Users:     users,
	})
}

func (h *Handler) GetUsersWithSQLInjection(c *gin.Context) {
	var req *models.RequestGetUsersWithInj
	logPrefix := "http handler: GetUsersWithSqlInjection:"
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Printf(logPrefix+" ShouldBindJSON: %v", err)
		return
	}
	if req.UserId == "" || req.MessageId == 0 || req.AuthKey == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid inputs",
		})
	}
	resp, err := h.bizClient.GetUsersWithSqlInjection(c, &api2.RequestGetUsersWithSQLInjection{MessageId: req.MessageId, UserId: req.UserId, AuthKey: req.AuthKey})
	if err != nil {
		log.Printf(logPrefix+" call rpc failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var users []*models.User
	for _, userProto := range resp.Users {
		users = append(users, models.UserProtoToModel(userProto))
	}
	c.JSON(http.StatusOK, &models.ResponseGetUsers{
		MessageId: resp.MessageId,
		Users:     users,
	})
}
