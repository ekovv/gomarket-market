package domains

//go:generate go run github.com/vektra/mockery/v3 --name=SessionService
type SessionService interface {
	Generate() (string, error)
}
