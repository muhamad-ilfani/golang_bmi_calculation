package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
)

type User struct {
	Fullname string  `json:"fullname" form:"fullname"`
	DOB      string  `json:"dob" form:"dob"`
	Height   float64 `json:"height" form:"height"`
	Weight   float64 `json:"weight" form:"weight"`
	Age      int     `json:"age" form:"age"`
	Status   string  `json:"status" form:"status"`
	BMI      float64 `json:"bmi" form:"bmi"`
}

var user User = User{}

func main() {
	e := echo.New()
	e.GET("/user", GetUserController)
	e.POST("/user", CreateUserController)
	e.Start(":8000")
}
func GetUserController(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}
func CreateUserController(c echo.Context) error {
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	CalculateAge(&user)
	CalculateBMI(&user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success",
		"fullname": user.Fullname,
		"age":      user.Age,
		"bmi":      user.BMI,
		"status":   user.Status,
	})
}
func CalculateAge(user *User) {
	date := strings.Split(user.DOB, "/")
	dt_now := strings.Split(time.Now().Format("01-02-2006"), "-")
	date_conv, _ := strconv.Atoi(date[2])
	dt_now_conv, _ := strconv.Atoi(dt_now[2])
	user.Age = dt_now_conv - date_conv
}
func CalculateBMI(user *User) {
	height := user.Height / 100
	user.BMI = user.Weight / (height * height)
	fmt.Println(user.BMI)
	if user.BMI <= 18.5 {
		user.Status = "Underweight"
	} else if user.BMI > 18.5 && user.BMI <= 24.9 {
		user.Status = "Normal"
	} else if user.BMI > 24.9 && user.BMI <= 29.9 {
		user.Status = "Overweight"
	} else {
		user.Status = "Obesity"
	}
}
