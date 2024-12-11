package database

import (
	"fmt"

	"github.com/Hexagonz/back-end-go/env"
	"github.com/Hexagonz/back-end-go/models"
	"github.com/Hexagonz/back-end-go/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase() (*gorm.DB, error) {
	db_conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.IfElse(
			env.DotEnvVariable("USERNAME") == "", "root", env.DotEnvVariable("USERNAME"),
		),
		utils.IfElse(
			env.DotEnvVariable("PASSWORD") == "", "", env.DotEnvVariable("PASSWORD"),
		),
		utils.IfElse(
			env.DotEnvVariable("HOST") == "", "127.0.0.1", env.DotEnvVariable("HOST"),
		),
		utils.IfElse(
			env.DotEnvVariable("PORT") == "", "3306", env.DotEnvVariable("PORT"),
		),
		utils.IfElse(
			env.DotEnvVariable("DB_NAME") == "", "mysql", env.DotEnvVariable("DB_NAME"),
		),
	)
	db, err := gorm.Open(mysql.Open(db_conn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		return nil, err
	}

	modelsToMigrate := []interface{}{
		&models.Users{},
		&models.RefreshToken{},
	}

	for _, model := range modelsToMigrate {
		if err := db.AutoMigrate(model); err != nil {
			return nil, err
		}
	}

	println("Database migration completed successfully!")
	return db, nil
}
