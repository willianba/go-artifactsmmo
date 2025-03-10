package actions

import (
	"fmt"

	"github.com/0xN0x/go-artifactsmmo/internal/api"
	"github.com/0xN0x/go-artifactsmmo/models"
)

func (c *ArtifactsMMO) Move(x int, y int) (*models.CharacterMovementData, error) {
	var move models.CharacterMovementData

	body := models.Movement{X: x, Y: y}
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/move", c.Config.GetUsername())).SetResultStruct(&move).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 404:
		return nil, models.ErrMapNotFound
	case 490:
		return nil, models.ErrAlreadyAtDestination
	}

	return &move, nil
}

func (c *ArtifactsMMO) Rest() (*models.Rest, error) {
	var ret models.Rest

	_, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/rest", c.Config.GetUsername())).SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *ArtifactsMMO) Equip(code string, slot models.Slot, quantity int) (*models.EquipRequest, error) {
	var equip models.EquipRequest

	body := models.ItemInventory{Code: code, Slot: slot, Quantity: quantity}
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/equip", c.Config.GetUsername())).SetResultStruct(&equip).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 404:
		return nil, models.ErrItemNotFound
	case 484:
		return nil, models.ErrTooMuchConsumables
	case 485:
		return nil, models.ErrItemAlreadyEquiped
	case 491:
		return nil, models.ErrSlotNotEmpty
	case 496:
		return nil, models.ErrLevelTooLow
	}

	return &equip, nil
}

func (c *ArtifactsMMO) Unequip(slot models.Slot, quantity int) (*models.EquipRequest, error) {
	var unequip models.EquipRequest

	body := models.RemoveItemInventory{Slot: slot, Quantity: quantity}
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/unequip", c.Config.GetUsername())).SetResultStruct(&unequip).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 404:
		return nil, models.ErrItemNotFound
	}

	return &unequip, nil
}

func (c *ArtifactsMMO) UseConsumable(code string, quantity int) (*models.Consumable, error) {
	var ret models.Consumable

	body := models.SimpleItem{Code: code, Quantity: quantity}
	_, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/use", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *ArtifactsMMO) Fight() (*models.CharacterFight, error) {
	var fight models.CharacterFight

	_, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/fight", c.Config.GetUsername())).SetResultStruct(&fight).Run()
	if err != nil {
		return nil, err
	}

	return &fight, nil
}

func (c *ArtifactsMMO) Gather() (*models.SkillData, error) {
	var skill models.SkillData

	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/gathering", c.Config.GetUsername())).SetResultStruct(&skill).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 493:
		return nil, models.ErrInsufficientSkillLevel
	case 598:
		return nil, models.ErrRessourceNotFound
	}

	return &skill, nil
}

func (c *ArtifactsMMO) Craft(code string, quantity int) (*models.SkillData, error) {
	var ret models.SkillData

	body := models.SimpleItem{Code: code, Quantity: quantity}
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/crafting", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 404:
		return nil, models.ErrCraftNotFound
	case 493:
		return nil, models.ErrInsufficientSkillLevel
	case 598:
		return nil, models.ErrWorkshopNotFound
	}

	return &ret, nil
}

func (c *ArtifactsMMO) Recycling(code string, quantity int) (*models.Recycling, error) {
	var ret models.Recycling

	body := models.SimpleItem{Code: code, Quantity: quantity}
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/recycling", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 404:
		return nil, models.ErrItemNotFound
	case 473:
		return nil, models.ErrCantBeRecycled
	case 493:
		return nil, models.ErrInsufficientSkillLevel
	case 598:
		return nil, models.ErrWorkshopNotFound
	}

	return &ret, nil
}

func (c *ArtifactsMMO) DeleteItem(code string, quantity int) (*models.ItemReponse, error) {
	var ret models.ItemReponse

	body := models.SimpleItem{Code: code, Quantity: quantity}
	_, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/delete", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

// TODO add get characters logs

func (c *ArtifactsMMO) GetMyCharactersInfo() (*[]models.Character, error) {
	var ret []models.Character

	_, err := api.NewRequest(c.Config).SetMethod("GET").SetURL("/my/characters").SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
