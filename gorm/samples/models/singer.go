package models

import (
	"context"
	"database/sql"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Singer struct {
	ID           int `gorm:"primarykey;autoIncrement:false"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime `gorm:"index"`
	Name         string
	Email        string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
	Age          int    `gorm:"<-:update"`          // allow read and update
	Birthday     sql.NullTime
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	//  custom data type not supported
	ActivatedDay ActivationDay `gorm:"-"` // ignore this field when write and read with struct
	Song         Song
}

// // `Song` belongs to `Singer`, `SingerID` is the foreign key
type Song struct {
	ID          int `gorm:"primarykey;autoIncrement:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
	Title       string
	Description string // default value `gorm:"default:placeholder"` not supported
	SingerID    int
}

func (s *Singer) BeforeCreate(tx *gorm.DB) (err error) {
	if !s.ActivatedAt.Valid {
		s.ActivatedAt = sql.NullTime{time.Now(), true}
	}
	// s.ActivatedAt = ActivationDay{ActivatedAt: s.ActivatedAt}
	return
}

// Create from customized data type
type ActivationDay struct {
	D           int
	ActivatedAt sql.NullTime
}

// Scan implements the sql.Scanner interface
func (d *ActivationDay) Scan(v interface{}) error {
	// Scan a value into struct from database driver
	d.D = v.(int)
	return nil
}

func (d *ActivationDay) GormDataType() string {
	return "INT64"
}

func (d *ActivationDay) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "EXTRACT(DAY FROM DATE ?)",
		Vars: []interface{}{d.ActivatedAt},
	}
}
