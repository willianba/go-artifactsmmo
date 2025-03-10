package actions

import (
	"fmt"
	"strconv"

	"github.com/0xN0x/go-artifactsmmo/internal/api"
	"github.com/0xN0x/go-artifactsmmo/models"
)

func (c *ArtifactsMMO) GetItems(craft_material string, craft_skill string, max_level int, min_level int, name string, item_type models.ItemType, page int, size int) (*[]models.Item, error) {
	var ret []models.Item
	req := api.NewRequest(c.Config).SetMethod("GET").SetURL("/items").SetResultStruct(&ret)

	if craft_material != "" {
		req.SetParam("craft_material", craft_material)
	}

	if craft_skill != "" {
		req.SetParam("craft_skill", craft_skill)
	}

	if name != "" {
		req.SetParam("name", name)
	}

	if item_type != "" {
		req.SetParam("type", string(item_type))
	}

	if min_level != 0 {
		req.SetParam("min_level", strconv.Itoa(min_level))
	}

	if max_level != 0 {
		req.SetParam("max_level", strconv.Itoa(max_level))
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

// Retrieve the details of an item
func (c *ArtifactsMMO) GetItem(code string) (*models.SingleItem, error) {
	var ret models.SingleItem

	res, err := api.NewRequest(c.Config).SetMethod("GET").SetURL(fmt.Sprintf("/items/%s", code)).SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 404 {
		return nil, models.ErrItemNotFound
	}

	return &ret, nil
}
