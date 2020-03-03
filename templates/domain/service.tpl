package {{.PackageName}}

import "context"

type Service interface {
	Entities(ctx context.Context) ([]string, error)
	Entity(ctx context.Context, UUID string) ([]string, error)
}

type service struct {
	entityService Service
}

func New(s Service) Service {
	return service{
		entityService: s,
	}
}

func (s service) Entities(ctx context.Context) ([]string, error) {
	return nil,nil
}

func (s service) Entity(ctx context.Context, UUID string) ([]string, error) {
	return nil,nil
}