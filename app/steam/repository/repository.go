package repository

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"main/app/steam/types"

	"github.com/gofor-little/env"
)

type SteamRepository struct{}

func (s *SteamRepository) GetSteamPlayers(ids []string) []types.SteamPlayer {
	notExpandedUrl := env.Get("STEAM_API_GET_USER_URL", "")
	url := os.ExpandEnv((notExpandedUrl + strings.Join(ids[:], ",")))

	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("%v", err)
	}

	defer response.Body.Close()

	data := &types.SteamPlayersResponse{}
	json.NewDecoder(response.Body).Decode(data)

	return data.Response.Players
}
