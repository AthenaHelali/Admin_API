package handler

import (
	"log"
	"net/http"

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
type WebId struct {
	Id uint `json:"website_id"`
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
	Email                   string `json:"email"`
	Name                    string `json:"name"`
	Password                string `json:"password"`
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
func (u *Users) signUp(c echo.Context) error {
	var req userRequest
	if err := c.Bind(&req); err != nil {
		log.Printf("can't build request to user :%v", err)
		return echo.ErrBadRequest
	}
	if u.Store.DuplicateAdmin(req.ID) {
		log.Printf("this id already exists in database : %v", req.ID)
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

func (a Users) Register(e *echo.Echo) {
	e.POST("/admin/signup", a.signUp)
	e.POST("/admin/signin", a.signin)
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

}
func extractID(c echo.Context) uint {
	e := c.Get("user").(*jwt.Token)
	claims := e.Claims.(jwt.MapClaims)
	id := uint(claims["id"].(float64))
	return id
}
