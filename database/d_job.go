package database

import (
    "fmt"
    "context"
)

const getAllJobsQuery = `
SELECT
    id,
    title,
    company_name,
    location,
    start_date,
    end_date
FROM
    jobs
ORDER BY
    id ASC
`

func (database *Database) GetAllJobs(ctx context.Context) ([]Job, error) {
    jobs := []Job{}
    res,err := database.db.QueryContext(ctx, getAllJobsQuery)
    if err != nil {
        return []Job{}, fmt.Errorf("Error executing getAllJobsQuery in d_jobs.go: %s", err)
    }
    for res.Next(){
        job := Job{}
        res.Scan(
            &job.ID,
            &job.Title,
            &job.CompanyName,
            &job.Location,
            &job.StartDate,
            &job.EndDate,
        )
        jobExperienceItems,err := database.GetExperienceItemsByID(ctx, job.ID)
        if err != nil {
            fmt.Println(fmt.Errorf("Error getting job experience items for job %s: %s", job, err))
        }
        job.ExperienceItems = jobExperienceItems
        fmt.Println("job: ", job)
        jobs = append(jobs, job)
    }
    return jobs, nil
}

const getJobExperienceItemsByIDQuery = `
SELECT
    B.id,
    B.job_id,
    B.text
FROM
    jobs A,
    experience_items B
WHERE
    A.id = B.Job_id AND 
    B.job_id = ?
ORDER BY
    B.id DESC
`

func (database *Database) GetExperienceItemsByID(ctx context.Context, jobID int) ([]ExperienceItem, error) {
    items := []ExperienceItem{}
    res,err := database.db.QueryContext(ctx, getJobExperienceItemsByIDQuery, jobID)
    if err != nil {
        return items, fmt.Errorf("Error querying getJobExperienceItemsByIDQuery in d_job.go: %s", err)
    }
    for res.Next(){
        item := ExperienceItem{}
        res.Scan(
            &item.ID,
            &item.JobID,
            &item.Text,
        )
        items = append(items, item)
    }
    return items, nil
}
