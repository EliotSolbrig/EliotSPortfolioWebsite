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

func (service *Service) GetProjectImagesFromID(ctx context.Context, projectID int) ([]database.ProjectImage, error) {
    images,err := service.Database.GetProjectImagesFromID(ctx, projectID)
    if err != nil {
        return []database.ProjectImage{}, err
    }
    return images, nil


}
