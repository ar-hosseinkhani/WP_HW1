package models

type RequestReqPQ struct {
	Nonce     string `json:"nonce"`
	MessageId int32  `json:"message_id"`
}

type ResponseReqPQ struct {
	Nonce       string `json:"nonce"`
	ServerNonce string `json:"server_nonce"`
	MessageId   int32  `json:"message_id"`
	P           int64  `json:"p"`
	G           int64  `json:"g"`
}

type RequestReqDHParams struct {
	Nonce       string `json:"nonce"`
	ServerNonce string `json:"server_nonce"`
	MessageId   int32  `json:"message_id"`
	A           int64  `json:"a"`
}

type ResponseReqDHParams struct {
	Nonce       string `json:"nonce"`
	ServerNonce string `json:"server_nonce"`
	MessageId   int32  `json:"message_id"`
	B           int64  `json:"b"`
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
	MessageId int32 `json:"message_id"`
	// array of user
}

type RequestGetUsersWithInj struct {
	UserId    string `json:"user_id"`
	MessageId int32  `json:"message_id"`
	AuthKey   int32  `json:"auth_key"`
}

type ResponseGetUsersWithInj struct {
	MessageId int32 `json:"message_id"`
	// array of user
}
