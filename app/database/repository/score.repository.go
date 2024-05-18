package repository

import (
	"main/app/database"
	"main/app/database/schema"
	"main/app/database/types"

	"github.com/thoas/go-funk"
)

type ScoreRepository struct{}

func (s *ScoreRepository) SaveScore(id string, score int) bool {
	data := schema.Score{
		Score:  score,
		UserID: id,
	}

	output := database.GetConnection().Create(&data)
	return output.RowsAffected == 1
}

func (s *ScoreRepository) GetScore() []types.GetScoreResponse {
	var scores []types.GetScoreResponse

	queryString := `
		SELECT DISTINCT s.user_id, u.name, MAX(score) AS score 
		FROM Scores s
		INNER JOIN Users u ON s.user_id = u.id
		GROUP BY s.user_id 
		ORDER BY s.score DESC
	`
	database.GetConnection().Raw(queryString).Scan(&scores)

	return funk.Map(scores, func(score types.GetScoreResponse) types.GetScoreResponse {
		return types.GetScoreResponse{UserID: score.UserID, Name: score.Name, Score: score.Score}
	}).([]types.GetScoreResponse)
}
