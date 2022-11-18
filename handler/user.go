package handler

import "github.com/labstack/echo/v4"

type User struct {
}

func (u User) GetAll(c echo.Context) error {
	return nil
}
func (u User) Get(c echo.Context) error {
	return nil

}
func (u User) login(c echo.Context) error {
	return nil
}
func (u User) SignUp(c echo.Context) error {
	return nil
}

func (u User) Register(e *echo.Echo) {
	e.GET("/all", u.GetAll)
	e.GET("/:id", u.Get)
	e.POST("/signup", u.SignUp)
	e.POST("/login", u.login)
}
