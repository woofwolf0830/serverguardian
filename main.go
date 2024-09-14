package main

import (
	// "fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/api/guaiguaidate", func(ctx *gin.Context) {
		dateFile, err := os.OpenFile("date", os.O_RDWR, 0644)
		if err != nil {
			dateFile, _ := os.Create("date")
			defer dateFile.Close()
			dateFile.Write([]byte(time.Now().Format("2006-01-02")))
			ctx.JSON(http.StatusOK, gin.H{
				"result": "OK",
				"time": time.Now().Format("2006-01-02"),
			})
		} else {
			date := make([]byte, 100)
			n, _ := dateFile.Read(date)
			guaiguaiDate, _ := time.Parse("2006-01-02", string(date[:n]))
			if guaiguaiDate.Before(time.Now()) {
				ctx.JSON(http.StatusOK, gin.H{
					"result": "EXPIRE",
					"time": string(date[:n]),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"result": "OK",
					"time": string(date[:n]),
				})
			}
		}
		defer dateFile.Close()
	})

	r.GET("/api/newguaiguai", func(ctx *gin.Context) {
		dateFile, err := os.Create("date")
		if err != nil {
			panic(err)
		}
		defer dateFile.Close()
		oneMonthLater := time.Now().AddDate(0, 1, 0)
		dateFile.Write([]byte(oneMonthLater.Format("2006-01-02")))
		ctx.JSON(http.StatusOK, "OK")
	})

	err := r.Run()
	if err != nil {
		panic(err)
	}
}