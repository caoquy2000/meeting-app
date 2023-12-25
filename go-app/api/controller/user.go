package controller

import (
	"net/http"
	"time"

	"github.com/caoquy2000/meeting-app/api/service"
	"github.com/caoquy2000/meeting-app/models"
	"github.com/caoquy2000/meeting-app/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserController struct {
	service service.UserService
}

func NewUserController(s service.UserService) UserController {
	return UserController{
		service: s,
	}
}

func (u *UserController) CreateUser(c *gin.Context) {
	var user models.UserRegister
	if err := c.ShouldBind(&user); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
		return
	}

	hashPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashPassword

	err := u.service.CreateUser(user)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Failed to create user")
		return
	}

	utils.SuccessJSON(c, http.StatusOK, "Successfully Created User")
}

func (u *UserController) Login(c *gin.Context) {
	var user models.UserLogin
	var hmacSampleSecret []byte

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
		return
	}

	dbUser, err := u.service.LoginUser(user)

	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Invalid Login Credential")
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": dbUser,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "Failed to get token")
		return
	}

	response := &utils.Response{
		Success: true,
		Message: "Token success",
		Data:    tokenString,
	}

	c.JSON(http.StatusOK, response)
}
