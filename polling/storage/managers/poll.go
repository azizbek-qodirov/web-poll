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
	query := "SELECT insert_poll($1, $2, $3)"
	optionsJSON, err := json.Marshal(poll.Options)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal options to JSON: %w", err)
	}
	feedbackJSON, err := json.Marshal(poll.Feedbacks)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal feedback to JSON: %w", err)
	}

	_, err = m.Conn.ExecContext(ctx, query, poll.Title, optionsJSON, feedbackJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to create poll: %w", err)
	}
	return &pb.Void{}, nil
}


func (m *PollManager) Update(ctx context.Context, poll *pb.PollUpdateReq) (*pb.Void, error) {
	// Pollni yangilash uchun so'rov
	query := "UPDATE polls SET title = $1 WHERE id = $2"

	// So'rov uchun argumentlarni tayyorlash
	args := []interface{}{poll.Title, poll.Id}

	// So'rovni bajarish
	_, err := m.Conn.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to update poll: %w", err)
	}

	return &pb.Void{}, nil
}


func (m *PollManager) GetByID(ctx context.Context, req *pb.ByID) (*pb.PollGetByIDRes, error) {
	query := "SELECT id, poll_num, title, options, feedbacks FROM polls WHERE id = $1"
	row := m.Conn.QueryRowContext(ctx, query, req.Id)

	var (
		id        string
		pollNum   int32
		title     string
		options   string
		feedbacks string
	)

	err := row.Scan(&id, &pollNum, &title, &options, &feedbacks)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("poll not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get poll by id: %w", err)
	}

	// Optionsni map orqali unmarshaling qilish
	var optionList []*pb.Option
	if options != "" {
		var rawOptions map[string]interface{}
		if err := json.Unmarshal([]byte(options), &rawOptions); err != nil {
			return nil, fmt.Errorf("failed to unmarshal options: %w", err)
		}

		// Agar sizda bir nechta variant bo'lsa, ularni to'g'ri konvertatsiya qilish
		for key, value := range rawOptions {
			option := &pb.Option{}
			if key == "variant" {
				if variant, ok := value.(string); ok {
					option.Variant = &variant
				}
			} else if key == "ball" {
				if ball, ok := value.(float64); ok {
					ballInt := int32(ball)
					option.Ball = &ballInt
				}
			}
			optionList = append(optionList, option)
		}
	}

	// Feedbacksni map orqali unmarshaling qilish
	var feedbackList []*pb.Feedback
	if feedbacks != "" {
		var rawFeedbacks map[string]interface{}
		if err := json.Unmarshal([]byte(feedbacks), &rawFeedbacks); err != nil {
			return nil, fmt.Errorf("failed to unmarshal feedbacks: %w", err)
		}

		for key, value := range rawFeedbacks {
			feedback := &pb.Feedback{}
			if key == "from" {
				if from, ok := value.(float64); ok {
					fromInt := int32(from)
					feedback.From = &fromInt
				}
			} else if key == "to" {
				if to, ok := value.(float64); ok {
					toInt := int32(to)
					feedback.To = &toInt
				}
			} else if key == "text" {
				if text, ok := value.(string); ok {
					feedback.Text = &text
				}
			}
			feedbackList = append(feedbackList, feedback)
		}
	}

	return &pb.PollGetByIDRes{
		Id:       &id,
		PollNum:  &pollNum,
		Title:    &title,
		Options:  optionList,
		Feedback: feedbackList,
	}, nil
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
	query := "SELECT id, poll_num, title, options, feedbacks FROM polls WHERE 1 = 1"
	var args []interface{}
	paramIndex := 1

	if req.UserId != nil {
		query += fmt.Sprintf(" AND user_id = $%d", paramIndex)
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
		var (
			id        string
			pollNum   int32
			title     string
			options   []byte
			feedbacks []byte
		)

		err := rows.Scan(&id, &pollNum, &title, &options, &feedbacks)
		if err != nil {
			return nil, fmt.Errorf("failed to scan poll: %w", err)
		}

		// Optionsni JSON ob'ekti sifatida unmarshaling qilish
		var optionList []*pb.Option
		if len(options) > 0 {
			var rawOptions []map[string]interface{}
			if err := json.Unmarshal(options, &rawOptions); err != nil {
				return nil, fmt.Errorf("failed to unmarshal options: %w", err)
			}

			for _, rawOption := range rawOptions {
				option := &pb.Option{}
				if variant, ok := rawOption["variant"].(string); ok {
					option.Variant = &variant
				}
				if ball, ok := rawOption["ball"].(float64); ok {
					ballInt := int32(ball)
					option.Ball = &ballInt
				}
				optionList = append(optionList, option)
			}
		}

		// Feedbacksni JSON ob'ekti sifatida unmarshaling qilish
		var feedbackList []*pb.Feedback
		if len(feedbacks) > 0 {
			var rawFeedbacks []map[string]interface{}
			if err := json.Unmarshal(feedbacks, &rawFeedbacks); err != nil {
				return nil, fmt.Errorf("failed to unmarshal feedbacks: %w", err)
			}

			for _, rawFeedback := range rawFeedbacks {
				feedback := &pb.Feedback{}
				if from, ok := rawFeedback["from"].(float64); ok {
					fromInt := int32(from)
					feedback.From = &fromInt
				}
				if to, ok := rawFeedback["to"].(float64); ok {
					toInt := int32(to)
					feedback.To = &toInt
				}
				if text, ok := rawFeedback["text"].(string); ok {
					feedback.Text = &text
				}
				feedbackList = append(feedbackList, feedback)
			}
		}

		polls = append(polls, &pb.PollGetByIDRes{
			Id:       &id,
			PollNum:  &pollNum,
			Title:    &title,
			Options:  optionList,
			Feedback: feedbackList,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return &pb.PollGetAllRes{
		Poll: polls,
	}, nil
}






