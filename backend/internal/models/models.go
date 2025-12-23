package models

import (
	"gorm.io/gorm"
	"time"
)

// LocationType define se é Museu, Ruína, etc.
type LocationType struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name"` // Ex: "Museum", "Ruin"
	Icon      *string   `json:"icon"`                        // Ex: "landmark", "history"
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Location é o ponto principal no mapa
type Location struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name        string  `gorm:"not null" json:"name"`
	Description string  `json:"description"`
	Latitude    float64 `gorm:"not null" json:"lat"`
	Longitude   float64 `gorm:"not null" json:"lng"`
	QRCodeToken *string `gorm:"uniqueIndex" json:"-"`

	LocationTypeID uint         `json:"location_type_id"`
	LocationType   LocationType `gorm:"foreignKey:LocationTypeID" json:"location_type"`

	Artifacts []Artifact `gorm:"foreignKey:LocationID" json:"artifacts,omitempty"`
}

type Artifact struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	LocationID  uint      `json:"location_id"` // Foreign Key para Location
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	Era         string    `json:"era"`
	ImageURL    string    `json:"image_url"`
	CreatedAt   time.Time `json:"-"`
}
