package handler

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"main.go/auth"
)

func (a Users) Register(e *echo.Echo) {
	e.POST("/admin/signup", a.SignUp)
	e.POST("/admin/signin", a.Signin)
	u := e.Group("users")
	u.Use(auth.JWT())
	u.POST("/create", a.CreateUser)
	u.POST("/get", a.GetAll)
	u.POST("/update", a.UpdateUser)
	u.POST("/delete", a.DeleteUser)
	w := e.Group("website")
	w.Use(auth.JWT())
	w.POST("/get", a.getWebsite)
	w.POST("/update", a.updateWebsite)
	p := e.Group("/plans")
	p.Use(auth.JWT())
	p.POST("/get", a.getPlan)
	p.POST("/update", a.updatePlan)
	p.POST("/delete", a.deletePlan)

}
func extractID(c echo.Context) uint {
	e := c.Get("user").(*jwt.Token)
	claims := e.Claims.(jwt.MapClaims)
	id := uint(claims["id"].(float64))
	return id
}
