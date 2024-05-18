package application

import (
	database "main/app/database/repository"
	entity "main/app/src/score/domain"
	infrastructure "main/app/src/score/infrastructure"
	steam "main/app/steam/repository"
)

var scoreRepository infrastructure.IScoreRepository = &database.ScoreRepository{}
var steamRepository infrastructure.ISteamRepository = &steam.SteamRepository{}

type Service struct{}

func (s *Service) SaveScore(uid string, score int) bool {
	return scoreRepository.SaveScore(uid, score)
}

func (s *Service) GetScore() []entity.Score {
	scores := scoreRepository.GetScore()

	var userIds []string
	for _, user := range scores {
		userIds = append(userIds, user.UserID)
	}

	players := steamRepository.GetSteamPlayers(userIds)

	var result []entity.Score
	for _, score := range scores {
		for _, player := range players {
			if score.UserID == player.SteamID {
				result = append(result, entity.Score{UserID: score.UserID, Name: player.PersonaName, Score: score.Score})
			}
		}
	}

	return result
}
