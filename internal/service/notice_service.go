package service

import (
	"time"
	"github.com/MananLed/upKeepz-cli/internal/model"
	"github.com/MananLed/upKeepz-cli/internal/repository"
)

type NoticeService struct {
	NoticeRepo repository.NoticeRepositoryInterface
}

func NewNoticeService(repo repository.NoticeRepositoryInterface) *NoticeService {
	return &NoticeService{NoticeRepo: repo}
}

func (s *NoticeService) IssueNotice(content string) error {
	notice := model.Notice{
		DateIssued: time.Now().Format("2006-01-02 15:04:05"),
		Content:    content,
	}
	return s.NoticeRepo.SaveNotice(notice)
}

func (s *NoticeService) GetNotices() ([]model.Notice, error) {
	return s.NoticeRepo.GetAllNotices()
}

func (s *NoticeService) GetNoticeByID(id int) (*model.Notice, error) {	
	return s.NoticeRepo.GetNoticeByID(id)
}