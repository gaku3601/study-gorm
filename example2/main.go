package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db, _ := gorm.Open("postgres",
		"user=postgres dbname=app sslmode=disable")
	defer db.Close()

	// create
	admin := Admin{
		Name:  "gaku",
		Age:   26,
		Isman: true,
		Careers: []Career{
			{
				Fromdate:    time.Date(2012, 4, 1, 23, 59, 59, 0, time.Local),
				Todate:      time.Date(2014, 3, 15, 23, 59, 59, 0, time.Local),
				Description: "大学生活最高！",
			},
			{
				Fromdate:    time.Date(2014, 4, 1, 23, 59, 59, 0, time.Local),
				Todate:      time.Date(2016, 8, 20, 23, 59, 59, 0, time.Local),
				Description: "しゃちく生活万歳！",
			},
			{
				Fromdate:    time.Date(2016, 9, 1, 23, 59, 59, 0, time.Local),
				Todate:      time.Date(2017, 5, 7, 23, 59, 59, 0, time.Local),
				Description: "ニート生活最高ぅううううううあsぢうふぁすdfh(｀・ω・´)！",
			},
			{
				Fromdate:    time.Date(2017, 5, 8, 23, 59, 59, 0, time.Local),
				Description: "不動産屋さん！しゃちく万歳！万歳！",
			},
		},
	}
	fmt.Printf("%#v\n", admin)
	db.Create(&admin)
	db.Save(&admin)
}

// テーブル名:adminsに対して、Admin or Adminsどちらの構造体名でもOK
// 管理者テーブル
type Admin struct {
	ID      int
	Name    string
	Age     int
	Isman   bool
	Careers []Career
}

// 管理者経歴テーブル
type Career struct {
	ID          int
	Fromdate    time.Time
	Todate      time.Time
	Description string
	AdminId     int // AdminsIdではダメ。 AdminIdでテーブルカラムを作成する必要がある
}
