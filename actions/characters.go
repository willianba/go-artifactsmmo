package actions

import (
	"fmt"

	"github.com/0xN0x/go-artifactsmmo/internal/api"
	"github.com/0xN0x/go-artifactsmmo/models"
)

func (c *ArtifactsMMO) CreateCharacter(name string, skin models.CharacterSkin) (*models.Character, error) {
	var character models.Character

	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL("/characters/create").SetResultStruct(&character).SetBody(&models.Character{Name: name, Skin: skin}).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 494:
		return nil, models.ErrNameAlreadyUsed
	case 495:
		return nil, models.ErrMaxCharactersReached
	}

	return &character, nil
}

func (c *ArtifactsMMO) DeleteCharacter(name string) (*models.Character, error) {
	var character models.Character

	_, err := api.NewRequest(c.Config).SetMethod("POST").SetURL("/characters/delete").SetResultStruct(&character).SetBody(&models.Character{Name: name}).Run()
	if err != nil {
		return nil, err
	}

	return &character, nil
}

func (c *ArtifactsMMO) GetCharacter(name string) (*models.Character, error) {
	var character models.Character

	_, err := api.NewRequest(c.Config).SetMethod("GET").SetURL(fmt.Sprintf("/characters/%s", name)).SetResultStruct(&character).Run()
	if err != nil {
		return nil, err
	}

	return &character, nil
}
