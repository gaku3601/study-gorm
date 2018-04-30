package main

import (
	"fmt"

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
				Description: "大学生活最高！",
			},
			{
				Description: "しゃちく生活万歳！",
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
	Description string
	AdminId     int // AdminsIdではダメ。 AdminIdでテーブルカラムを作成する必要がある
}
