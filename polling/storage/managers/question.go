package managers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	pb "poll-service/genprotos"
)

type QuestionManager struct {
	Conn *sql.DB
}

func NewQuestionManager(conn *sql.DB) *QuestionManager {
	return &QuestionManager{Conn: conn}
}

func (m *QuestionManager) Create(ctx context.Context, question *pb.QuestionCreateReq) (*pb.Void, error) {
	query := "SELECT insert_question($1, $2)"
	_, err := m.Conn.ExecContext(ctx, query, question.Content, question.PollId)
	if err != nil {
		return nil, fmt.Errorf("failed to create question: %w", err)
	}
	return &pb.Void{}, nil
}

func (m *QuestionManager) GetByID(ctx context.Context, req *pb.ByID) (*pb.QuestionGetByIDRes, error) {
	query := "SELECT id, content, poll_id FROM questions WHERE id = $1"
	row := m.Conn.QueryRowContext(ctx, query, req.Id)

	var res pb.QuestionGetByIDRes
	err := row.Scan(&res.Id, &res.Content, &res.PollId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("question not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get question by id: %w", err)
	}
	return &res, nil
}

func (m *QuestionManager) Update(ctx context.Context, question *pb.QuestionUpdateReq) (*pb.Void, error) {
	query := "UPDATE questions SET content = $1 WHERE id = $2"
	_, err := m.Conn.ExecContext(ctx, query, question.Content, question.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to update question: %w", err)
	}
	return &pb.Void{}, nil
}

func (m *QuestionManager) Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	query := "DELETE FROM questions WHERE id = $1"
	_, err := m.Conn.ExecContext(ctx, query, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete question: %w", err)
	}
	return &pb.Void{}, nil
}

func (m *QuestionManager) GetAll(ctx context.Context, req *pb.QuestionGetAllReq) (*pb.QuestionGetAllRes, error) {
	questionQuery := "SELECT id, num, content, poll_id FROM questions WHERE poll_id = $1"
	rows, err := m.Conn.QueryContext(ctx, questionQuery, req.PollId)
	if err != nil {
		return nil, fmt.Errorf("failed to get all questions: %w", err)
	}
	defer rows.Close()

	var questions []*pb.Question
	for rows.Next() {
		var question pb.Question
		err := rows.Scan(&question.Id, &question.Num, &question.Content, &question.PollId)
		if err != nil {
			return nil, fmt.Errorf("failed to scan question: %w", err)
		}
		questions = append(questions, &question)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	pollQuery := "SELECT id, poll_num, title, options FROM polls WHERE id = $1"
	var poll pb.PollGetByIDRes
	var pollNum int32
	var options, title, pollId string
	var optionList []*pb.Option

	if options != "" {
		if err := json.Unmarshal([]byte(options), &optionList); err != nil {
			return nil, fmt.Errorf("failed to unmarshal options: %s", err.Error())
		}
	}

	err = m.Conn.QueryRowContext(ctx, pollQuery, req.PollId).Scan(&pollId, &pollNum, &title, &options)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.QuestionGetAllRes{
				Question: []*pb.Question{},
				Poll:     nil,
			}, nil
		}
		return nil, fmt.Errorf("failed to get poll details: %w", err)
	}

	poll.Id = pollId
	poll.PollNum = pollNum
	poll.Title = title
	poll.Options = optionList

	return &pb.QuestionGetAllRes{
		Question: questions,
		Poll:     &poll,
	}, nil
}
