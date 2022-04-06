package model

import (
	"encoding/json"
	"time"

	"github.com/wmcff/zoogeek/repository"

	"golang.org/x/crypto/bcrypt"
)

const selectAccount = "select * " +
	" from user "

type User struct {
	ID        int64     `gorm:"column:id" json:"id"`
	RealName  string    `gorm:"column:real_name" json:"real_name"` //  真实名称
	Email     string    `gorm:"column:email" json:"email"`         //  邮箱
	Phone     string    `gorm:"column:phone" json:"phone"`         //  电话号码
	NickName  string    `gorm:"column:nick_name" json:"nick_name"` //  昵称
	Sex       string    `gorm:"column:sex" json:"sex"`             //  性别 0女 1男 2其他
	Age       float64   `gorm:"column:age" json:"age"`             //  年龄
	Country   string    `gorm:"column:country" json:"country"`     //  国家
	City      string    `gorm:"column:city" json:"city"`           //  城市
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	ForestId  int64     `gorm:"column:forest_id" json:"forest_id"`
	Password  string    `gorm:"column:password" json:"password"`
}

type RecordUser struct {
	ID       int64   `gorm:"column:id" json:"id"`
	RealName string  `gorm:"column:real_name" json:"real_name"` //  真实名称
	Email    string  `gorm:"column:email" json:"email"`         //  邮箱
	Phone    string  `gorm:"column:phone" json:"phone"`         //  电话号码
	NickName string  `gorm:"column:nick_name" json:"nick_name"` //  昵称
	Sex      string  `gorm:"column:sex" json:"sex"`             //  性别 0女 1男 2其他
	Age      float64 `gorm:"column:age" json:"age"`             //  年龄
	Country  string  `gorm:"column:country" json:"country"`     //  国家
	City     string  `gorm:"column:city" json:"city"`           //  城市
	ForestId int64   `gorm:"column:forest_id" json:"forest_id"`
	Password string  `gorm:"column:password" json:"password"`
}

func (User) TableName() string {
	return "user"
}

func NewUser(name string, password string, authorityID uint) *User {
	return &User{RealName: name, Password: password}
}

func NewUserWithPlainPassword(name string, password string, authorityID uint) *User {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return &User{RealName: name, Password: string(hashed)}
}

func (a *User) FindByName(rep repository.Repository, name string) (*User, error) {
	var user *User

	var rec RecordUser
	rep.Raw(selectAccount+" where real_name = ?", name).Scan(&rec)
	user = convertToUser(&rec)

	return user, nil
}

func convertToUser(rec *RecordUser) *User {
	return &User{ID: rec.ID, RealName: rec.RealName, Password: rec.Password}
}

// Create persists this account data.
func (a *User) Create(rep repository.Repository) (*User, error) {
	if error := rep.Select("nick_name", "password").Create(a).Error; error != nil {
		return nil, error
	}
	return a, nil
}

// ToString is return string of object
func (a *User) ToString() (string, error) {
	bytes, error := json.Marshal(a)
	return string(bytes), error
}
