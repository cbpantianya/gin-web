package database

import "time"

// AccessToken stores the access token for a user
type AccessToken struct {
	UserID      string    `gorm:"column:user_id;primaryKey;not null;index"`
	AccessToken string    `gorm:"column:access_token;not null;index"`
	ExpiredAt   time.Time `gorm:"column:expired_at;not null"`
	CreateAt    time.Time `gorm:"column:create_at;not null"`
}

type UserInfo struct {
	UserID       string    `gorm:"column:user_id;primaryKey;not null;index"`
	Ico          string    `gorm:"column:ico"`
	Name         string    `gorm:"column:name;not null"`
	Sex          int       `gorm:"column:sex"`
	BirthDate    time.Time `gorm:"birth_date"`
	PersonalSign string    `gorm:"personal_sign"`
}

type VCodeRequestRecord struct {
	CreateAt    time.Time
	PhoneNumber string `gorm:"column:phone_number;not null;index"`
}
