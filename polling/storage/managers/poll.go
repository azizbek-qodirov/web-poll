package managers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	pb "poll-service/genprotos"
)

type PollManager struct {
	Conn *sql.DB
}

func NewPollManager(conn *sql.DB) *PollManager {
	return &PollManager{Conn: conn}
}

func (m *PollManager) Create(ctx context.Context, poll *pb.PollCreateReq) (*pb.Void, error) {
	query := "INSERT INTO polls (title, options) VALUES ($1, $2)"
	optionsJSON, err := json.Marshal(poll.Options)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal options to JSON: %w", err)
	}

	_, err = m.Conn.ExecContext(ctx, query, poll.Title, optionsJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to create poll: %w", err)
	}
	return &pb.Void{}, nil
}

func (m *PollManager) GetByID(ctx context.Context, req *pb.ByID) (*pb.PollGetByIDRes, error) {
	query := "SELECT id, title, options FROM polls WHERE id = $1"
	row := m.Conn.QueryRowContext(ctx, query, req.Id)

	var res pb.PollGetByIDRes
	err := row.Scan(&res.Id, &res.PollNum, &res.Title, &res.Options)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("poll not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get poll by id: %w", err)
	}
	return &res, nil
}

func (m *PollManager) Update(ctx context.Context, poll *pb.PollUpdateReq) (*pb.Void, error) {
	query := "UPDATE polls SET  title = $1 WHERE id = $2"
	_, err := m.Conn.ExecContext(ctx, query, poll.Title, poll.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to update poll: %w", err)
	}
	return &pb.Void{}, nil
}

func (m *PollManager) Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	query := "DELETE FROM polls WHERE id = $1"
	_, err := m.Conn.ExecContext(ctx, query, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete poll: %w", err)
	}
	return &pb.Void{}, nil
}

func (m *PollManager) GetAll(ctx context.Context, req *pb.PollGetAllReq) (*pb.PollGetAllRes, error) {
	query := "SELECT id, poll_num, title, options FROM polls WHERE 1 = 1"
	var args []interface{}
	paramIndex := 1

	if req.UserId != "" {
		query += fmt.Sprintf(" AND user_id = %d", paramIndex)
		args = append(args, req.UserId)
		paramIndex++
	}

	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get all polls: %w", err)
	}
	defer rows.Close()

	var polls []*pb.PollGetByIDRes
	for rows.Next() {
		var poll pb.PollGetByIDRes
		err := rows.Scan(&poll.Id, &poll.PollNum, &poll.Title, &poll.Options)
		if err != nil {
			return nil, fmt.Errorf("failed to scan poll: %w", err)
		}
		polls = append(polls, &poll)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return &pb.PollGetAllRes{
		Poll: polls,
	}, nil
}
