package models

import "github.com/jinzhu/gorm"

type Dept struct {
	gorm.Model
	Dm   string `gorm:"index"`
	Name string
}

func ExistDeptByDm(dm string) bool {
	var dept Dept
	db.Select("id").Where("dm = ?", dm).First(&dept)
	if dept.ID > 0 {
		return true
	}
	return false
}
func GetDeptByDm(dm string) (dept Dept) {
	db.Where("dm = ?", dm).First(&dept)
	return
}
func AddDept(data map[string]interface{}) bool {
	db.Create(&Dept{
		Dm:   data["dm"].(string),
		Name: data["name"].(string),
	})
	return true
}
func EditDept(id int, data map[string]interface{}) bool {
	db.Model(&Dept{}).Where("id = ?", id).Updates(data)
	return true
}
func RemoveDept(id int) bool {
	db.Where("id = ?", id).Delete(&Dept{})
	return true
}
