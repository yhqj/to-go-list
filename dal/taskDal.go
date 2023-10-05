package dal

import (
	"example.com/m/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateTodoTask(todoTask *model.TodoTask) error {
	gorm, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/togolist?charset=utf8mb4&parseTime=True&loc=Local"))
	gorm.Table("todo_tasks")
	gorm.Create(todoTask)
	return nil
}
