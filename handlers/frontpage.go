package handlers

import "github.com/gin-gonic/gin"

func FrontPage(c *gin.Context) {
    c.HTML(200, "frontpage.html", gin.H{
        "Title": "Desa Digital - Halaman Depan",
    })
}
