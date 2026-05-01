package structs

type Room struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	ConnectionKey string `json:"connection_key"`
	CreatedAt     string `json:"created_at"`
}
