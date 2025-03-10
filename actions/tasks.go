package actions

import (
	"fmt"
	"strconv"

	"github.com/0xN0x/go-artifactsmmo/internal/api"
	"github.com/0xN0x/go-artifactsmmo/models"
)

func (c *ArtifactsMMO) GetTasks(skill models.SkillType, task_type models.TaskType, max_level int, min_level int, page int, size int) (*[]models.TaskFull, error) {
	var ret []models.TaskFull
	req := api.NewRequest(c.Config).SetMethod("GET").SetURL("/tasks/list").SetResultStruct(&ret)

	if skill != "" {
		req.SetParam("skill", string(skill))
	}

	if task_type != "" {
		req.SetParam("task_type", string(task_type))
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

func (c *ArtifactsMMO) GetTask(code string) (*models.TaskFull, error) {
	var ret models.TaskFull

	res, err := api.NewRequest(c.Config).SetMethod("GET").SetURL(fmt.Sprintf("/tasks/list/%s", code)).SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 404 {
		return nil, models.ErrTaskNotFound
	}

	return &ret, nil
}

func (c *ArtifactsMMO) GetTasksRewards(code string, page int, size int) (*[]models.TaskRewardFull, error) {
	var ret []models.TaskRewardFull
	req := api.NewRequest(c.Config).SetMethod("GET").SetURL("/tasks/rewards").SetResultStruct(&ret)

	if code != "" {
		req.SetParam("code", code)
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

func (c *ArtifactsMMO) GetTaskReward(code string) (*models.TaskRewardFull, error) {
	var ret models.TaskRewardFull

	res, err := api.NewRequest(c.Config).SetMethod("GET").SetURL(fmt.Sprintf("/tasks/rewards/%s", code)).SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 404 {
		return nil, models.ErrRewardNotFound
	}

	return &ret, nil
}

func (c *ArtifactsMMO) CompleteTask() (*models.TaskRewardData, error) {
	var taskReward models.TaskRewardData

	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/task/complete", c.Config.GetUsername())).SetResultStruct(&taskReward).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 598:
		return nil, models.ErrTaskmasterNotFound
	}

	return &taskReward, nil
}

func (c *ArtifactsMMO) TaskExchange() (*models.TaskRewardData, error) {
	var task models.TaskRewardData

	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/task/exchange", c.Config.GetUsername())).SetResultStruct(&task).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 598:
		return nil, models.ErrTaskmasterNotFound
	}

	return &task, nil
}

func (c *ArtifactsMMO) AcceptNewTask() (*models.TaskData, error) {
	var task models.TaskData

	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/task/new", c.Config.GetUsername())).SetResultStruct(&task).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 598:
		return nil, models.ErrTaskmasterNotFound
	}

	return &task, nil
}

func (c *ArtifactsMMO) TaskTrade(code string, quantity int) (*models.TaskTradeData, error) {
	var task models.TaskTradeData

	body := models.SimpleItem{Code: code, Quantity: quantity}
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/task/trade", c.Config.GetUsername())).SetResultStruct(&task).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 598:
		return nil, models.ErrTaskmasterNotFound
	}

	return &task, nil
}

func (c *ArtifactsMMO) TaskCancel() (*models.TaskCancelled, error) {
	var task models.TaskCancelled

	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/task/cancel", c.Config.GetUsername())).SetResultStruct(&task).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 598:
		return nil, models.ErrTaskmasterNotFound
	}

	return &task, nil
}
