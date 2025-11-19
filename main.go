package main

import (
    "desa-digital/database"
    "desa-digital/handlers"
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/gin-gonic/gin"
)

func main() {
    database.InitDB() // Inisialisasi database

    r := gin.Default()
    r.LoadHTMLGlob("templates/*") // Load templates
    r.Static("/static", "./static") // Serve static files

    // Session setup
    store := cookie.NewStore([]byte("secret"))
    r.Use(sessions.Sessions("desa-session", store))

    // Routes
    r.GET("/", handlers.FrontPage)
    r.GET("/login", handlers.LoginPage)
    r.POST("/login", handlers.Login)
    r.GET("/register", handlers.RegisterPage)
    r.POST("/register", handlers.Register)
    r.GET("/dashboard", handlers.Dashboard)
    r.GET("/course/:id", handlers.CoursePage)
    r.POST("/logout", handlers.Logout)

    r.Run(":8080") // Jalankan di port 8080
}
