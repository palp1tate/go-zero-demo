// Code generated by goctl. DO NOT EDIT.
package types

type Base struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type LoginRequest struct {
	Mobile   string `form:"mobile"`
	Password string `form:"password"`
}

type LoginResponse struct {
	Base
	Data UserInfo `json:"data"`
}

type RegisterRequest struct {
	Name     string `form:"name"`
	Mobile   string `form:"mobile"`
	Password string `form:"password"`
}

type RegisterResponse struct {
	Base
	Data UserInfo `json:"data"`
}

type UserInfo struct {
	UserID int64  `json:"uid"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}
