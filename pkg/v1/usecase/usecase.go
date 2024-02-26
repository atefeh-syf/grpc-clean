package usecase

import (
	"errors"

	"github.com/atefeh-syf/grpc-clean/internal/models"
	interfaces "github.com/atefeh-syf/grpc-clean/pkg/v1"
	//userRepo "github.com/atefeh-syf/grpc-clean/pkg/v1/repository"
	"gorm.io/gorm"
)

type UseCase struct {
	repo interfaces.RepoInterface
}

func New(repo interfaces.RepoInterface) interfaces.UseCaseInterface {
	return &UseCase{repo}
}

func (uc *UseCase) Create(user models.User) (models.User, error) {
	return uc.repo.Create(user)
}

func (uc *UseCase) Get(id string) (models.User, error) {
	var user models.User
	var err error

	if user, err = uc.repo.Get(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, err
		}
		return models.User{}, err
	}
	return user, nil
}

func (uc *UseCase) Update(id string,user models.User) (models.User, error) {
	var err error

	if _, err = uc.repo.Get(id); err != nil {
		return user, err
	}

	err = uc.repo.Update(id, user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (uc *UseCase) Delete(id string) error {
	var err error

	if _, err = uc.Get(id); err != nil {
		return err
	}

	err = uc.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
