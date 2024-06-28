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
