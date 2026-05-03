package structs

type CreateSubTokenRequest struct {
	RoomID    int64  `json:"roomId"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	ExpiresAt string `json:"expiresAt"`
}
