package handlers

import (
	"BakeMeets/api/models"
	"BakeMeets/repository"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type UserHandler struct {
	Ur repository.IUserRepostory
}

func (uh *UserHandler) SignUp(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	userRes, err := uh.Ur.SignUp(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, userRes)
}


func (uh *UserHandler) Login(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	tokenString, err := uh.Ur.Login(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	// postmanで動作確認する場合cookie.Secureコメントアウト
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uh *UserHandler) Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	// postmanで動作確認する場合cookie.Secureコメントアウト
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uh *UserHandler) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
