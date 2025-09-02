package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `gorm:"uniqueIndex;size:32;not null" json:"username"`
	Password string    `gorm:"not null" json:"-"` // 不在json中返回
	Email    string    `gorm:"uniqueIndex;size:128;not null" json:"email"`
	Posts    []Post    `json:"-"`
	Comments []Comment `json:"-"`
}

type Post struct {
	gorm.Model
	Title    string    `gorm:"size:255;not null" json:"title"`
	Content  string    `gorm:"type:text;not null" json:"content"`
	UserId   uint      `json:"user_id"`
	User     User      `json:"author"`
	Comments []Comment `json:"-"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null" json:"content"`
	UserId  uint   `json:"user_id"`
	User    User   `json:"author"`
	PostId  uint   `json:"post_id"`
	Post    Post   `json:"-"`
}
