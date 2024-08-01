package handlers

import (
	pb "auth-service/genprotos"

	"google.golang.org/grpc"
)

type HTTPHandler struct {
	User     pb.UserServiceClient
	Poll     pb.PollServiceClient
	Question pb.QuestionServiceClient
}

func NewHandler(conn *grpc.ClientConn) *HTTPHandler {
	return &HTTPHandler{
		User:     pb.NewUserServiceClient(conn),
		Poll:     pb.NewPollServiceClient(conn),
		Question: pb.NewQuestionServiceClient(conn),
	}
}
