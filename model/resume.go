package model

import (
	"encoding/json"

	"github.com/wmcff/serve-data/repository"
)

// Resume defines struct of resume data.
type Resume struct {
	ID        int64  `gorm:"column:id" json:"id"`
	Name      string `gorm:"column:name" json:"name"`           //  森林名称
	Area      string `gorm:"column:area" json:"area"`           //  森林面积
	Longitude string `gorm:"column:longitude" json:"longitude"` //  经度
	Latitude  string `gorm:"column:latitude" json:"latitude"`   //  纬度
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt string `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName returns the table name of resume struct and it is used by gorm.
func (Resume) TableName() string {
	return "resume"
}

// NewResume is constructor
func NewResume(name string) *Resume {
	return &Resume{Name: name}
}

// Exist returns true if a given resume exits.
func (c *Resume) Exist(rep repository.Repository, id uint) (bool, error) {
	var count int64
	if error := rep.Where("id = ?", id).Count(&count).Error; error != nil {
		return false, error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// FindByID returns a resume full matched given resume's ID.
func (c *Resume) FindByID(rep repository.Repository, id uint) (*Resume, error) {
	var resume Resume
	if error := rep.Where("id = ?", id).First(&resume).Error; error != nil {
		return nil, error
	}
	return &resume, nil
}

// FindAll returns all resumes of the resume table.
func (c *Resume) FindAll(rep repository.Repository) (*[]Resume, error) {
	var resumes []Resume
	if error := rep.Find(&resumes).Error; error != nil {
		return nil, error
	}
	return &resumes, nil
}

// Create persists this resume data.
func (c *Resume) Create(rep repository.Repository) (*Resume, error) {
	if error := rep.Create(c).Error; error != nil {
		return nil, error
	}
	return c, nil
}

// ToString is return string of object
func (c *Resume) ToString() (string, error) {
	bytes, error := json.Marshal(c)
	return string(bytes), error
}
