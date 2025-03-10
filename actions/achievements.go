package actions

import (
	"fmt"
	"strconv"

	"github.com/0xN0x/go-artifactsmmo/internal/api"
	"github.com/0xN0x/go-artifactsmmo/models"
)

func (c *ArtifactsMMO) GetAchievements(ach_type models.AchievementType, page int, size int) (*[]models.Achievement, error) {
	var ret []models.Achievement
	req := api.NewRequest(c.Config).SetMethod("GET").SetURL("/achievements").SetResultStruct(&ret)

	if ach_type != "" {
		req.SetParam("ach_type", string(ach_type))
	}

	if page != 0 {
		req.SetParam("page", strconv.Itoa(page))
	}

	if size != 0 {
		req.SetParam("size", strconv.Itoa(size))
	}

	_, err := req.Run()
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *ArtifactsMMO) GetAchievement(code string) (*models.Achievement, error) {
	var ret models.Achievement

	res, err := api.NewRequest(c.Config).SetMethod("GET").SetURL(fmt.Sprintf("/achievements/%s", code)).SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 404 {
		return nil, models.ErrAchievementNotFound
	}

	return &ret, nil
}
