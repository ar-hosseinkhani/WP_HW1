package main

import (
	api "WP/biz/api/proto"
	"WP/biz/models"
)

func SexToProto(sex string) api.UserSex {
	if sex == "MALE" {
		return api.UserSex_UserSex_MALE
	} else {
		return api.UserSex_UserSex_FEMALE
	}
}

func UserToProto(user *models.User) *api.User {
	return &api.User{
		Name:   user.Name,
		Sex:    SexToProto(user.Sex),
		Age:    user.Age,
		Family: user.Family,
		Id:     user.Id,
	}
}
