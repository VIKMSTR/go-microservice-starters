package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gogin-rest-db/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func main() {
	var db *gorm.DB
	router := gin.Default()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&UserModel{})
	if err != nil {
		panic("failed to migrate")
	}
	router.Use(middleware.Database(db)) //passing database object to middleware- then we can use it in our handlers
	router.POST("/user", addUser)
	router.PUT("/user", addUser)
	router.GET("/user/:id", getUser)
	router.GET("/users", getUsers)
	router.DELETE("/user/:id", deleteUser)

	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil { //here we bind from json to type
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" { //checking the values (naive auth method)
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	//runs on 8080 by default
	err = router.Run()
	if err != nil {
		panic(err)
	}
}

func deleteUser(context *gin.Context) {
	var user UserModel
	db := context.MustGet("DB").(*gorm.DB)
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}
	db.First(&user, id)
	if user.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	db.Delete(&user, id)
	context.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("user %s %s deleted", user.FirstName, user.LastName)})
}

func getUsers(context *gin.Context) {
	var userModels []UserModel
	db := context.MustGet("DB").(*gorm.DB)
	db.Find(&userModels)
	context.JSON(http.StatusOK, toUsers(userModels))
}

func getUser(c *gin.Context) {
	var userModel UserModel
	db := c.MustGet("DB").(*gorm.DB)
	db.First(&userModel, c.Param("id"))
	user := toUser(userModel)
	c.JSON(200, &user)
}

func addUser(c *gin.Context) {
	var json User
	if err := c.ShouldBindJSON(&json); err != nil { //here we bind from json to type
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.MustGet("DB").(*gorm.DB)
	userModel := toUserModel(json)
	db.Create(&userModel)
	c.JSON(200, gin.H{"status": fmt.Sprintf("user %s %s added successfully", userModel.FirstName, userModel.LastName), "id": userModel.ID})
}

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type User struct {
	FirstName   string `json:"firstname" binding:"required"`
	LastName    string `json:"lastname"  binding:"required"`
	YearOfBirth uint   `json:"yearofbirth"  binding:"required"`
}

type UserModel struct {
	gorm.Model
	FirstName   string
	LastName    string
	YearOfBirth uint
}

func toUser(usermodel UserModel) User {
	return User{FirstName: usermodel.FirstName, LastName: usermodel.LastName, YearOfBirth: usermodel.YearOfBirth}
}

func toUserModel(user User) UserModel {
	return UserModel{FirstName: user.FirstName, LastName: user.LastName, YearOfBirth: user.YearOfBirth}
}

func toUsers(userModels []UserModel) []User {
	var users []User
	for _, userModel := range userModels {
		users = append(users, toUser(userModel))
	}
	return users
}
