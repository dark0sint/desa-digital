package handlers

import (
    "desa-digital/database"
    "desa-digital/models"
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
    "strconv"
)

func CoursePage(c *gin.Context) {
    session := sessions.Default(c)
    if session.Get("user_id") == nil {
        c.Redirect(302, "/login")
        return
    }

    id, _ := strconv.Atoi(c.Param("id"))
    var course models.Course
    database.DB.First(&course, id)

    c.HTML(200, "course.html", gin.H{
        "Course": course,
    })
}
