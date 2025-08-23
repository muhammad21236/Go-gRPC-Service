//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/muhammad21236/Go-gRPC-Service/internal/rocket Store

package rocket

import "context"

// Package rocket provides functionality related to rockets.
type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

// Store defines the interface for rocket storage operations.
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service provides methods to manage rockets.
type Service struct {
	Store Store
}

// New creates a new instance of the rocket service.
// It initializes the service and returns it.
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

// GetRocketByID retrieves a rocket by its ID.
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// InsertRocket adds a new rocket to the store.
func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// DeleteRocket removes a rocket from the store by its ID.
func (s Service) DeleteRocket(ctx context.Context, id string) error {
	if err := s.Store.DeleteRocket(id); err != nil {
		return err
	}
	return nil
}
