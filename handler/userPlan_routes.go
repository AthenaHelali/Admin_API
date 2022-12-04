package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"main.go/auth"
	"main.go/model"
)

type PlanReq struct {
	User_id uint      `json:"user_id" validate:"required"`
	Plan_id uint      `json:"plan_id"  validate:"required,min=1,max =2"`
	ID      uint      `json:"id"`
	Ex_time time.Time `json:"ex_time"`
}
type PlanResponse struct {
	ID         uint      `json:"id"`
	Plan_id    uint      `json:"plan_id"`
	User_id    uint      `json:"user_id"`
	Ex_time    time.Time `json:"ex_time"`
	Created_at time.Time `json:"created_at"`
	Token      string    `json:"token"`
}

func NewPlanResponse(plan *model.User_plan) *PlanResponse {
	token, _ := auth.GenerateJWT(plan.Plan_id+plan.ID+plan.User_id, "")
	pr := &PlanResponse{Plan_id: plan.Plan_id, User_id: plan.User_id, ID: plan.ID, Created_at: plan.Created_at, Ex_time: plan.Ex_time, Token: token}
	return pr
}

func (u *Users) getPlan(c echo.Context) error {
	var input PlanReq
	if err := c.Bind(&input); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "Invalid inputs", err)
	}
	plan, err := u.Store.GetPlan(input.User_id)
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, "can't get plan", err)
	}
	return c.JSON(http.StatusOK, NewPlanResponse(plan))
}
func (u *Users) updatePlan(c echo.Context) error {
	var input PlanReq
	if err := c.Bind(&input); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "Invalid inputs", err)
	}
	plan, err := u.Store.GetPlan(input.User_id)
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, "can't get plan", err)
	}
	newPlan := model.User_plan{
		ID:         input.ID,
		User_id:    input.User_id,
		Plan_id:    input.Plan_id,
		Updated_at: time.Now(),
		Ex_time:    input.Ex_time,
		Created_at: plan.Created_at,
	}
	if err := u.Store.UpdatePlan(&newPlan); err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, NewPlanResponse(&newPlan))
}
func (u *Users) deletePlan(c echo.Context) error {
	var input PlanReq
	if err := c.Bind(&input); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "Invalid inputs", err)
	}
	if err := u.Store.DeletePlan(input.User_id); err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, NewPlanResponse(new(model.User_plan)))
}
