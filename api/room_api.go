package api

import (
	"github.com/gin-gonic/gin"
	roomService "hotel_cms/service/room"
	"strconv"
)

// GetRooms 获取所有房间信息
func GetRooms(c *gin.Context) {

	c.JSON(200, roomService.Infos())
}

// GetRoom 获取所有房间信息
func GetRoom(c * gin.Context)  {
	roomId := c.Param("id")
	if roomId == "" {
		c.JSON(404, nil)
	} else {
		c.JSON(200, roomService.Info(strconv.Atoi(roomId)))
	}
}
