package actions

import (
	"fmt"
	"strconv"

	"github.com/0xN0x/go-artifactsmmo/internal/api"
	"github.com/0xN0x/go-artifactsmmo/models"
)

func (c *ArtifactsMMO) GetMonsters(drop string, max_level int, min_level int, page int, size int) (*[]models.Monster, error) {
	var ret []models.Monster
	req := api.NewRequest(c.Config).SetMethod("GET").SetURL("/monsters").SetResultStruct(&ret)

	if drop != "" {
		req.SetParam("drop", drop)
	}

	if min_level != 0 {
		req.SetParam("min_level", strconv.Itoa(min_level))
	}

	if max_level != 0 {
		req.SetParam("max_level", strconv.Itoa(max_level))
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

func (c *ArtifactsMMO) GetMonster(code string) (*models.Monster, error) {
	var ret models.Monster

	res, err := api.NewRequest(c.Config).SetMethod("GET").SetURL(fmt.Sprintf("/monsters/%s", code)).SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 404 {
		return nil, models.ErrMonsterNotFound
	}

	return &ret, nil
}
