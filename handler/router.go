package handler

import (
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"main.go/auth"
	"main.go/model"
)

type WebId struct {
	Id uint `json:"website_id"`
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
