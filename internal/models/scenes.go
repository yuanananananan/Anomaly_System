package models

import "gorm.io/datatypes"

type Domain string

const (
	DomainLand  Domain = "land"
	DomainWater Domain = "water"
)

type Scene struct {
	BaseModel
	SceneKey string         `gorm:"type:varchar(64);uniqueIndex;not null" json:"scene_key"`
	Name     string         `gorm:"type:varchar(128);not null;default:''" json:"name"`
	Domain   Domain         `gorm:"type:enum('land','water');not null" json:"domain"`
	Location string         `gorm:"type:varchar(256);not null;default:''" json:"location"`
	MetaJSON datatypes.JSON `gorm:"type:json" json:"meta_json"`
}

func (Scene) TableName() string { return "scenes" }
