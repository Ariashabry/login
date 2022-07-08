package main

import (
	controllers "github.com/ariashabry/login/contoller"
	"github.com/ariashabry/login/middlewares"
	"github.com/ariashabry/login/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {

	//connect dan migrate models to db
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}
	log.Println("connect to db success")

	err = models.MigrateModel(db)

	if err != nil {
		log.Println(err)
	}
	log.Println("migrate to db success")

	//add this for production
	//gin.SetMode(gin.ReleaseMode)

	// you can use this or new ones
	//r := gin.Default()

	r := gin.New()

	r.SetTrustedProxies([]string{"127.0.0.1"})

	public := r.Group("/api")

	h := controllers.Context{Gin: r, DB: db}
	public.POST("/register", h.Register)
	public.POST("/login", h.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", h.CurrentUser)

	log.Printf("API connected at port: %v", 8080)
	r.Run(":8080")

}
