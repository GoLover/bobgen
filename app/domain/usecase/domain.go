package usecase

import "github.com/golover/bobgen/domain"

type DomainUsecase struct {
	dr domain.GenericRepository
}

func NewDomainUsecase() DomainUsecase {
	return DomainUsecase{}
}

func (du DomainUsecase) CreateDomain(dn string) error {
	du.AddDURToDomain(dn, dn)
	err := du.dr.CreateFile(dn+".go", "domain/")
	return err
}
func (du DomainUsecase) AddDURToDomain(dn, un string) error {
	err := du.AddDeliveryToDomain(dn, un)
	err = du.AddRepositoryToDomain(dn, un)
	err = du.AddUsecaseToDomain(dn, un)
	return err
}
func (du DomainUsecase) AddDeliveryToDomain(dn, un string) error {
	return du.dr.CreateFile(un+".go", "app/"+dn+"/delivery/")
}
func (du DomainUsecase) AddUsecaseToDomain(dn, un string) error {
	return du.dr.CreateFile(un+".go", "app/"+dn+"/usecase/")
}
func (du DomainUsecase) AddRepositoryToDomain(dn, un string) error {
	return du.dr.CreateFile(un+".go", "app/"+dn+"/repository/")
}
