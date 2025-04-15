package models

import "seotrang.com/rgpc-microservice-user/userpb"

type Server struct {
	userpb.UnimplementedUserServiceServer
}
