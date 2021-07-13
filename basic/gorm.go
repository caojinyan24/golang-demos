package basic

import (
	"fmt"
	"gorm.io/gorm"
	"time"

	"gorm.io/driver/mysql"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var db *gorm.DB

func initDb() {
	dbGorm, err := gorm.Open(mysql.Open("username:password@tcp(host:port)/database?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(fmt.Sprintf("Init db: connect db error, %v", err))
	}
	db = dbGorm
}
func TestGormConnect() {
	initDb()
	defer db.Rollback()
	tx := db.Begin()
	//err := Db.CreateTable(&Product{Code: "L1212", Price: 1000}).Error
	err := tx.Table("products").Where("code=1").Update("code", "abc1").Error
	fmt.Println("create table  err,err=", err)
	err = db.Table("products").Create(&Product{Code: "L1211", Price: 1000}).Error
	fmt.Println("create err,err=", err)
	tx.Commit()
}

type ApplyInfo struct {
	ID               int64     `json:"id"`
	ApplyUser        string    ` json:"applyUser"`
	AuditUser        string    ` json:"auditUser"`
	UserDepartment   string    ` json:"userDepartment" `
	AppId            int64     ` json:"AppId"`
	AuditStatus      int64     ` json:"auditStatus"`
	ApplyDesc        string    ` json:"applyDesc"`
	ApplyType        int       ` json:"applyType"`
	ApplyServiceName string    ` json:"applyServiceName"`
	ApplyParams      string    ` json:"applyParams"`
	AuditResult      string    ` json:"auditResult"`
	AuditTime        time.Time `json:"auditTime"`
}

//两张表做连接,并取指定的字段读取到结构体中
func TestScan() {
	initDb()
	var data = make([]ApplyInfo, 0)
	var count int64 = 0
	conn := db.Table("apply_info").Select("apply_info.id as id,apply_info.apply_user as apply_user,b.audit_status as audit_status,apply_info.audit_user as audit_user,apply_info.app_id as app_id,apply_info.apply_service_name as apply_service_name,b.created_at as created_at,apply_info.apply_desc as apply_desc").Joins("inner join audit_history b on apply_info.id=b.apply_id").Where("b.audit_user=?", "caojinyan")
	conn.Count(&count)
	conn = conn.Offset(0).Limit(10)
	err := conn.Scan(&data).Error
	fmt.Println(data, count, err)
}
