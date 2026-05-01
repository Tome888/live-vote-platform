package structs

type CreateRoomResponse struct {
	RoomID        int64  `json:"room_id"`
	Name          string `json:"name"`
	Iat           int64  `json:"iat"`
	ConnectionKey string `json:"connection_key"`
}
