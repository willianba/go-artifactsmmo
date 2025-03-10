package actions

import (
	"fmt"
	"strconv"

	"github.com/0xN0x/go-artifactsmmo/internal/api"
	"github.com/0xN0x/go-artifactsmmo/models"
)

func (c *ArtifactsMMO) GetAllBadges(page int, size int) (*[]models.Badge, error) {
	var ret []models.Badge

	req := api.NewRequest(c.Config).SetMethod("GET").SetURL("/badges").SetResultStruct(&ret)

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

func (c *ArtifactsMMO) GetBadge(code string) (*models.Badge, error) {
	var ret models.Badge

	_, err := api.NewRequest(c.Config).SetMethod("GET").SetURL(fmt.Sprintf("/badges/%s", code)).SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
