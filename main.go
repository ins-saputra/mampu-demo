package main

import (
	"mampu-demo/handler"
	"mampu-demo/model"
	"mampu-demo/repository"
	"mampu-demo/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:Admin123@tcp(127.0.0.1:3306)/wallet_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("gagal connect ke db..." + err.Error())
	}

	db.AutoMigrate(&model.Wallet{})

	repo := repository.NewWalletRepository(db)
	svc := service.NewWalletService(repo)
	hdl := handler.NewWalletHandler(svc)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/balance/:userId", hdl.GetBalance)
	e.POST("/withdraw", hdl.Withdraw)

	e.Logger.Fatal(e.Start(":8080"))
}
