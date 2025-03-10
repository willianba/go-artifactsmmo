package models

type CharacterLeaderboardSort string

const (
	CombatSort          CharacterLeaderboardSort = "combat"
	WoodcuttingSort     CharacterLeaderboardSort = "woodcutting"
	MiningSort          CharacterLeaderboardSort = "mining"
	FishingSort         CharacterLeaderboardSort = "fishing"
	WeaponCraftingSort  CharacterLeaderboardSort = "weaponcrafting"
	GearCraftingSort    CharacterLeaderboardSort = "gearcrafting"
	JewelryCraftingSort CharacterLeaderboardSort = "jewelrycrafting"
	CookingSort         CharacterLeaderboardSort = "cooking"
	AlchemySort         CharacterLeaderboardSort = "alchemy"
)

type CharacterLeaderboard struct {
	Position               int           `json:"position"`
	Name                   string        `json:"name"`
	Account                string        `json:"account"`
	Skin                   CharacterSkin `json:"skin"`
	Level                  int           `json:"level"`
	TotalXp                int           `json:"total_xp"`
	MiningLevel            int           `json:"mining_level"`
	MiningTotalXp          int           `json:"mining_total_xp"`
	WoodcuttingLevel       int           `json:"woodcutting_level"`
	WoodcuttingTotalXp     int           `json:"woodcutting_total_xp"`
	FishingLevel           int           `json:"fishing_level"`
	FishingTotalXp         int           `json:"fishing_total_xp"`
	WeaponCraftingLevel    int           `json:"weaponcrafting_level"`
	WeaponCraftingTotalXp  int           `json:"weaponcrafting_total_xp"`
	GearCraftingLevel      int           `json:"gearcrafting_level"`
	GearCraftingTotalXp    int           `json:"gearcrafting_total_xp"`
	JewelryCraftingLevel   int           `json:"jewelrycrafting_level"`
	JewelryCraftingTotalXp int           `json:"jewelrycrafting_total_xp"`
	CookingLevel           int           `json:"cooking_level"`
	CookingTotalXp         int           `json:"cooking_total_xp"`
	AlchemyLevel           int           `json:"alchemy_level"`
	AlchemyTotalXp         int           `json:"alchemy_total_xp"`
	Gold                   int           `json:"gold"`
}

type AccountLeaderboardSort string

const (
	AchievementsPointsSort AccountLeaderboardSort = "achievements_points"
	GoldSort               AccountLeaderboardSort = "gold"
)

type AccountLeaderboard struct {
	Position           int           `json:"position"`
	Account            string        `json:"account"`
	Status             AccountStatus `json:"status"`
	AchievementsPoints int           `json:"achievements_points"`
	Gold               int           `json:"gold"`
}
