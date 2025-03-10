package actions

import (
	"github.com/0xN0x/go-artifactsmmo/internal/api"
	"github.com/0xN0x/go-artifactsmmo/models"
)

func (c *ArtifactsMMO) GetAccountDetails() (*models.AccountDetails, error) {
	var ret models.AccountDetails

	_, err := api.NewRequest(c.Config).SetMethod("GET").SetURL("/my/details").SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *ArtifactsMMO) ChangePassword(currentPassword string, newPassword string) (*string, error) {
	var ret string

	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL("/my/change-password").SetParam("current_password", currentPassword).SetParam("new_password", newPassword).SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 458:
		return nil, models.ErrUseDifferentPassword
	}

	return &ret, nil
}
