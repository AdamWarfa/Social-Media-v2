package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"somev2/internal/models"
)

func GetGamesByDate(date string) []models.NbaGame {

	var apiResponse struct {
		Data []models.NbaGame `json:"data"`
	}

	url := fmt.Sprintf("https://www.balldontlie.io/api/v1/games?dates[]=%v", date)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	err = json.Unmarshal(body, &apiResponse)

	if err != nil {
		fmt.Println(err)
	}

	return apiResponse.Data
}
