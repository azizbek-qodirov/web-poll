package service

import (
	"context"
	pb "poll-service/genprotos"

	st "poll-service/storage"
)

type ResultService struct {
	storage st.StorageI
	pb.UnimplementedResultServiceServer
}

func NewResultService(storage st.StorageI) *ResultService {
	return &ResultService{storage: storage}
}

func (s *ResultService) CreateResult(ctx context.Context, req *pb.CreateResultReq) (*pb.CreateResultRes, error) {
	resAnswer, err := s.storage.Result().GetByIDRes(ctx, &pb.ByID{Id: req.PollId})
	if err != nil {
		return nil, err
	}
	var total int32
	var feed string
	for _, v := range resAnswer.Answers {
		total += v.AnswerPoint
	}
	poll, err := s.storage.Poll().GetByID(ctx, &pb.ByID{Id: req.PollId})
	if err != nil {
		return nil, err
	}

	for _, v := range poll.Feedback {
		if total >= v.From && total < v.To {
			feed = v.Text
			break
		}
	}
	res, err := s.storage.Result().CreateResult(ctx, req)
	if err != nil {
		return nil, err
	}
	res.Feed = feed
	return res, nil
}

func (s *ResultService) SavePollAnswer(ctx context.Context, req *pb.SavePollAnswerReq) (*pb.Void, error) {
	return s.storage.Result().SavePollAnswer(ctx, req)
}

func (s *ResultService) GetResultsInExcel(ctx context.Context, req *pb.Void) (*pb.ExcelResultsRes, error) {
	return s.storage.Result().GetResultsInExcel(ctx, req)
}
