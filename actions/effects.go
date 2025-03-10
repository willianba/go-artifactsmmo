package actions

import (
	"fmt"
	"strconv"

	"github.com/0xN0x/go-artifactsmmo/internal/api"
	"github.com/0xN0x/go-artifactsmmo/models"
)

func (c *ArtifactsMMO) GetAllEffects(page int, size int) (*[]models.Effect, error) {
	var ret []models.Effect

	req := api.NewRequest(c.Config).SetMethod("GET").SetURL("/effects").SetResultStruct(&ret)

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

func (c *ArtifactsMMO) GetEffect(code string) (*models.Effect, error) {
	var ret models.Effect

	_, err := api.NewRequest(c.Config).SetMethod("GET").SetURL(fmt.Sprintf("/effects/%s", code)).SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
