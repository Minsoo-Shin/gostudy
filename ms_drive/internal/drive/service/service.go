package service

import "github.com/Minsoo-Shin/ms_drive/internal/drive/repository"

type DriveService interface {
}

type driveService struct {
	repository.DriverRepository
}

func New() {

}
