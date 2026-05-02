package structs

import "github.com/golang-jwt/jwt/v5"

type RoomToken struct {
	RoomId int64  `json:"room_id"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}
