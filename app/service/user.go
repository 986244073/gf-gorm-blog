package user

import (
	"errors"
	"gf-app/app/model"
)

func CheckUser(name string) error {
	var user model.User
	db.Where("username =?", name).First(&user)
	if user.ID > 0 {
		return errors.New("已经存在用户")
	}
	return nil
}
func CreateUser(data *model.User) error {
	err := db.Create(&data).Error
	if err != nil {
		panic("查询错误")
	}
	return nil
}
