package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "/Users/bytedance/GolandProjects/to-go-list/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// [user[:password]@][net[(addr)]]/dbname[?param1=value1&paramN=valueN]
	gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/togolist?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // reuse your gorm db

	g.GenerateModel("todo_tasks")
	g.Execute()
}
