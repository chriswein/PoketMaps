package main

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Map_part struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Container struct {
	Data []Map_part `json:"data"`
}

type MetaData struct {
	NumParts int `json:"numparts"`
	N int `json:"n"`
	M int `json:"m"`
	DetailLevels []int `json:"detaillevels"`
}


var maps = []Map_part{}

func getMetaData(c *gin.Context){
	md := MetaData{NumParts:81,N:9,M:9, DetailLevels: []int{25,50,100}}
	c.IndentedJSON(http.StatusOK, md)
}

func getMaps(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, maps)
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PATCH, DELETE, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()
	// router.Use(CorsMiddleware()) // use for local development
	content, err := os.ReadFile("./images/100/info.json")
	if err != nil {
		log.Fatal("Error while opening file", err)
	}
	var da Container
	err = json.Unmarshal(content, &da)
	if err != nil {
		log.Fatal("Error while parsing file", err)
	}
	for _, element := range da.Data {
		maps = append(maps, element)
		fmt.Print(element)
	}

	router.GET("/maps", getMaps)
	router.GET("/meta", getMetaData)
	router.Static("/images/25/", "./images/25/")
	router.Static("/images/50/", "./images/50/")
	router.Static("/images/75/", "./images/75/")
	router.Static("/images/100/", "./images/100/")
	router.StaticFile("/", "./public/index.html") // Serve index for root
	router.NoRoute(func(c *gin.Context) {
	    // Check if the requested file exists in the public folder
	    c.File("./public" + c.Request.URL.Path)
	})
	router.Run("127.0.0.1:6767")
}                       
                        
