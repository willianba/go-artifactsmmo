package models

type TaskType string

const (
	TaskTypeMonsters TaskType = "monsters"
	TaskTypeItems    TaskType = "items"
)

type Task struct {
	Code    string     `json:"code"`
	Type    TaskType   `json:"type"`
	Total   int        `json:"total"`
	Rewards TaskReward `json:"rewards"`
}

type TaskReward struct {
	Items []SimpleItem `json:"items"`
	Gold  int          `json:"gold"`
}

type TaskTrade SimpleItem
