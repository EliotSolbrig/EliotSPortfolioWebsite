package service

import (
    "main/database"
)

type IService interface {

}

type Service struct {
    Database database.IDatabase
}

func NewService(database database.IDatabase) *Service {
    return &Service{
        Database: database,
    }
}
