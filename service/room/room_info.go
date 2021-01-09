package room

import (
	"hotel_cms/db_op"
	"hotel_cms/vo"
)

// Infos 房间信息
func Infos() vo.Response {
	result, err := db_op.SelectRooms()
	return DueWithResult(result, err)
}

// Info 房间信息
func Info(roomId int, err error) vo.Response {
	if err != nil {
		return vo.ParamErr(err, "参数转化错误")
	}
	result, err := db_op.SelectRoomById(roomId)
	return DueWithResult(result, err)
}

func DueWithResult(result interface{}, err error) vo.Response {
	if err != nil {
		return vo.Response{
			Data: result,
			Code: 200,
			Msg:  "success",
		}
	} else {
		return vo.Err(300, "getRoom infos error: ", err)
	}
}