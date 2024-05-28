package main

import (
	"backend/database"
	"backend/router"
	"os"

	"github.com/gin-gonic/gin"
)

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}
// 		c.Next()
// 	}
// }

func main() {
	// connecting to database
	database.Connect()
	// // // closing database when application close
	defer database.Close()
	// // creating gin engine
	ngin := gin.Default()
	// cors := CORSMiddleware()
	// ngin.Use(cors)
	// // all router api
	rtr := router.NewRouter(ngin)
	rtr.ClientRouter()
	// // // auto migrate data in development mode
	if mode := os.Getenv("MODE"); mode == "development" {
		rtr.AdminRouter()
	}
	// running server with port 8000
	// ngin.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
	// })
	// prt := os.Getenv("PORT")
	ngin.Run(":8000")
}
