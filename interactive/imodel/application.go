package imodel

// Application 申请加入房间
type Application struct {
	// 申请加入的房间ID
	ApplyRoomID string
	// 申请验证信息
	ApplyContent string
}

func CreateApplication(roomID string, applyContent string) Application {
	return Application{
		ApplyRoomID:  roomID,
		ApplyContent: applyContent,
	}
}
