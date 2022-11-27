package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
	ID                     uint   `json:"id"`
	Email                  string `json:"email"`
	Name                   string `json:"name"`
	Password               string `json:"password"`
	Phone                  string `json:"phone"`
	Company_name           string `json:"company_name"`
	Job_title              string `json:"job_title"`
	Active                 bool   `json:"active"`
	Subscribe_news         bool   `json:"subscribe_news"`
	Subscribe_notification bool   `json:"subscribe_notification"`
}

func NewUserResponse(user *model.User) *UserResponse {
	token, _ := auth.GenerateJWT(user.ID, user.Email)
	ur := &UserResponse{Id: user.ID, Email: user.Email, Token: token}
	return ur
}
func (u *Users) signUp(c echo.Context) error {
	var req userRequest
	if err := c.Bind(&req); err != nil {
		log.Printf("can't build request to user :%v", err)
		return echo.ErrBadRequest
	}
	pass, _ := model.HashPassword(req.Password)
	NewAdmin := &model.Admin{
		ID:       req.ID,
		Password: pass,
		Email:    req.Email,
	}
	if err := u.Store.NewAdmin(c.Request().Context(), NewAdmin); err != nil {
		log.Printf("can't signup Admin with id : %v", req.ID)
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusCreated, NewAdmin)

}

func (u *Users) signin(c echo.Context) error {
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
func (u *Users) CreateUser(c echo.Context) error {
	var req userRequest
	if err := c.Bind(&req); err != nil {
		log.Printf("can't build request to user :%v", err)
		return echo.ErrBadRequest
	}
	pass, _ := model.HashPassword(req.Password)
	newUser := &model.User{
		ID:                     req.ID,
		Created_at:             time.Now(),
		Updated_at:             time.Now(),
		Name:                   req.Name,
		Password:               pass,
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
func (u *Users) GetAll(c echo.Context) error {
	users, err := u.Store.GetAll()
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusAccepted, users)
}
func (u *Users) Get(c echo.Context) error {
	return nil

}
func (u *Users) UpdateUser(c echo.Context) error {
	var input userRequest
	if err := c.Bind(&input); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "Invalid inputs", err)
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
	user.Subscribe_notification = input.Subscribe_notification
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

func (a Users) Register(e *echo.Echo) {
	e.POST("/admin/signup", a.signUp)
	e.POST("/admin/signin", a.signin)
	e.GET("/all", a.GetAll)
	e.GET("/:id", a.Get)
	u := e.Group("users")
	u.Use(auth.JWT())
	u.POST("/create", a.CreateUser)
	u.POST("/get", a.GetAll)
	u.POST("/update", a.UpdateUser)
	u.POST("/delete", a.DeleteUser)
	// w :=e.Group("website")
	// e.Use(auth.jwt())
	// w.POST("/get",a.GetWebsite)
	//w.Post("/update".a.UpdateWebsite)

}
func extractID(c echo.Context) uint {
	e := c.Get("user").(*jwt.Token)
	claims := e.Claims.(jwt.MapClaims)
	id := uint(claims["id"].(float64))
	return id
}
