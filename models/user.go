package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name   string
	DeptDm string `json:"dept_dm"`
	Dept   Dept   `gorm:"foreign_key:DeptDm;association_key:Dm"`
}

func GetUsers(dm string) (users []User) {
	db.Where("dm = ?", dm).Find(&users)
	return
}
func AddUser(data map[string]interface{}) bool {
	if !ExistDeptByDm(data["dept_dm"].(string)) {
		AddDept(map[string]interface{}{"dm": data["dept_dm"].(string)})
	}
	db.Create(&User{
		Name:   data["name"].(string),
		DeptDm: data["dept_dm"].(string),
	})
	return true
}
func RemoveUser(id int) {
	db.Where("id = ?", id).Delete(&User{})
}
