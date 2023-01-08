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

type GenericUseCases interface {
	Initiate() error
	CreateDomain(name string) error
	DeleteDomain(name string) error
}
type GenericRepository interface {
	CreateFile(name, pathFromBase string) error
	DeleteFile(pathFromBase string) error
}
