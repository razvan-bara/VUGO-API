package services

import (
	"context"
	"database/sql"
	"github.com/go-openapi/swag"
	"github.com/razvan-bara/VUGO-API/api/sdto"
	db "github.com/razvan-bara/VUGO-API/db/sqlc"
)

type IQuestionService interface {
	SaveQuestion(ctx context.Context, quizID int64, question *sdto.QuizFormQuestionsItems0) (*db.Question, error)
}

type QuestionService struct {
	storage db.Storage
}

func NewQuestionService(storage db.Storage) *QuestionService {
	return &QuestionService{storage: storage}
}

func (qs *QuestionService) SaveQuestion(ctx context.Context, quizID int64, question *sdto.QuizFormQuestionsItems0) (*db.Question, error) {
	questionArgs := &db.CreateQuestionParams{
		QuizID: quizID,
		Title:  swag.StringValue(question.Title),
		Body: sql.NullString{
			String: question.Body,
			Valid:  false,
		},
	}

	if question.Body != "" {
		questionArgs.Body.Valid = true
	}

	savedQuestion, err := qs.storage.CreateQuestion(ctx, questionArgs)
	if err != nil {
		return nil, err
	}

	return savedQuestion, err
}
