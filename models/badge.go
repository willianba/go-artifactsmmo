package models

type BadgeCondition SimpleItem

type Badge struct {
	Code        string           `json:"code"`
	Season      int              `json:"season"`
	Description string           `json:"description"`
	Conditions  []BadgeCondition `json:"conditions"`
}
