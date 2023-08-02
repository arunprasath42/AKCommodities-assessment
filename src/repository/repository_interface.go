package repository

type MysqlRepository interface {
	Insert(req interface{}) error
	FindAll(req interface{}, condition interface{}) error
}
