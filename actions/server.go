package actions

import (
	"github.com/0xN0x/go-artifactsmmo/internal/api"
	"github.com/0xN0x/go-artifactsmmo/models"
)

func (c *ArtifactsMMO) GetServerStatus() (*models.ServerStatus, error) {
	var ret models.ServerStatus

	_, err := api.NewRequest(c.Config).SetMethod("GET").SetURL("/").SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	return &ret, nil

}
