package models

import "gorm.io/datatypes"

type AlgorithmType string

const (
	AlgoDetector AlgorithmType = "detector"
	AlgoAnomaly  AlgorithmType = "anomaly"
)

type Algorithm struct {
	BaseModel
	Name              string         `gorm:"type:varchar(64);uniqueIndex;not null" json:"name"`
	Type              AlgorithmType  `gorm:"type:enum('detector','anomaly');not null" json:"type"`
	SupportedDomains  datatypes.JSON `gorm:"type:json;not null" json:"supported_domains_json"`
	DefaultParamsJSON datatypes.JSON `gorm:"type:json;not null" json:"default_params_json"`
	ParamSchemaJSON   datatypes.JSON `gorm:"type:json;not null" json:"param_schema_json"`
	IsActive          bool           `gorm:"not null;default:true;index" json:"is_active"`
}

func (Algorithm) TableName() string { return "algorithms" }
