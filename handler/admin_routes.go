package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"main.go/auth"
	"main.go/model"
)

func (u *Users) SignUp(c echo.Context) error {
	var req userRequest
	if err := c.Bind(&req); err != nil {
		log.Printf("can't build request to admin user :%v", err)
		return echo.ErrBadRequest
	}
	pass, _ := model.HashPassword(req.Password)
	NewAdmin := &model.Admin{
		Password: pass,
		Email:    req.Email,
	}
	if err := u.Store.NewAdmin(c.Request().Context(), NewAdmin); err != nil {
		log.Printf("can't signup Admin with id : %v", req.ID)
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusCreated, NewAdmin)

}

func (u *Users) Signin(c echo.Context) error {
	authRequest := &userAuthRequest{}
	if err := c.Bind(authRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "error binding user request", err)
	}

	admin, err := u.Store.GetAdminByEmail(authRequest.Email)
	if err != nil {
		return echo.ErrInternalServerError
	}
	passValidate := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(authRequest.Password))

	if !(passValidate == nil) {
		return echo.ErrUnauthorized
	}
	token, _ := auth.GenerateJWT(admin.ID, admin.Email)
	user := UserResponse{
		Id:    admin.ID,
		Email: admin.Email,
		Token: token,
	}
	return c.JSON(http.StatusOK, user)

}
