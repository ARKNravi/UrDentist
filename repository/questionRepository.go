package repository

import (
	"github.com/ARKNravi/HACKFEST-BE/database"
	"github.com/ARKNravi/HACKFEST-BE/model"
)

type QuestionRepository struct {}

func NewQuestionRepository() *QuestionRepository {
	return &QuestionRepository{}
}

func (r *QuestionRepository) CreateQuestion(question *model.Question) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	result := db.Create(question)
	return result.Error
}

func (r *QuestionRepository) GetAllQuestions() ([]model.Question, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	var questions []model.Question
	result := db.Find(&questions)
	return questions, result.Error
}

func (r *QuestionRepository) GetQuestionByID(id uint) (*model.Question, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	var question model.Question
	result := db.First(&question, id)
	return &question, result.Error
}