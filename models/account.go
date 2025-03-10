package models

type AccountStatus string

const (
	Standard    AccountStatus = "standard"
	Founder     AccountStatus = "founder"
	GoldFounder AccountStatus = "gold_founder"
	VIPFounder  AccountStatus = "vip_founder"
)

type AccountDetails struct {
	Username          string        `json:"username"`
	Email             string        `json:"email"`
	Subscribed        bool          `json:"subscribed"`
	Status            AccountStatus `json:"status"`
	Badges            []string      `json:"badges"`
	Gems              int           `json:"gems"`
	AchivementsPoints int           `json:"achivements_points"`
	Banned            bool          `json:"banned"`
	BanReason         string        `json:"ban_reason"`
}
