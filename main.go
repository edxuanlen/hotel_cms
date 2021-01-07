package main

import (
	"hotel_cms/entity"
	"hotel_cms/util"
	"hotel_cms/vo"
	"time"
)

func main() {
	user := entity.User{
		Id:           1,
		IdentityCard: "12",
		Name:         "dmr",
		Phone:        "123",
		Password:     "111",
		Salary:       5,
		Status:       1,
		Access:       4,
		EntryTime:    time.Time{},
		LastPayday:   time.Time{},
		CreateTime:   time.Time{},
		UpdateTime:   time.Time{},
	}

	simple := vo.UserSimpleInfo{}
	util.Convert(&user, &simple)


	//// 从配置文件读取配置
	//conf.Init()
	//
	//// 装载路由
	//r := server.NewRouter()
	//r.Run(":3000")
}
