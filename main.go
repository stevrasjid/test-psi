package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Number1 struct {
	VoucherDiscount int
	ProductPrice float64	
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},	
        AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"*"},
        ExposeHeaders:    []string{"*"}, 
        AllowCredentials: true,
        MaxAge: 300,
    }))
	r.POST("/api/number1", getPointNumber1)
	r.POST("/api/number2", getToken)
	r.GET("/api/number4", getUsers)
	r.GET("/api/number5", getArray)
	
	r.Run(":8080")
}

func getPointNumber1(c *gin.Context){
	var number1 Number1
	if err := c.ShouldBindJSON(&number1); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	point := (number1.ProductPrice * (float64(number1.VoucherDiscount) / 100)) * 0.02

	c.JSON(http.StatusOK, gin.H{
		"point" : point,
	})
}

func getToken(c *gin.Context) {
	claims:= jwt.MapClaims{
		"id" : uuid.New(),
		"Username" : "username123",
		"Session_exp" : time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error() ,
		})		
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : "success",
		"token" :  tokenString,
	})
}

func getUsers(c *gin.Context) {
	page,_ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	result, _ := strconv.Atoi(c.Query("result"))
	if result == 0 {
		result = 10
	}


	url := fmt.Sprintf("https://randomuser.me/api?results=%d&page=%d",result, page)
	response, err := http.Get(url)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
    }
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	var responseBody APIResponse
	json.Unmarshal(responseData, &responseBody)

	var results []Result
	for _, data := range responseBody.Results {
		var result Result
		result.Name = fmt.Sprintf("%s, %s %s", data.Name.Title, data.Name.First, data.Name.Last)
		result.Age = data.Dob.Age
		result.Email = data.Email
		result.Cell = data.Cell
		result.Phone = data.Phone
		result.Location = fmt.Sprintf("%s, %s, %s, %s, %s", strconv.Itoa(data.Location.Street.Number), data.Location.Street.Name, data.Location.City, data.Location.State, data.Location.Country)

		var pictures []string
		pictures = append(pictures, data.Picture.Thumbnail)
		pictures = append(pictures, data.Picture.Medium)
		pictures = append(pictures, data.Picture.Large)
		result.Pictures = pictures
		
		results = append(results, result)		
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data" :  results,
		"countData" : len(results),
	})
}

func getArray(c *gin.Context) {
	colors := []string{"merah", "kuning", "hijau", "pink", "ungu"}
	clothes := []string{"baju", "celana", "topi", "jaket", "sepatu"}
	ads := []string{"Diskon", "Sale", "Diskon", "Sale", "Sale"}

	color := c.Query("color")
	if color != "" {
		colors = append(colors, color)
	}

	var result []string
	j:=0

	for index, color := range colors {
		if (j > len(clothes)) {
			j = 0
		}

		if (index == 0 || index + 1 == len(colors)){
			result = append(result, fmt.Sprintf("%s %s %s", clothes[j], color, ads[j]))
			continue;
		}
		
		if (index % 2 == 1) {
			result = append(result, fmt.Sprintf("%s %s %s", clothes[index + 1], color, ads[j]))
		} else if (index % 2 == 0){
			result = append(result, fmt.Sprintf("%s %s %s", clothes[index - 1], color, ads[j]))
		}
		j++
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "success",
		"data" : result,
	})
}