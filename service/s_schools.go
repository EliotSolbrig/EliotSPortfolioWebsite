package service

import (
    "context"
    "main/database"
)

func (service *Service) GetAllSchools(ctx context.Context) ([]*database.School, error) {
    schools,err := service.Database.GetAllSchools(ctx)
    if err != nil {
        return []*database.School{}, err
    }
    return schools, nil
}

func (service *Service) GetSchoolClassesFromSchoolID(ctx context.Context, schoolID int) ([]database.Class, error) {
    classes,err := service.GetSchoolClassesFromSchoolID(ctx, schoolID)
    if err != nil {
        return []database.Class{}, err
    }
    return classes, nil
}

