package service

import (
	"github.com/yoshinorihisakawa/api-taskmanagement/domain/model"
	"github.com/yoshinorihisakawa/api-taskmanagement/usecase/presenter"
	"github.com/yoshinorihisakawa/api-taskmanagement/usecase/repository"
)

type taskUserSetService struct {
	TaskRepository        repository.TaskRepository
	TaskPresenter         presenter.TaskPresenter
	UserRepository        repository.UserRepository
	TaskUserSetRepository repository.TaskUserSetRepository
	TaskUserSetPresenter  presenter.TaskUserSetPresenter
}

type TaskUserSetService interface {
	GetTaskUserSet(id uint) (*model.TaskUserSetResponse, error)
}

func NewTaskUsetSetService(
	tr repository.TaskRepository,
	pre presenter.TaskPresenter,
	ur repository.UserRepository,
	tur repository.TaskUserSetRepository,
	tup presenter.TaskUserSetPresenter,
) TaskUserSetService {
	return &taskUserSetService{
		tr,
		pre,
		ur,
		tur,
		tup,
	}
}

func (taskUserSetService *taskUserSetService) GetTaskUserSet(id uint) (*model.TaskUserSetResponse, error) {
	// FIXME:クエリが汚いからリファクタリングすること
	// taskIdに紐づくtaskUserSetを取得
	taskUserSet, err := taskUserSetService.TaskUserSetRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	// 上記の結果を利用してsendUserを取得
	sendUser, err := taskUserSetService.UserRepository.FindByID(taskUserSet.SendUserID)
	if err != nil {
		return nil, err
	}

	// idにひもづくreceiveUserIdsを取得
	receiveUsersIDs, err := taskUserSetService.TaskUserSetRepository.FindAllIDsByID(id)
	if err != nil {
		return nil, err
	}
	var ids []*int64
	for _, c := range receiveUsersIDs {
		id := int64(c.ReceiveUserID)
		ids = append(ids, &id)
	}

	// 上記の結果を利用して、receiveUserを複数取得
	receiveUsers, err := taskUserSetService.UserRepository.FindAllByID(ids)
	if err != nil {
		return nil, err
	}

	// taskを取得
	task, err := taskUserSetService.TaskRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	// presenterでレスポンスように加工
	return taskUserSetService.TaskUserSetPresenter.ResponseTaskUserSet(task, sendUser, receiveUsers)
}
