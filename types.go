package main

import (
	"time"
)

type Question struct {
	ID        uint      `gorm:"primaryKey,autoIncrement"`
	Type      uint      `gorm:"index" form:"type"`
	Question  string    `form:"question"`
	MediaUrl  string    `form:"media_url"`
	MediaType uint      `form:"media_type"`
	Locale    string    `gorm:"index" form:"locale"`
	Answer    string    `form:"answer"`
	Answer1   string    `form:"answer1"`
	Answer2   string    `form:"answer2"`
	Answer3   string    `form:"answer3"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
