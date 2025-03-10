package actions

import (
	"fmt"
	"strconv"

	"github.com/0xN0x/go-artifactsmmo/internal/api"
	"github.com/0xN0x/go-artifactsmmo/models"
)

func (c *ArtifactsMMO) GetBankDetails() (*models.Bank, error) {
	var ret models.Bank

	_, err := api.NewRequest(c.Config).SetMethod("GET").SetURL("/my/bank").SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *ArtifactsMMO) GetBankItems(code string, page int, size int) (*[]models.SimpleItem, error) {
	var ret []models.SimpleItem

	req := api.NewRequest(c.Config).SetMethod("GET").SetURL("/my/bank/items").SetResultStruct(&ret)

	if code != "" {
		req.SetParam("item_code", string(code))
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

func (c *ArtifactsMMO) DepositBank(code string, quantity int) (*models.BankItemTransaction, error) {
	var ret models.BankItemTransaction

	body := models.SimpleItem{Code: code, Quantity: quantity}

	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/bank/deposit", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 404:
		return nil, models.ErrItemNotFound
	case 461:
		return nil, models.ErrTransactionInProgress
	case 462:
		return nil, models.ErrBankFull
	case 598:
		return nil, models.ErrBankNotFound
	}

	return &ret, nil
}

func (c *ArtifactsMMO) DepositBankGold(quantity int) (*models.BankGoldTransaction, error) {
	var ret models.BankGoldTransaction

	body := models.Gold{Quantity: quantity}
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/bank/deposit/gold", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 461:
		return nil, models.ErrTransactionInProgress
	case 492:
		return nil, models.ErrInsufficientGold
	case 598:
		return nil, models.ErrBankNotFound
	}

	return &ret, nil
}

func (c *ArtifactsMMO) WithdrawBank(code string, quantity int) (*models.BankItemTransaction, error) {
	var ret models.BankItemTransaction

	body := models.SimpleItem{Code: code, Quantity: quantity}
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/bank/withdraw", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 404:
		return nil, models.ErrItemNotFound
	case 461:
		return nil, models.ErrTransactionInProgress
	case 598:
		return nil, models.ErrBankNotFound
	}

	return &ret, nil
}

func (c *ArtifactsMMO) WithdrawBankGold(quantity int) (*models.BankGoldTransaction, error) {
	var ret models.BankGoldTransaction

	body := models.Gold{Quantity: quantity}
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/bank/withdraw/gold", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 460:
		return nil, models.ErrInsufficientGold
	case 461:
		return nil, models.ErrTransactionInProgress
	case 598:
		return nil, models.ErrBankNotFound
	}

	return &ret, nil
}

func (c *ArtifactsMMO) BuyBankExpansion() (*models.BankTransaction, error) {
	var ret models.BankTransaction

	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/bank/buy_expansion", c.Config.GetUsername())).SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 492:
		return nil, models.ErrInsufficientGold
	case 598:
		return nil, models.ErrBankNotFound
	}

	return &ret, nil
}
