package handlers

import (
    "desa-digital/database"
    "desa-digital/models"
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
    session := sessions.Default(c)
    userID := session.Get("user_id")
    if userID == nil {
        c.Redirect(302, "/login")
        return
    }

    var courses []models.Course
    database.DB.Find(&courses)

    c.HTML(200, "dashboard.html", gin.H{
        "Courses": courses,
    })
}
