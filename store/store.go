package store

import (
	"context"
	"fmt"
	"log"

	"github.com/labstack/echo"
	"gorm.io/gorm"
	"main.go/model"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(database *gorm.DB) *UserPostgres {
	return &UserPostgres{db: database}
}
func (store *UserPostgres) Save(ctx context.Context, m *model.User) error {
	if err := store.db.Table("users").Create(m).Error; err != nil {
		log.Printf("user creation on Postgres failed: %v", err)
		return echo.ErrInternalServerError
	}
	return nil
}
func (store *UserPostgres) Get(ctx context.Context, id uint) (*model.User, error) {
	user := new(model.User)
	if err := store.db.Table("users").First(&user, id); err != nil {
		log.Printf("user not found in Postgres: %v", err)
		return nil, echo.ErrInternalServerError
	}
	return user, nil

}
func (store *UserPostgres) GetUserByEmail(email string) (*model.User, error) {
	user := new(model.User)
	if err := store.db.Table("users").Where("email = ?", email).First(user).Error; err != nil {
		log.Printf("user not found in database: %v", err)
		return nil, echo.ErrInternalServerError
	}
	return user, nil
}
func (store *UserPostgres) GetAdminByEmail(email string) (*model.Admin, error) {
	admin := new(model.Admin)
	if err := store.db.Where("email = ?", email).First(admin).Error; err != nil {
		log.Printf("admin not found in database: %v", err)
		return nil, echo.ErrInternalServerError
	}
	return admin, nil
}
func (store *UserPostgres) DeleteUser(id uint) (*model.User, error) {
	user_plan := new(model.User_plan)
	user := new(model.User)
	if err := store.db.Table("user_plans").Where("user_id = ?", id).Delete(&user_plan).Error; err != nil {
		log.Printf("couldn't delete user plan with id : %v. %v", id, err)
		return nil, echo.ErrInternalServerError
	}
	if err := store.db.Table("users").Delete(&user, id).Error; err != nil {
		log.Printf("couldn't delete user with id : %v. %v", id, err)
		return nil, echo.ErrInternalServerError
	}
	return user, nil
}
func (store *UserPostgres) DuplicateUser(id uint) bool {
	var user model.User
	if store.db.Table("users").Find(&user, id).RowsAffected > 0 {
		return true
	}
	return false

}

func (store *UserPostgres) GetAll() ([]model.User, error) {
	var users []model.User
	if err := store.db.Find(&users).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return users, nil
}
func (store *UserPostgres) Update(ctx context.Context, m *model.User) error {
	if err := store.db.Table("users").Save(&m); err != nil {
		return err.Error
	}
	return nil
}

func (store *UserPostgres) NewAdmin(ctx context.Context, m *model.Admin) error {
	if !store.db.Migrator().HasTable(&model.Admin{}) {

		store.db.Migrator().CreateTable(&model.Admin{})
	}
	if err := store.db.Table("admins").Create(m).Error; err != nil {
		log.Printf("admin sign up on Postgres failed: %v", err)
		return echo.ErrInternalServerError
	}
	return nil

}
func (store *UserPostgres) DuplicateAdmin(id uint) bool {
	var user model.User
	if store.db.Table("admins").Find(&user, id).RowsAffected > 0 {
		return true
	}
	return false

}

//plan functions

func (store *UserPostgres) GetPlan(user_id uint) (*model.User_plan, error) {
	var plan = new(model.User_plan)
	if err := store.db.Where("user_id = ?", user_id).Find(&plan).Error; err != nil {
		log.Printf("can't find plan  with user id : %v", user_id)
		return nil, err
	}
	return plan, nil
}
func (store *UserPostgres) UpdatePlan(plan *model.User_plan) error {
	if err := store.db.Save(&plan).Error; err != nil {
		log.Printf("can't save the plan  with id : %v", plan.Plan_id)
		return err
	}
	return nil
}
func (store *UserPostgres) DeletePlan(id uint) error {
	plan := new(model.User_plan)
	if err := store.db.Where("user_id = ?", id).Delete(&plan).Error; err != nil {
		log.Printf("can't delete the user plan  with id : %v", id)
		return err
	}
	return nil
}

//website functions

func (store *UserPostgres) GetWebsite(id uint) (*model.Website, error) {
	// /w1 := new(model.Website_v1)
	web := new(model.Website)
	if err := store.db.Find(&web).Preload("Website_v1").Error; err != nil {
		fmt.Println(err)
		log.Printf("can't get the website : %v", id)
		return nil, err
	}
	return web, nil
}
func (store *UserPostgres) UpdateWebsite(web *model.Website) error {
	if err := store.db.Save(&web).Error; err != nil {
		log.Printf("can't save the website with id : %v", web.ID)
		return err
	}
	return nil
}
