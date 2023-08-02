package repository

import (
	"web-api/utils/database"
)

var Repo MysqlRepository

type MySqlRepositoryRepo struct{}

func MySqlInit() {
	Repo = &MySqlRepositoryRepo{}
}

/***Inserting data to database***/
func (r *MySqlRepositoryRepo) Insert(req interface{}) error {
	if err := database.DB.Debug().Create(req).Error; err != nil {
		return err
	}
	return nil
}

/***Fetching data from database***/
func (r *MySqlRepositoryRepo) FindAll(req interface{}, condition interface{}) error {
	if err := database.DB.Debug().Where(condition).Find(req).Error; err != nil {
		return err
	}
	return nil
}
