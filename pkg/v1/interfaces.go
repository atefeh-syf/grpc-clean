package v1

import "github.com/atefeh-syf/grpc-clean/internal/models"

type RepoInterface interface {
	Create(models.User) (models.User ,error)
	Get(id string) (models.User, error)
	Update(id string, user models.User)  error
	Delete(id string) error
}

type UseCaseInterface interface {
	Create(models.User) (models.User ,error)
	Get(id string) (models.User, error)
	Update(id string, user models.User) (models.User ,error)
	Delete(id string) error
}
