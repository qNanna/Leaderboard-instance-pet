package infrastructure

import (
	databaseTypes "main/app/database/types"
	types "main/app/steam/types"

	"gorm.io/gorm"
)

type IUserRepository interface {
	SaveUser(uid string) *gorm.DB
}

type IScoreRepository interface {
	SaveScore(id string, score int) bool
	GetScore() []databaseTypes.GetScoreResponse
}

type ISteamRepository interface {
	GetSteamPlayers(ids []string) []types.SteamPlayer
}
