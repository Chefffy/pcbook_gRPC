package client

import (
	"context"
	"google.golang.org/grpc"
	"pcbook-gRPC/pb"
	"time"
)

type AuthClient struct {
	service pb.AuthServiceClient
	username string
	password string
}

func NewAuthClient(cc *grpc.ClientConn,username string,password string) *AuthClient{
	service := pb.NewAuthServiceClient(cc)
	return &AuthClient{service,username,password}
}

// Login 用户登录返回令牌
func (client *AuthClient) Login()(string,error){
	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	req := &pb.LoginRequest{
		Username: client.username,
		Password: client.password,
	}

	res , err := client.service.Login(ctx,req)
	if err != nil{
		return "",err
	}

	return res.GetAccessToken(),nil
}