package managers

import (
	"context"
	"database/sql"
	"fmt"

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

func (m *ResultManager) GetPoll(ctx context.Context, req *pb.GetPollReq) (*pb.GetPollRes, error) {
	var userID string
	query := `SELECT user_id FROM results WHERE user_id = $1`
	row := m.Conn.QueryRowContext(ctx, query, req.UserId)
	err := row.Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no poll found for user_id: %s", req.UserId)
		}
		return nil, err
	}

	pollQuery := `SELECT poll_id FROM results WHERE user_id = $1`
	pollRows, err := m.Conn.QueryContext(ctx, pollQuery, req.UserId)
	if err != nil {
		return nil, err
	}
	defer pollRows.Close()

	var polls []*pb.GetPoll
	for pollRows.Next() {
		var pollID string
		err = pollRows.Scan(&pollID)
		if err != nil {
			return nil, err
		}

		polls = append(polls, &pb.GetPoll{
			PollId: pollID,
		})
	}

	return &pb.GetPollRes{
		UserId: req.UserId,
		Poll:   polls,
	}, nil
}
