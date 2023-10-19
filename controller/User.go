package controller

import (
	"net/http"
	"os"
	"set-up-Golang/config"
	"set-up-Golang/helper"
	"set-up-Golang/model"
	request "set-up-Golang/model/Request"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user model.User
	request := request.Register{}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"msg":   "this field is required, Please fill",
			"reson": "error" + "(" + err.Error() + ")",
			"data":  nil,
		})
		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	db := config.ConnectDB()
	user = model.User{
		Nama:     request.Nama,
		Email:    request.Email,
		Password: string(hash),
	}

	err = db.Create(&user).Error

	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"msg":    "Succes Insert Data",
		"reason": nil,
		"data": gin.H{
			"nama":  user.Nama,
			"email": user.Email,
		},
	})
}

func Login(c *gin.Context) {
	login := request.Login{}

	if err := c.BindJSON(&login); err != nil {
		helper.Badrequest(c, http.StatusBadRequest, err)
		return
	}

	db := config.ConnectDB()

	// validasi Email
	var user model.User
	if err := db.Where("email = ?", login.Email).Find(&user).Error; err != nil {
		helper.ErrorCustom(c, http.StatusInternalServerError, "email user tidak ditemukan", err)
		return
	}

	// validasi Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		helper.ErrorCustom(c, http.StatusInternalServerError, "password salah, harap ulangi", err)
		return
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id_user":  user.Id,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
		"namaUser": user.Nama,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		panic(err)
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Autharization", tokenString, 3600*24*30, "", "", false, true)
	helper.SuccesResponse(c, http.StatusOK, "Success Login", user)
}

func Logout(c *gin.Context) {
	c.SetCookie("Autharization", "", -1, "", "", false, true)

	helper.SuccesResponse(c, http.StatusOK, "Logout Success", nil)
}
