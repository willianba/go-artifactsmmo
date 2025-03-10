package actions

import (
	"fmt"
	"strconv"

	"github.com/0xN0x/go-artifactsmmo/internal/api"
	"github.com/0xN0x/go-artifactsmmo/models"
)

func (c *ArtifactsMMO) GetMaps(content_code string, content_type models.MapContentType, page int, size int) (*[]models.MapSchema, error) {
	var ret []models.MapSchema
	req := api.NewRequest(c.Config).SetMethod("GET").SetURL("/maps").SetResultStruct(&ret)

	if content_code != "" {
		req.SetParam("content_code", content_code)
	}

	if content_type != "" {
		req.SetParam("content_type", string(content_type))
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

func (c *ArtifactsMMO) GetMap(x int, y int) (*models.MapSchema, error) {
	var ret models.MapSchema

	res, err := api.NewRequest(c.Config).SetMethod("GET").SetURL(fmt.Sprintf("/maps/%d/%d", x, y)).SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 404 {
		return nil, models.ErrMapNotFound
	}

	return &ret, nil
}
