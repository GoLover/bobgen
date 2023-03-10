package domain

type Medium int

const (
	GRPC Medium = iota
	HTTP
	Kafka
)

type Domain struct {
	name       string
	delivery   delivery
	repository repository
	usecase    usecase
}

type usecase struct {
	name string
	args map[string]string
}

type delivery struct {
	name string
	kind Medium
}

type repository struct {
	name string
}

type DomainUseCase interface {
	CreateDomain(name string) error
	AddDURToDomain(dn, un string) error
	AddDeliveryToDomain(dn, un string) error
	AddRepositoryToDomain(dn, un string) error
	AddUsecaseToDomain(dn, un string) error
	//DeleteDomain(name string) error
}
type GenericRepository interface {
	CreateFile(name, pathFromBase string) error
	WriteFile(name, pathFromBase string, content []byte) error
	ReadFile(name, pathFromBase string) ([]byte, error)
	DeleteFile(pathFromBase string) error
}
