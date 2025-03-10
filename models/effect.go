package models

type EffectType string

const (
	EquipmentEffect  EffectType = "equipment"
	ConsumableEffect EffectType = "consumable"
	CombatEffect     EffectType = "combat"
)

type EffectSubtype string

const (
	StatSubtype      EffectSubtype = "stat"
	OtherSubtype     EffectSubtype = "other"
	HealSubtype      EffectSubtype = "heal"
	BuffSubtype      EffectSubtype = "buff"
	DebuffSubtype    EffectSubtype = "debuff"
	SpecialSubtype   EffectSubtype = "special"
	GatheringSubtype EffectSubtype = "gathering"
	TeleportSubtype  EffectSubtype = "teleport"
	GoldSubtype      EffectSubtype = "gold"
)

type Effect struct {
	Name        string        `json:"name"`
	Code        string        `json:"code"`
	Description string        `json:"description"`
	Type        EffectType    `json:"type"`
	Subtype     EffectSubtype `json:"subtype"`
}
