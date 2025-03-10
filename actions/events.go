package actions

import (
	"strconv"

	"github.com/0xN0x/go-artifactsmmo/internal/api"
	"github.com/0xN0x/go-artifactsmmo/models"
)

func (c *ArtifactsMMO) GetAllEvents(page int, size int) (*[]models.Event, error) {
	var ret []models.Event

	req := api.NewRequest(c.Config).SetMethod("GET").SetURL("/events").SetResultStruct(&ret)

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

func (c *ArtifactsMMO) GetAllActiveEvents(page int, size int) (*[]models.ActiveEvent, error) {
	var ret []models.ActiveEvent

	req := api.NewRequest(c.Config).SetMethod("GET").SetURL("/events/active").SetResultStruct(&ret)

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
