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

	// insert
	admins := []*Admin{
		{Name: "gaku", Age: 26, Isman: true},
		{Name: "gakuko", Age: 15, Isman: false},
		{Name: "gakuko", Age: 15, Isman: false},
		{Name: "gakuko", Age: 15, Isman: false},
		{Name: "gakuko", Age: 15, Isman: false},
		{Name: "gakuko", Age: 15, Isman: false},
		{Name: "gakuko", Age: 15, Isman: false},
	}
	for _, v := range admins {
		db.Create(&v)
	}

	// fetch
	var list []*Admin
	db.Where("name LIKE ?", "%ko%").Order("id DESC").Offset(0).Limit(3).Find(&list)
	for _, v := range list {
		fmt.Printf("%#v\n", v)
	}
	fmt.Println("----------------------")

	// update
	var admin Admin
	db.First(&admin, 1) // where id = 3と同じ

	admin.Name = "gakkaaaaaaaa"
	db.Save(&admin) // 更新対象が存在しない場合、SAVEは新規作成を実施する

	// where delete
	db.Where("name Like ?", "%ko%").Delete(&Admin{})
	db.Delete(&admin, 5) // where id = 2
	//db.Delete(&Admin{})    // all delete

	// 最後に全てをfetchし確認
	db.Order("id DESC").Find(&list)
	for _, v := range list {
		fmt.Printf("%#v\n", v)
	}
}

// テーブル名:adminsに対して、Admin or Adminsどちらの構造体名でもOK
type Admin struct {
	ID    int
	Name  string
	Age   int
	Isman bool
}
