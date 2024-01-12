package repository

import (
	"time"

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

func (r *QuestionRepository) AnswerQuestion(id uint, answer string, dentistID uint) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	var question model.Question
	result := db.First(&question, id)
	if result.Error != nil {
		return result.Error
	}
	now := time.Now()
	question.Answer = answer
	question.DentistID = &dentistID
	question.AnsweredAt = &now
	result = db.Model(&question).Updates(model.Question{Answer: answer, DentistID: &dentistID, AnsweredAt: &now})
	return result.Error
}