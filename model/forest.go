package model

import (
	"encoding/json"

	"github.com/wmcff/zoogeek/repository"
)

// Forest defines struct of forest data.
type Forest struct {
	ID        int64  `gorm:"column:id" json:"id"`
	Name      string `gorm:"column:name" json:"name"`           //  森林名称
	Area      string `gorm:"column:area" json:"area"`           //  森林面积
	Longitude string `gorm:"column:longitude" json:"longitude"` //  经度
	Latitude  string `gorm:"column:latitude" json:"latitude"`   //  纬度
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt string `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName returns the table name of forest struct and it is used by gorm.
func (Forest) TableName() string {
	return "forest"
}

// NewForest is constructor
func NewForest(name string) *Forest {
	return &Forest{Name: name}
}

// Exist returns true if a given forest exits.
func (c *Forest) Exist(rep repository.Repository, id uint) (bool, error) {
	var count int64
	if error := rep.Where("id = ?", id).Count(&count).Error; error != nil {
		return false, error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// FindByID returns a forest full matched given forest's ID.
func (c *Forest) FindByID(rep repository.Repository, id uint) (*Forest, error) {
	var forest Forest
	if error := rep.Where("id = ?", id).First(&forest).Error; error != nil {
		return nil, error
	}
	return &forest, nil
}

// FindAll returns all forests of the forest table.
func (c *Forest) FindAll(rep repository.Repository) (*[]Forest, error) {
	var forests []Forest
	if error := rep.Find(&forests).Error; error != nil {
		return nil, error
	}
	return &forests, nil
}

// Create persists this forest data.
func (c *Forest) Create(rep repository.Repository) (*Forest, error) {
	if error := rep.Create(c).Error; error != nil {
		return nil, error
	}
	return c, nil
}

// ToString is return string of object
func (c *Forest) ToString() (string, error) {
	bytes, error := json.Marshal(c)
	return string(bytes), error
}
