package main

import (
	"fmt"
	"log"
	"template/config"
	"template/repository"
	"template/service"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	config.LoadConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
			NoLowerCase:   true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.InitializeRepository(db)
	svc := service.NewCardService(repos.CardRepo)

	fmt.Println("Starting debug scrap...")
	err = svc.ScrapCards()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Success!")
	}
}
