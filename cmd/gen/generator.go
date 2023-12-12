package main

import (
	"shorturl/dal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const (
	dsn = "root:123456@tcp(127.0.0.1:3306)/shorturl?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
)

func connectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// 生成shorturl表的query
	g.UseDB(connectDB(dsn))
	g.ApplyBasic(g.GenerateAllTable()...)
	g.ApplyInterface(func(model.ImpulseSender) {}, g.GenerateModel("sequence"))
	g.Execute()
}
