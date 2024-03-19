package model

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(255);comment:用户名称;NOT NULL" json:"name"`
	Mobile   string `gorm:"column:mobile;type:char(11);comment:用户电话;NOT NULL" json:"mobile"`
	Password string `gorm:"column:password;type:varchar(512);comment:用户密码;NOT NULL" json:"password"`
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

func (m *UserModel) FindUserByMobile(ctx context.Context, mobile string) (*User, error) {
	var user User
	result := m.db.WithContext(ctx).Where("mobile = ?", mobile).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (m *UserModel) CreateUser(ctx context.Context, user *User) error {
	result := m.db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
