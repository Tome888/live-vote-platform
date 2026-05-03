package structs

type CreateRoomRequest struct {
	Name          string `json:"name"`
	VoteType      string `json:"voteType"`
	RoomExpiresAt string `json:"roomExpiresAt"`
}
