package database

import (
    "fmt"
    "context"
)

const getAllProjectsQuery = `
SELECT
    id,
    name,
    url,
    github_repo_url,
    desc
FROM
    projects
ORDER BY
    id DESC
`

func (database *Database) GetAllProjects(ctx context.Context) ([]*Project, error) {
    projects := []*Project{}
    res,err := database.db.QueryContext(ctx, getAllProjectsQuery)
    if err != nil {
        return []*Project{}, fmt.Errorf("Error querying getAllProjectsQuery in d_projects.go: %s", err)
    }
    for res.Next(){
        project := Project{}
        res.Scan(
            &project.ID,
            &project.Name,
            &project.URL,
            &project.GithubRepoURL,
            &project.Description,
        )
        projects = append(projects, &project)
    }
    return projects, nil
}

const getProjectImagesFromIDQuery = `
SELECT
    A.id,
    A.project_id,
    A.image_link
FROM
    project_images A,
    projects B
WHERE
   B.id = A.project_id AND
   B.id = ?
ORDER BY
    A.id DESC
`

func (database *Database) GetProjectImagesFromID(ctx context.Context, projectID int) ([]ProjectImage, error) {
    images := []ProjectImage{}
    res,err := database.db.QueryContext(ctx, getProjectImagesFromIDQuery, projectID)
    if err != nil {
        return images, fmt.Errorf("Error querying getProjectImagesFromIDQuery in d_projects.go: %s", err)
    }
    for res.Next(){
        image := ProjectImage{}
        res.Scan(
            &image.ID,
            &image.ProjectID,
            &image.ImageLink,
        )
        images = append(images, image)
    }
    return images, nil
}
