package controllers

import (
	"example/auth"
	"example/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var user []models.User
	err := models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var userResponse [] models.UserResponse = make([]models.UserResponse,len(user))
		for i, user := range user {
			userResponse[i] = TransferUser(user)
		}
		c.JSON(http.StatusOK, userResponse)
	}
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	user.Prepare()
	var err error
	err = user.Validate("create")
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}
	err = models.CreateUser(&user)
	if err !=nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, TransferUser(user))
	}
}

func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, TransferUser(user))
	}
}

func UpdateUser(c *gin.Context){
	var user models.User
	id := c.Params.ByName("id")
	err := models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = user.Validate("create")
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}
	err = models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, TransferUser(user))
	}
}

func DeleteUser(c *gin.Context){
	var user models.User
	id := c.Params.ByName("id")
	err :=models.DeleteUser(&user, id)
    if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

func Login(c *gin.Context) {
	var user models.User
	if err:= c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	
	if err:= models.GetUserByUsernameAndPassword(&user,user.Username, user.Password); err!= nil {
		 c.JSON(http.StatusUnprocessableEntity, err.Error())
		 return
	 }

	 token, err := auth.CreateToken(user.Id)
	 
	 if err != nil {
		 c.JSON(http.StatusUnprocessableEntity, err.Error())
	   return
		}

		c.JSON(http.StatusOK, token)
}

func TransferUser(user models.User) models.UserResponse {
	var userResponse models.UserResponse
	userResponse.Id 		= user.Id
	userResponse.Name 		= user.Name
	userResponse.Email 	= user.Email
	userResponse.Phone		= user.Phone
	userResponse.Address 	= user.Address
	userResponse.Username	= user.Username
	userResponse.Role		= user.Role
	return userResponse
}


