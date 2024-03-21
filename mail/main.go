package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	dataSourceFormat := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"

	dataSource := fmt.Sprintf(
		dataSourceFormat,
		"acloset-kr-dev.ckogrfmvncuu.ap-northeast-2.rds.amazonaws.com",
		5432,
		"test",
		"test",
		"acloset-dev",
	)
	db, err := gorm.Open(postgres.Open(dataSource), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	_ = db.AutoMigrate(&data{})

	// 시작 시간
	startTime := time.Now()

	var iter int = 1000
	ch := make(chan int, iter)
	for i := 0; i < iter; i++ {
		index := i
		go func() {
			//다음 시퀀스 값을 얻기 위해 raw SQL 쿼리 사용
			var result int
			tx := db.Begin()

			if err := db.Raw("SELECT nextval('id_seq')").Scan(&result).Error; err != nil {
				panic("Failed to get nextval from sequence" + err.Error())
			}
			if randInt := rand.Intn(100); randInt > 90 {
				if err := tx.Exec("crash").Error; err != nil {
					tx.Rollback()
				}
			}
			tx.Commit()

			if err := db.Create(&data{
				Name: strconv.Itoa(index),
			}).Error; err != nil {
				panic(err)
			}

			//fmt.Println("Next value:", result)

		}()
	}
	// 경과 시간
	elapsedTime := time.Since(startTime)

	fmt.Printf("실행시간: %s\n", elapsedTime)

	var mmap = make(map[int]bool, 0)

	for {
		select {
		case s := <-ch:
			if ok := mmap[s]; !ok {
				mmap[s] = true
			} else {
				panic("duplicate")
			}
		case <-time.After(time.Second * 1):
			return
		}
	}

}

type data struct {
	gorm.Model
	Name string
}
