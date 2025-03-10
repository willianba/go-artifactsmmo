package actions

import (
	"strconv"

	"github.com/0xN0x/go-artifactsmmo/internal/api"
	"github.com/0xN0x/go-artifactsmmo/models"
)

func (c *ArtifactsMMO) GetCharactersLeaderboard(page int, size int, sort models.CharacterLeaderboardSort) (*models.CharacterLeaderboard, error) {
	var ret models.CharacterLeaderboard

	req := api.NewRequest(c.Config).SetMethod("GET").SetURL("/leaderboard/characters").SetResultStruct(&ret)

	if page != 0 {
		req.SetParam("page", strconv.Itoa(page))
	}

	if size != 0 {
		req.SetParam("size", strconv.Itoa(size))
	}

	if sort != "" {
		req.SetParam("sort", string(sort))
	}

	_, err := req.Run()
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *ArtifactsMMO) GetAccountsLeaderboard(page int, size int, sort models.AccountLeaderboardSort) (*models.AccountLeaderboard, error) {
	var ret models.AccountLeaderboard

	req := api.NewRequest(c.Config).SetMethod("GET").SetURL("/leaderboard/accounts").SetResultStruct(&ret)

	if page != 0 {
		req.SetParam("page", strconv.Itoa(page))
	}

	if size != 0 {
		req.SetParam("size", strconv.Itoa(size))
	}

	if sort != "" {
		req.SetParam("sort", string(sort))
	}

	_, err := req.Run()
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
