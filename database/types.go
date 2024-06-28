package database

type Job struct {
    ID int
    Title string
    CompanyName string
    Location string
    StartDate string
    EndDate string
    ExperienceItems []ExperienceItem
}

type ExperienceItem struct {
    ID int
    JobID int
    Text string
}

