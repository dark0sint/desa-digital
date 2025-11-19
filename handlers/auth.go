package handlers

import (
    "desa-digital/database"
    "desa-digital/models"
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

func LoginPage(c *gin.Context) {
    c.HTML(200, "login.html", nil)
}

func RegisterPage(c *gin.Context) {
    c.HTML(200, "register.html", nil)
}

func Login(c *gin.Context) {
    username := c.PostForm("username")
    password := c.PostForm("password")

    var user models.User
    if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
        c.HTML(400, "login.html", gin.H{"Error": "Username atau password salah"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        c.HTML(400, "login.html", gin.H{"Error": "Username atau password salah"})
        return
    }

    session := sessions.Default(c)
    session.Set("user_id", user.ID)
    session.Save()

    c.Redirect(302, "/dashboard")
}

func Register(c *gin.Context) {
    username := c.PostForm("username")
    password := c.PostForm("password")

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

    user := models.User{Username: username, Password: string(hashedPassword), Role: "user"}
    if err := database.DB.Create(&user).Error; err != nil {
        c.HTML(400, "register.html", gin.H{"Error": "Username sudah ada"})
        return
    }

    c.Redirect(302, "/login")
}

func Logout(c *gin.Context) {
    session := sessions.Default(c)
    session.Clear()
    session.Save()
    c.Redirect(302, "/")
}
