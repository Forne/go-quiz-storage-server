package main

import (
	"gorm.io/gorm"
	"time"
)

type Question struct {
	ID         uint           `json:"id" gorm:"primaryKey,autoIncrement"`
	CategoryID int            `json:"category_id" form:"category_id"`
	Category   Category       `json:"-"`
	TypeID     int            `json:"type_id" form:"type_id" gorm:"index"`
	Text       string         `json:"text" form:"text"`
	MediaType  int            `json:"media_type" form:"media_type"`
	MediaUrl   string         `json:"media_url" form:"media_url"`
	LocaleCode string         `json:"locale_code" form:"locale_code"`
	Locale     Locale         `json:"-" form:"locale" gorm:"references:Code" gorm:"foreignKey:LocaleCode"`
	Answers    []Answer       `json:"answers" gorm:"foreignKey:QuestionID"`
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type Category struct {
	ID   int    `json:"id" gorm:"primaryKey,autoIncrement"`
	Text string `json:"text" form:"text"`
}

type Answer struct {
	ID         int      `json:"id" gorm:"primaryKey,autoIncrement"`
	QuestionID int      `json:"question_id" form:"question_id"`
	Question   Question `json:"-"`
	Text       string   `json:"text" form:"text"`
	Valid      bool     `json:"is_valid" form:"is_valid"`
}

type Locale struct {
	Code     string `json:"code" gorm:"primaryKey,autoIncrement"`
	Language string `json:"language"`
	Country  string `json:"country"`
}
