package service

import (
    "context"
    "main/database"
)

type IService interface {
    GetAllJobs(ctx context.Context) ([]database.Job, error)
    GetExperienceItemsByID(ctx context.Context, jobID int) ([]database.ExperienceItem, error)
}

type Service struct {
    Database database.IDatabase
}

func NewService(database database.IDatabase) *Service {
    return &Service{
        Database: database,
    }
}
