package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"main.go/model"
	"main.go/store"
)

type App interface {
	SignUp(echo.Context) error
	login(echo.Context, model.User) error
}
type Users struct {
	Store store.UserPostgre
}

func (u *Users) GetAll(c echo.Context) error {
	return nil
}
func (u *Users) Get(c echo.Context) error {
	return nil

}
func (u *Users) signin(c echo.Context) error {
	return nil
}
func (u *Users) SignUp(c echo.Context) error {
	var req model.User
	if err := c.Bind(&req); err != nil {
		log.Printf("can't build request to user :%v", err)
		return echo.ErrBadRequest
	}
	newUser := &model.User{
		ID:                     req.ID,
		Created_at:             req.Created_at,
		Updated_at:             req.Updated_at,
		Name:                   req.Name,
		Password:               req.Password,
		Phone:                  req.Phone,
		Company_name:           req.Company_name,
		Job_title:              req.Job_title,
		Active:                 req.Active,
		Subscribe_news:         req.Subscribe_news,
		Subscribe_notification: req.Subscribe_notification,
	}
	if err := u.Store.Save(c.Request().Context(), newUser); err != nil {
		log.Printf("can't signup user with id : %v", req.ID)
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusCreated, u)

}

func (a Users) Register(e *echo.Echo) {
	e.GET("/all", a.GetAll)
	e.GET("/:id", a.Get)
	e.POST("/user/signup", a.SignUp)
	e.POST("/user/signin", a.signin)
}
