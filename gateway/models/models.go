package models

import api "WP/biz/api/proto"

type RequestReqPQ struct {
	Nonce     string `json:"nonce"`
	MessageId int32  `json:"message_id"`
}

type ResponseReqPQ struct {
	Nonce       string `json:"nonce"`
	ServerNonce string `json:"server_nonce"`
	MessageId   int32  `json:"message_id"`
	P           int32  `json:"p"`
	G           int32  `json:"g"`
}

type User struct {
	Name   string `json:"name"`
	Family string `json:"family"`
	Age    int32  `json:"age"`
	Id     int32  `json:"id"`
	Sex    string `json:"sex"`
}

type RequestReqDHParams struct {
	Nonce       string `json:"nonce"`
	ServerNonce string `json:"server_nonce"`
	MessageId   int32  `json:"message_id"`
	A           int32  `json:"a"`
}

type ResponseReqDHParams struct {
	Nonce       string `json:"nonce"`
	ServerNonce string `json:"server_nonce"`
	MessageId   int32  `json:"message_id"`
	B           int32  `json:"b"`
}

//--------------------------------------------
// --------------- BIZ Models ----------------
// -------------------------------------------

type RequestGetUsers struct {
	UserId    int32 `json:"user_id"`
	MessageId int32 `json:"message_id"`
	AuthKey   int32 `json:"auth_key"`
}

type ResponseGetUsers struct {
	MessageId int32   `json:"message_id"`
	Users     []*User `json:"users"`
}

type RequestGetUsersWithInj struct {
	UserId    string `json:"user_id"`
	MessageId int32  `json:"message_id"`
	AuthKey   int32  `json:"auth_key"`
}

type ResponseGetUsersWithInj struct {
	MessageId int32   `json:"message_id"`
	Users     []*User `json:"users"`
}

func UserProtoToModel(user *api.User) *User {
	var sex string
	if user.Sex == api.UserSex_UserSex_FEMALE {
		sex = "FEMALE"
	} else {
		sex = "MALE"
	}
	return &User{
		Name:   user.Name,
		Family: user.Family,
		Sex:    sex,
		Age:    user.Age,
		Id:     user.Id,
	}
}
