package repository

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/MananLed/upKeepz-cli/constants"
	"github.com/MananLed/upKeepz-cli/internal/model"
)

type FeedbackRepository struct {
	mu sync.Mutex
}

type FeedbackRepositoryInterface interface {
	SaveFeedback(model.Feedback) error
	GetFeedbacksByID(string) ([]model.Feedback, error)
	GetAllFeedbacks() ([]model.Feedback, error)
}

func (f *FeedbackRepository) loadFeedbacks() ([]model.Feedback, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	fileData, err := os.ReadFile(string(constants.FeedbackDataPath))
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Feedback{}, nil
		}
		return nil, err
	}

	var feedbacks []model.Feedback
	if err := json.Unmarshal(fileData, &feedbacks); err != nil {
		return nil, err
	}
	return feedbacks, nil
}

func (f *FeedbackRepository) GetFeedbacksByID(id string) ([]model.Feedback, error) {
	feedbacks, err := f.loadFeedbacks()
	if err != nil {
		return nil, err
	}
	var feedbackOfResident []model.Feedback
	for _, f := range feedbacks {
		if f.ResidentID == id {
			feedbackOfResident = append(feedbackOfResident, f) 
		}
	}
	return feedbackOfResident, nil
}

func (f *FeedbackRepository) GetAllFeedbacks() ([]model.Feedback, error){
	return f.loadFeedbacks()
}

func (f *FeedbackRepository) SaveFeedback(feedback model.Feedback) error {
	feedbacks, err := f.loadFeedbacks()

	if err != nil{
		return err 
	}

	feedbacks = append(feedbacks, feedback)

	f.mu.Lock()
	defer f.mu.Unlock()

	data, err := json.MarshalIndent(feedbacks, "", " ")
	if err != nil{
		return err 
	}

	return os.WriteFile(string(constants.FeedbackDataPath), data, 0644)
}
