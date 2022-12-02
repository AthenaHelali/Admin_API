package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"main.go/auth"
	"main.go/model"
	"main.go/store"
)

type Users struct {
	Store store.UserPostgres
}
type reqId struct {
	Id uint `json:"id"`
}
type userAuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserResponse struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}
type userRequest struct {
	ID                      uint   `json:"id"`
	Email                   string `json:"email" validate:"email,required"`
	Name                    string `json:"name"`
	Password                string `json:"password" validate:"required,min=8,max=16"`
	Ouath_id                string `json:"ouath_id"`
	Phone                   string `json:"phone"`
	Company_name            string `json:"company_name"`
	Job_title               string `json:"job_title"`
	Active                  bool   `json:"active"`
	Subscribe_news          bool   `json:"subscribe_news"`
	Subscribe_notifications bool   `json:"subscribe_notifications"`
}

func NewUserResponse(user *model.User) *UserResponse {
	token, _ := auth.GenerateJWT(user.ID, user.Email)
	ur := &UserResponse{Id: user.ID, Email: user.Email, Token: token}
	return ur
}

func (u *Users) CreateUser(c echo.Context) error {
	var req userRequest
	if err := c.Bind(&req); err != nil {
		log.Printf("can't build request to user :%v", err)
		return echo.ErrBadRequest
	}
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "failed to validate struct")
	}

	pass, _ := model.HashPassword(req.Password)
	newUser := &model.User{
		ID:                      req.ID,
		Email:                   req.Email,
		Created_at:              time.Now(),
		Updated_at:              time.Now(),
		Name:                    req.Name,
		Password:                pass,
		Phone:                   req.Phone,
		Oauth_id:                req.Ouath_id,
		Company_name:            req.Company_name,
		Job_title:               req.Job_title,
		Active:                  req.Active,
		Subscribe_news:          req.Subscribe_news,
		Subscribe_notifications: req.Subscribe_notifications,
	}
	if u.Store.DuplicateUser(req.ID) {
		log.Printf("this id already exists in database : %v", req.ID)
		return echo.ErrBadRequest
	}
	if err := u.Store.Save(c.Request().Context(), newUser); err != nil {
		log.Printf("can't signup user with id : %v", req.ID)
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusCreated, newUser)

}
func (u *Users) GetAll(c echo.Context) error {
	users, err := u.Store.GetAll()
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusAccepted, users)
}
func (u *Users) UpdateUser(c echo.Context) error {
	var input userRequest
	if err := c.Bind(&input); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "Invalid inputs", err)
	}
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "failed to validate struct")
	}
	user, err := u.Store.GetUserByEmail(input.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found", err)
	}
	user.Email = input.Email
	user.Password = input.Password
	user.Name = input.Name
	user.Company_name = input.Company_name
	user.Active = input.Active
	user.Job_title = input.Job_title
	user.Phone = input.Phone
	user.Subscribe_news = input.Subscribe_news
	user.Subscribe_notifications = input.Subscribe_notifications
	user.Password, _ = model.HashPassword(user.Password)
	if err := u.Store.Update(c.Request().Context(), user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not update user in the database", err)
	}

	return c.JSON(http.StatusCreated, NewUserResponse(user))

}
func (u *Users) DeleteUser(c echo.Context) error {
	var req reqId
	if err := c.Bind(&req); err != nil {
		log.Printf("can't build request to user :%v", err)
		return echo.ErrBadRequest

	}
	deletedUser, err := u.Store.DeleteUser(req.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "can't delete user")
	}
	return c.JSON(http.StatusOK, NewUserResponse(deletedUser))
}
