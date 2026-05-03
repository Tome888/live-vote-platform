package structs

import "github.com/golang-jwt/jwt/v5"

type RoomToken struct {
	RoomId        int64  `json:"room_id"`
	Name          string `json:"name"`
	Role          string `json:"role"`
	VoteType      string `json:"vote_type"`
	RoomExpiresAt string `json:"room_expires_at"`
	jwt.RegisteredClaims
}
