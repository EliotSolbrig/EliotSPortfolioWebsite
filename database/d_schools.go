package database

import (
    "fmt"
    "context"
)

const getAllSchoolsQuery = `
SELECT
    id,
    name,
    major,
    graduated,
    degree,
    school_name,
    school_type_id,
    gpa,
    location,
    transcript_path,
    logo_path
FROM
    schools
ORDER BY id DESC
`

func (database *Database) GetAllSchools(ctx context.Context) ([]*School, error) {
    res,err := database.db.QueryContext(ctx, getAllSchoolsQuery)
    if err != nil {
        return []*School{}, fmt.Errorf("Error querying getAllSchoolsQuery in d_schools.go: %s", err)
    }
    schools := []*School{}
    for res.Next(){
        school := School{}
        res.Scan(
            &school.ID,
            &school.Name,
            &school.Major,
            &school.Graduated,
            &school.Degree,
            &school.SchoolName,
            &school.SchoolTypeID,
            &school.GPA,
            &school.Location,
            &school.TranscriptPath,
            &school.LogoPath,
        )
        classes,_ := database.GetSchoolClassesFromSchoolID(ctx, school.ID)
        school.Classes = classes
        schools = append(schools, &school)
    }
    return schools, nil
}

const getSchoolClassesFromSchoolIDQuery = `
SELECT
    A.id,
    A.school_id,
    A.abbreviation,
    A.name,
    A.grade,
    A.term,
    A.credits,
    class_files_url
FROM
    classes A,
    schools B
WHERE
    A.school_id = B.id AND
    B.id = ?
ORDER BY A.id DESC
`

func (database *Database) GetSchoolClassesFromSchoolID(ctx context.Context, schoolID int) ([]Class, error) {
    res,err := database.db.QueryContext(ctx, getSchoolClassesFromSchoolIDQuery, schoolID)
    if err != nil {
        return []Class{}, fmt.Errorf("Error executing getSchoolClassesFromSchoolIDQuery in d_schools.go: %s", err)
    }
    classes := []Class{}
    for res.Next(){
        class := Class{}
        res.Scan(
            &class.ID,
            &class.SchoolID,
            &class.Abbreviation,
            &class.Name,
            &class.Grade,
            &class.Term,
            &class.Credits,
            &class.ClassFilesURL,
        )
        classes = append(classes, class)
    }
    return classes, nil
}
