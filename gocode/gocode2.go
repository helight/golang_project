package main

import (
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pikanezi/mapslice"
	"github.com/tealeg/xlsx"
)

type Role struct {
	Id     int    `gorm:"primary_key"` //主键id
	Name   string `gorm:"not null"`    //角色名
	Remark string
	Status int
}

type RoleBack struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
	Status string `json:"status"`
}

//规定表名
func (Role) TableName() string {
	return "role"
}

var DB *gorm.DB

//初始化并保持连接
func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:123456@/test?charset=utf8&parseTime=True&loc=Local")
	//    DB.LogMode(true)//打印sql语句
	if err != nil {
		log.Fatalf("database connect is err:%s", err.Error())
	} else {
		log.Print("connect database is success")
	}
	err = DB.DB().Ping()
	if err != nil {
		DB.DB().Close()
		log.Fatalf("Error on opening database connection: %s", err.Error())
	}
}

func main() {
	var role RoleBack
	var rolelists []RoleBack
	rolelist, _ := GetRoleList()
	for i := 0; i < len(rolelist); i++ {
		role.Id = strconv.Itoa(rolelist[i].Id)
		role.Name = rolelist[i].Name
		role.Remark = rolelist[i].Remark
		role.Status = strconv.Itoa(rolelist[i].Status)
		rolelists = append(rolelists, role)
	}
	id, _ := mapslice.ToStrings(rolelists, "Id")
	name, _ := mapslice.ToStrings(rolelists, "Name")
	remark, _ := mapslice.ToStrings(rolelists, "Remark")
	status, _ := mapslice.ToStrings(rolelists, "Status")

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("Sheet1")
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "编号"
	cell = row.AddCell()
	cell.Value = "名称"
	cell = row.AddCell()
	cell.Value = "状态"
	cell = row.AddCell()
	cell.Value = "备注"
	for i := 0; i < len(id); i++ {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = id[i]
		cell = row.AddCell()
		cell.Value = name[i]
		cell = row.AddCell()
		cell.Value = status[i]
		cell = row.AddCell()
		cell.Value = remark[i]
	}
	file.Save("File.xlsx")
}

func GetRoleList() (rolelist []Role, err error) {
	err = DB.Find(&rolelist).Error
	return rolelist, err
}
