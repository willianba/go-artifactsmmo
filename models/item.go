package models

type ItemType string

const (
	ItemConsumable ItemType = "consumable"
	ItemBodyArmor  ItemType = "body_armor"
	ItemWeapon     ItemType = "weapon"
	ItemResource   ItemType = "resource"
	ItemLegArmor   ItemType = "leg_armor"
	ItemHelmet     ItemType = "helmet"
	ItemBoots      ItemType = "boots"
	ItemShield     ItemType = "shield"
	ItemAmulet     ItemType = "amulet"
	ItemRing       ItemType = "ring"
	ItemArtifact   ItemType = "artifact"
	ItemCurrency   ItemType = "currency"
)

type SkillType string

const (
	Mining      SkillType = "mining"
	WoodCutting SkillType = "woodcutting"
	Fishing     SkillType = "fishing"
)

type Effect struct {
	Code  string `json:"code"`
	Value int    `json:"value"`
}

type Item struct {
	Name        string   `json:"name"`
	Code        string   `json:"code"`
	Level       int      `json:"level"`
	Type        string   `json:"type"`
	SubType     string   `json:"subtype"`
	Description string   `json:"description"`
	Effects     []Effect `json:"effects"`
	Craft       Craft    `json:"craft"`
	Tradeable   bool     `json:"tradeable"`
}

type Craft struct {
	Skill    string       `json:"skill"`
	Level    int          `json:"level"`
	Items    []SimpleItem `json:"items"`
	Quantity int          `json:"quantity"`
}

type SimpleItem struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type Drop SimpleItem

type Detail struct {
	Xp    int    `json:"xp"`
	Items []Drop `json:"items"`
}

type Gold struct {
	Quantity int `json:"quantity"`
}

type ItemsArray struct {
	Items []SimpleItem `json:"items"`
}

type DropFull struct {
	Code        string `json:"code"`
	Rate        int    `json:"rate"`
	MinQuantity int    `json:"min_quantity"`
	MaxQuantity int    `json:"max_quantity"`
}

type Resource struct {
	Name  string     `json:"name"`
	Code  string     `json:"code"`
	Skill SkillType  `json:"skill"`
	Level int        `json:"level"`
	Drops []DropFull `json:"drops"`
}
