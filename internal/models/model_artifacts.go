package models

import "gorm.io/datatypes"

type ModelArtifact struct {
	BaseModel
	AlgoID      uint64         `gorm:"not null;index" json:"algo_id"`
	Name        string         `gorm:"type:varchar(128);not null" json:"name"`
	Version     string         `gorm:"type:varchar(64);not null;index" json:"version"`
	ArtifactURI string         `gorm:"type:varchar(512);not null" json:"artifact_uri"`
	MetaJSON    datatypes.JSON `gorm:"type:json" json:"meta_json"`
	IsActive    bool           `gorm:"not null;default:true;index" json:"is_active"`
}

func (ModelArtifact) TableName() string { return "models" }
