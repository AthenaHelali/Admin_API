package store

import (
	"context"
	"fmt"
	"log"

	"github.com/labstack/echo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"main.go/model"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(database *gorm.DB) *UserPostgres {
	return &UserPostgres{db: database}
}
func (store *UserPostgres) Save(ctx context.Context, m *model.User) error {
	fmt.Println("akbar")
	fmt.Println(m)
	if err := store.db.Create(&m).Error; err != nil {
		log.Printf("user creation on Postgres failed: %v", err)
		return echo.ErrInternalServerError
	}
	return nil
}
func (store *UserPostgres) Get(ctx context.Context, id uint) (*model.User, error) {
	user := new(model.User)
	if err := store.db.First(&user, id); err != nil {
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
func (store *UserPostgres) DeleteUser(email string) (*model.User, error) {
	user := new(model.User)
	if err := store.db.Clauses(clause.Returning{}).Where("email = ?", email).Delete(&user).Error; err != nil {
		log.Printf("couldn't delete user with email : %v. %v", email, err)
		return nil, echo.ErrInternalServerError
	}
	return user, nil
}
func (store *UserPostgres) GetAll() ([]model.User, error) {
	var users []model.User
	if err := store.db.Find(&users).Error; err != nil {
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
	if err := store.db.Table("admins").Create(m).Error; err != nil {
		log.Printf("admin sign up on Postgres failed: %v", err)
		return echo.ErrInternalServerError
	}
	return nil

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
func (store *UserPostgres) DeletePlan(id uint) (model.User_plan, error) {
	plan := new(model.User_plan)
	if err := store.db.Clauses(clause.Returning{}).Where("user_id = ?", id).Delete(&plan).Error; err != nil {
		log.Printf("can't delete the user plan  with id : %v", id)
		return *plan, err
	}
	return *plan, nil
}

//website functions

func (store *UserPostgres) GetWebsite(web_id uint) (*model.Website, error) {
	web := new(model.Website)
	fmt.Println(web_id)
	if err := store.db.Preload("Website_v1").Find(&web, web_id).Error; err != nil {
		log.Printf("can't get the website : %v", web_id)
		return nil, err
	}
	fmt.Println(web)
	return web, nil
}
func (store *UserPostgres) UpdateWebsite(web *model.Website) error {
	website := new(model.Website)
	if err := store.db.Model(&website).Where("id = ?", web.ID).Omit("id", "created_at", "user_id", "site_key", "secret_key").Updates(web).Error; err != nil {
		log.Printf("can't save the website with id : %v", web.ID)
		return err
	}
	return nil
}
