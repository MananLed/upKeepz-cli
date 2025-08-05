package repository

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/MananLed/upKeepz-cli/constants"
	"github.com/MananLed/upKeepz-cli/internal/model"
)

type NoticeRepositoryInterface interface {
	SaveNotice(notice model.Notice) error
	GetAllNotices() ([]model.Notice, error)
	GetNoticeByID(id int) (*model.Notice, error)
}

type NoticeRepository struct {
	mu sync.Mutex
}

func (r *NoticeRepository) loadNotices() ([]model.Notice, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	fileData, err := os.ReadFile(string(constants.NoticeDataPath))
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Notice{}, nil
		}
		return nil, err
	}

	var notices []model.Notice
	if err := json.Unmarshal(fileData, &notices); err != nil {
		return nil, err
	}

	return notices, nil
}

func (r *NoticeRepository) SaveNotice(notice model.Notice) error {
	notices, err := r.loadNotices()
	if err != nil {
		return err
	}

	notice.ID = len(notices) + 1
	notices = append(notices, notice)

	r.mu.Lock()
	defer r.mu.Unlock()
	data, err := json.MarshalIndent(notices, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(string(constants.NoticeDataPath), data, 0644)
}

func (r *NoticeRepository) GetAllNotices() ([]model.Notice, error) {
	return r.loadNotices()
}

func (r *NoticeRepository) GetNoticeByID(id int) (*model.Notice, error) {
	notices, err := r.loadNotices()
	if err != nil {
		return nil, err
	}

	for _, n := range notices {
		if n.ID == id {
			return &n, nil
		}
	}

	return nil, errors.New("notice not found")
}
