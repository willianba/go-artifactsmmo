package actions

import (
	"fmt"
	"strconv"

	"github.com/0xN0x/go-artifactsmmo/internal/api"
	"github.com/0xN0x/go-artifactsmmo/models"
)

func (c *ArtifactsMMO) GetGESellOrders(code string, page int, size int) (*[]models.GEOrder, error) {
	var ret []models.GEOrder

	req := api.NewRequest(c.Config).SetMethod("GET").SetURL("/my/grandexchange/orders").SetResultStruct(&ret)

	if code != "" {
		req.SetParam("code", string(code))
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

func (c *ArtifactsMMO) GetGESellHistory(code string, id string, page int, size int) (*[]models.GEOrderHistory, error) {
	var ret []models.GEOrderHistory

	req := api.NewRequest(c.Config).SetMethod("GET").SetURL("/my/grandexchange/history").SetResultStruct(&ret)

	if code != "" {
		req.SetParam("code", string(code))
	}

	if id != "" {
		req.SetParam("id", string(id))
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

func (c *ArtifactsMMO) GEBuyItem(code string, quantity int, price int) (*models.GETransaction, error) {
	var ret models.GETransaction

	body := models.GEItem{Code: code, Quantity: quantity}
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/grandexchange/buy", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 404:
		return nil, models.ErrOrderNotFound
	case 434:
		return nil, models.ErrOfferDoesNotContainItems
	case 435:
		return nil, models.ErrYouCantBuyToYourself
	case 436:
		return nil, models.ErrTransactionOther
	case 486:
		return nil, models.ErrTransactionCharacter
	case 492:
		return nil, models.ErrInsufficientGold
	case 598:
		return nil, models.ErrGENotFound
	}

	return &ret, nil
}

func (c *ArtifactsMMO) GECreateSellOrder(code string, quantity int, price int) (*models.GEOrderCreated, error) {
	var ret models.GEOrderCreated

	body := models.GEItem{Code: code, Quantity: quantity, Price: price}
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/grandexchange/sell", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 404:
		return nil, models.ErrItemNotFound
	case 433:
		return nil, models.ErrCantCreateOrdersAtTheSameTime
	case 437:
		return nil, models.ErrItemCannotBeSold
	case 486:
		return nil, models.ErrTransactionCharacter
	case 492:
		return nil, models.ErrInsufficientGold
	case 598:
		return nil, models.ErrGENotFound
	}

	return &ret, nil
}

func (c *ArtifactsMMO) GECancelSellOrder(id string) (*models.GETransaction, error) {
	var ret models.GETransaction

	body := struct {
		Id string `json:"id"`
	}{
		Id: id,
	}
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/grandexchange/cancel", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 404:
		return nil, models.ErrOrderNotFound
	case 436:
		return nil, models.ErrTransactionOther
	case 438:
		return nil, models.ErrCantCancelOrder
	case 486:
		return nil, models.ErrTransactionCharacter
	case 598:
		return nil, models.ErrGENotFound
	}

	return &ret, nil
}
