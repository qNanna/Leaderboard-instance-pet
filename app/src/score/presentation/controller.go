package presentation

import (
	"encoding/json"
	"fmt"
	"net/http"

	application "main/app/src/score/application"
	dto "main/app/src/score/presentation/dto"
)

type ScoreController struct{}

var service = &application.Service{}

func (c ScoreController) SaveScore(res http.ResponseWriter, req *http.Request) {
	data := &dto.SaveScoreDto{}
	uid := req.Header.Get("uid")

	json.NewDecoder(req.Body).Decode(data)

	isAdded := service.SaveScore(uid, data.Score)
	jsonResponse := fmt.Sprintf(`{ "Done": %v }`, isAdded)

	res.Write([]byte(jsonResponse))
}

func (c ScoreController) GetScore(res http.ResponseWriter, req *http.Request) {
	scores := service.GetScore()

	response, _ := json.Marshal(scores)
	res.Write(response)
}
