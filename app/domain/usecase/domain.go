package usecase

import (
	"strings"
	"unicode"

	"github.com/golover/bobgen/domain"
)

type DomainUsecase struct {
	dr domain.GenericRepository
}

func NewDomainUsecase() DomainUsecase {
	return DomainUsecase{}
}

func (du DomainUsecase) CreateDomain(dn string) error {
	du.AddDURToDomain(dn, dn)
	dn = correctLetterAndNumber(dn)
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
	dn = correctLetterAndNumber(dn)
	un = correctLetterAndNumber(un)
	return du.dr.CreateFile(un+".go", "app/"+dn+"/delivery/")
}
func (du DomainUsecase) AddUsecaseToDomain(dn, un string) error {
	dn = correctLetterAndNumber(dn)
	un = correctLetterAndNumber(un)
	return du.dr.CreateFile(un+".go", "app/"+dn+"/usecase/")
}
func (du DomainUsecase) AddRepositoryToDomain(dn, un string) error {
	dn = correctLetterAndNumber(dn)
	un = correctLetterAndNumber(un)
	return du.dr.CreateFile(un+".go", "app/"+dn+"/repository/")
}
func correctLetterAndNumber(dn string) string {
	dn = strings.TrimSpace(dn)
	var correctedStr string
	makeItUpper := false
	for _, k := range dn {
		b := k
		if !unicode.IsLetter(b) && !unicode.IsNumber(b) {
			makeItUpper = true
			continue
		}
		if makeItUpper {
			b = unicode.ToUpper(b)
			makeItUpper = false
		}
		correctedStr += string(b)
	}
	return correctedStr
}
