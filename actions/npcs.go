package actions

import (
	"fmt"
	"strconv"

	"github.com/0xN0x/go-artifactsmmo/internal/api"
	"github.com/0xN0x/go-artifactsmmo/models"
)

func (c *ArtifactsMMO) GetAllNPCs(npc_type models.NPCType, page int, size int) (*[]models.NPC, error) {
	var ret []models.NPC

	req := api.NewRequest(c.Config).SetMethod("GET").SetURL("/npcs").SetResultStruct(&ret)

	if npc_type != "" {
		req.SetParam("type", string(npc_type))
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

func (c *ArtifactsMMO) GetNPC(code string) (*models.NPC, error) {
	var ret models.NPC

	res, err := api.NewRequest(c.Config).SetMethod("GET").SetURL(fmt.Sprintf("/npcs/%s", code)).SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 404 {
		return nil, models.ErrNPCNotFound
	}

	return &ret, nil
}

func (c *ArtifactsMMO) GetNPCItems(code string, page int, size int) (*[]models.NPCItem, error) {
	var ret []models.NPCItem

	req := api.NewRequest(c.Config).SetMethod("GET").SetURL(fmt.Sprintf("/npcs/%s/items", code)).SetResultStruct(&ret)

	if page != 0 {
		req.SetParam("page", strconv.Itoa(page))
	}

	if size != 0 {
		req.SetParam("size", strconv.Itoa(size))
	}

	res, err := req.Run()
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 404 {
		return nil, models.ErrNPCNotFound
	}

	return &ret, nil
}

func (c *ArtifactsMMO) NPCBuyItem(code string, quantity int) (*models.NPCTransactionResponse, error) {
	var ret models.NPCTransactionResponse

	body := models.SimpleItem{Code: code, Quantity: quantity}
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/npc/buy", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 404:
		return nil, models.ErrItemNotFound
	case 441:
		return nil, models.ErrItemCannotBePurchased
	case 492:
		return nil, models.ErrInsufficientGold
	case 598:
		return nil, models.ErrNPCNotFoundOnThisMap
	}

	return &ret, nil
}

func (c *ArtifactsMMO) NPCSellItem(code string, quantity int) (*models.NPCTransactionResponse, error) {
	var ret models.NPCTransactionResponse

	body := models.SimpleItem{Code: code, Quantity: quantity}
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/npc/sell", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 404:
		return nil, models.ErrItemNotFound
	case 442:
		return nil, models.ErrItemCannotBeSold
	case 492:
		return nil, models.ErrInsufficientGold
	case 598:
		return nil, models.ErrNPCNotFoundOnThisMap
	}

	return &ret, nil
}
