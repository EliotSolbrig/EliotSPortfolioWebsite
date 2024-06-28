package service

import (
    "context"
    "main/database"
)

func (service *Service) GetAllProjects(ctx context.Context) ([]*database.Project, error) {
   projects, err := service.Database.GetAllProjects(ctx)
   if err != nil {
       return []*database.Project{}, err
   }
   return projects, nil
    
}
