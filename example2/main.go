package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db, _ := gorm.Open("postgres",
		"user=postgres dbname=app sslmode=disable")
	defer db.Close()

	// create
	admin := &Admin{
		Name:  "gaku",
		Age:   26,
		Isman: true,
		Careers: []*Career{
			{
				Fromdate:    time.Date(2012, 4, 1, 0, 0, 0, 0, time.Local),
				Todate:      time.Date(2014, 3, 15, 0, 0, 0, 0, time.Local),
				Description: "大学生活最高！",
			},
			{
				Fromdate:    time.Date(2014, 4, 1, 0, 0, 0, 0, time.Local),
				Todate:      time.Date(2016, 8, 20, 0, 0, 0, 0, time.Local),
				Description: "しゃちく生活万歳！",
			},
			{
				Fromdate:    time.Date(2016, 9, 1, 0, 0, 0, 0, time.Local),
				Todate:      time.Date(2017, 5, 7, 0, 0, 0, 0, time.Local),
				Description: "ニート生活最高ぅううううううあsぢうふぁすdfh(｀・ω・´)！",
			},
			{
				Fromdate:    time.Date(2017, 5, 8, 0, 0, 0, 0, time.Local),
				Description: "不動産屋さん！しゃちく万歳！万歳！",
			},
		},
	}
	db.Create(admin)

	// fetch
	var myAdmin Admin
	db.First(&myAdmin, 1)
	var careers []*Career
	db.Model(&myAdmin).Related(&careers)

	// update
	myAdmin.Name = "updateName"
	db.Save(&myAdmin)
	var updateCareer Career
	db.First(&updateCareer, 1)
	updateCareer.Description = "update description!"
	db.Save(&updateCareer)

	// delete 以下で関連も削除される
	var deleteAdmin Admin
	db.Delete(&deleteAdmin, 1)
}

// テーブル名:adminsに対して、Admin or Adminsどちらの構造体名でもOK
// 管理者テーブル
type Admin struct {
	ID      int
	Name    string
	Age     int
	Isman   bool
	Careers []*Career
}

// 管理者経歴テーブル
type Career struct {
	ID          int
	Fromdate    time.Time
	Todate      time.Time
	Description string
	AdminId     int // AdminsIdではダメ。 AdminIdでテーブルカラムを作成する必要がある
}
