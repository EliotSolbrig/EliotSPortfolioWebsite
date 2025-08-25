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
    desc,
    project_type_id,
    disabled,
    list_order,
	description_html
FROM
    projects
WHERE
   project_type_id LIKE ? 
ORDER BY
    list_order ASC
`

func (database *Database) GetAllProjects(ctx context.Context, typeID int) ([]*Project, error) {
    var typeString string = "%%"
    if typeID != 0 {
        typeString = fmt.Sprintf("%d",typeID)
    }
    fmt.Println("typeString: ", typeString)
    projects := []*Project{}
    res,err := database.db.QueryContext(ctx, getAllProjectsQuery, typeString)
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
            &project.ProjectType.ID,
            &project.Disabled,
            &project.ListOrder, 
            &project.DescriptionHTML, 
        )
        images,err := database.GetProjectImagesFromID(ctx, project.ID)
        if err != nil {
            fmt.Println(fmt.Errorf("Error getting project images from id for %s: %s", project, err))
        }
        project.ProjectImages = images

        projectType,err := database.GetProjectTypeFromID(ctx, project.ProjectType.ID)
        if err != nil {
            fmt.Println(fmt.Errorf("Error getting project type from id for %s: %s", project, err))
        }
        project.ProjectType = *projectType

        projects = append(projects, &project)
    }
    return projects, nil
}

const getProjectImagesFromIDQuery = `
SELECT
    A.id,
    A.project_id,
    A.image_link,
    A.caption,
    A.alt_link,
    A.image_is_video
FROM
    project_images A,
    projects B
WHERE
   B.id = A.project_id AND
   B.id = ?
ORDER BY
    A.id ASC
`

func (database *Database) GetProjectImagesFromID(ctx context.Context, projectID int) ([]*ProjectImage, error) {
    images := []*ProjectImage{}
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
            &image.Caption,
            &image.AltLink,
            &image.ImageIsVideo,
        )
        images = append(images, &image)
    }
    return images, nil
}

const getProjectTypeFromIDQuery = `
SELECT
    id,
    name,
    desc
FROM
    project_type
WHERE
    id = ?
`

func (database *Database) GetProjectTypeFromID(ctx context.Context, typeID int) (*ProjectType, error) {
    res,err := database.db.QueryContext(ctx, getProjectTypeFromIDQuery, typeID)
    if err != nil {
        return &ProjectType{}, fmt.Errorf("Error querying getProjectTypeFromIDQuery in d_projects.go: %s", err)
    }
    res.Next()
    projectType := ProjectType{}
    res.Scan(
        &projectType.ID,
        &projectType.Name,
        &projectType.Description,
    )
    return &projectType, nil
}
