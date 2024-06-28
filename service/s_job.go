package service

import (
    "context"
    "main/database"
)

func (service *Service) GetAllJobs(ctx context.Context) ([]database.Job, error) {
    jobs,err := service.Database.GetAllJobs(ctx)
    if err != nil {
        return []database.Job{}, nil
    }
    return jobs, nil
}

func (service *Service) GetExperienceItemsByID(ctx context.Context, jobID int) ([]database.ExperienceItem, error) {
    items,err := service.Database.GetExperienceItemsByID(ctx, jobID)
    if err != nil {
        return []database.ExperienceItem{}, err
    }
    return items, nil
}
