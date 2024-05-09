package controllers

import (
	"net/http"

	// "os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/webbalaka/MassageShop_backend/initializers"
	"github.com/webbalaka/MassageShop_backend/models"
	"golang.org/x/crypto/bcrypt"
)

func hashPassWord(password string) (string, error) {
	resultPassWord, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(resultPassWord), err
}

func matchPassword(systemPassword string, incomePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(systemPassword), []byte(incomePassword))
}

func sendTokenResponse(u *models.User, s int, c *gin.Context) {
	jwtExpire, err := strconv.Atoi("30")
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	claims := jwt.MapClaims{
		"ID":  u.ID,
		"exp": time.Now().Add(10 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	jwtSecret := []byte("SecretYouShouldHide")
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Duration(jwtExpire) * time.Minute),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(c.Writer, cookie)
	c.JSON(s, gin.H{
		"success": true,
		"token":   tokenString,
	})
}

func Register(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid input data",
		})
	}

	hash, err := hashPassWord(input.Password)

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}

	newUser := models.User{
		Name:     input.Name,
		Tel:      input.Tel,
		Email:    input.Email,
		Role:     input.Role,
		Password: hash,
	}

	result := initializers.DB.Create(&newUser)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": result.Error.Error(),
		})
		return
	}

	sendTokenResponse(&newUser, 200, c)
}

func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	login := models.User{
		Email:    input.Email,
		Password: input.Password,
	}

	var user models.User

	result := initializers.DB.Where("email = ?", login.Email).First(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid credentials",
		})
		return
	}

	if err := matchPassword(user.Password, login.Password); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid credentials",
		})
		return
	}

	sendTokenResponse(&user, 200, c)
}

func Logout(c *gin.Context){
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(200, gin.H{
		"success": true,
		"message": "Logged out",
	})
}

func GetMe(c *gin.Context) {
	userClaims, exists  := c.Get("user")
	if !exists {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Not authorize to access this route",
		})
		return
	}

	var user models.User

	claims, ok := userClaims.(jwt.MapClaims)

	if !ok {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Not authorize to access this route",
		})
		return
	}

	id := claims["ID"].(float64)


	result := initializers.DB.Find(&user, id)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": user,
	})
}
