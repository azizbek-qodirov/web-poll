package managers

import (
	"context"
	"database/sql"

	pb "poll-service/genprotos"

	"github.com/google/uuid"
)

type ResultManager struct {
	Conn *sql.DB
}

func NewResultManager(conn *sql.DB) *ResultManager {
	return &ResultManager{Conn: conn}
}

func (m *ResultManager) CreateResult(ctx context.Context, req *pb.CreateResultReq) (*pb.Void, error) {
	query := `INSERT INTO results (id, user_id, poll_id) VALUES ($1, $2, $3)`
	_, err := m.Conn.ExecContext(ctx, query, uuid.NewString(), req.UserId, req.PollId)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (m *ResultManager) SavePollAnswer(ctx context.Context, req *pb.SavePollAnswerReq) (*pb.Void, error) {
	query := "INSERT INTO poll_answers (id, result_id, question_num, answer) VALUES ($1, $2, $3, $4)"
	_, err := m.Conn.ExecContext(ctx, query, uuid.NewString(), req.ResultId, req.QuestionNum, req.Answer)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
