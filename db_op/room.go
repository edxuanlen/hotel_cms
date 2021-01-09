package db_op

import (
	"hotel_cms/common"
	"hotel_cms/entity"
	"time"
)

type Room entity.Room

// GetRoom 用ID 获取房间信息
func GetRoom(Id interface{}) (Room, error) {
	var room Room
	result := DB.First(&room, Id)
	return room, result.Error
}

// UpdateRoomPrice 更改该房间类型的房间单价
func UpdateRoomPrice(Type byte, NewPrice int) (bool, error) {
	result := DB.Table(common.Room).
		Where("type = ?", Type).
		Update("price", NewPrice).
		Update("modify_time", time.Now())
	return result.RowsAffected > 0, result.Error
}

// SelectRooms 查看房间列表
func SelectRooms () ([]Room, error) {
	var rooms []Room
	result := DB.Find(&rooms)
	return rooms, result.Error
}

// SelectRoom 查看某个房间信息
func SelectRoomById(roomId int) (Room, error) {
	var room Room
	result := DB.Where("where id = ?", roomId).First(&room)
	return room, result.Error
}