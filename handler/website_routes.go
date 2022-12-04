package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"main.go/auth"
	"main.go/model"
)

type WebId struct {
	Id uint `json:"id"`
}
type WebsiteResponse struct {
	ID         uint             `json:"id"`
	Created_at time.Time        `json:"created_at"`
	Updated_at time.Time        `json:"updated_at"`
	User_id    uint             `json:"user_id"`
	Site_key   string           `json:"site_key"`
	Secret_key string           `json:"secret_key"`
	Label      string           `json:"label"`
	Alert      bool             `json:"alert"`
	Subdomain  bool             `json:"subdomain"`
	Version    uint             `json:"version"`
	Website_v1 model.Website_v1 `json:"web_v1"`
	Token      string           `json:"token"`
}

func NewWebResponse(web *model.Website) *WebsiteResponse {
	token, _ := auth.GenerateJWT(web.ID, web.Label)
	wr := &WebsiteResponse{
		ID:         web.ID,
		Created_at: web.Created_at,
		Updated_at: web.Updated_at,
		User_id:    web.User_id,
		Site_key:   web.Site_key,
		Secret_key: web.Secret_key,
		Label:      web.Label,
		Alert:      web.Alert,
		Subdomain:  web.Subdomain,
		Version:    web.Version,
		Token:      token,
	}
	return wr
}

func (u *Users) getWebsite(c echo.Context) error {
	var input WebId
	if err := c.Bind(&input); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "Invalid inputs", err)
	}
	fmt.Println(input)
	web, err := u.Store.GetWebsite(input.Id)
	if err != nil {
		return echo.ErrBadRequest
	}
	fmt.Println(web)

	return c.JSON(http.StatusOK, NewWebResponse(web))

}
func (u *Users) updateWebsite(c echo.Context) error {
	var input WebsiteResponse
	if err := c.Bind(&input); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "Invalid inputs", err)
	}
	newWeb := &model.Website{
		ID:         input.ID,
		User_id:    input.User_id,
		Site_key:   input.Site_key,
		Secret_key: input.Secret_key,
		Label:      input.Label,
		Alert:      input.Alert,
		Subdomain:  input.Subdomain,
		Version:    input.Version,
	}
	err := u.Store.UpdateWebsite(newWeb)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not update website in the database", err)
	}
	return c.JSON(http.StatusOK, NewWebResponse(newWeb))
}
