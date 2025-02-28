package model

type Platform struct {
	Id     string
	UserId string `firestore:"user_id"`
	Type   string `firestore:"type"`
}

func NewPlatform(id, userId, platformType string) *Platform {
	return &Platform{
		Id:     id,
		UserId: userId,
		Type:   platformType,
	}
}
