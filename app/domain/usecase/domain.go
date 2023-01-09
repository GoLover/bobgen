package usecase

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/dave/jennifer/jen"
	"github.com/golover/bobgen/domain"
)

type DomainUsecase struct {
	dr domain.GenericRepository
}

func NewDomainUsecase(dr domain.GenericRepository) DomainUsecase {
	return DomainUsecase{dr: dr}
}

func (du DomainUsecase) CreateDomain(dn string) error {
	dn = correctLetterAndNumber(dn)
	err := du.dr.CreateFile(dn+".go", "domain/")
	err = du.dr.WriteFile(dn+".go", "domain/", []byte("package domain\n"))
	err = du.dr.CreateFile(dn+".go", "app/"+dn+"/delivery/")
	err = du.dr.WriteFile(dn+".go", "app/"+dn+"/delivery/", []byte("package delivery\n"))
	err = du.dr.CreateFile(dn+".go", "app/"+dn+"/repository/")
	err = du.dr.WriteFile(dn+".go", "app/"+dn+"/repository/", []byte("package repository\n"))
	err = du.dr.CreateFile(dn+".go", "app/"+dn+"/usecase/")
	err = du.dr.WriteFile(dn+".go", "app/"+dn+"/usecase/", []byte("package usecase\n"))
	du.AddDURToDomain(dn, dn)
	return err
}
func (du DomainUsecase) addToDomain(dn string, content []byte) error {
	err := du.dr.WriteFile(dn+".go", "domain/", content)
	return err
}
func (du DomainUsecase) AddDURToDomain(dn, un string) error {
	err := du.AddDeliveryToDomain(dn, un)
	err = du.AddRepositoryToDomain(dn, un)
	err = du.AddUsecaseToDomain(dn, un)
	return err
}
func (du DomainUsecase) AddDeliveryToDomain(dn, deln string) error {
	dn = correctLetterAndNumber(dn)
	deln = correctLetterAndNumber(deln)
	err := du.dr.WriteFile(deln+".go", "app/"+dn+"/delivery/", []byte(fmt.Sprintf("%#v\n", jen.Type().Id(deln+"Delivery").Struct())))
	err = du.addToDomain(dn, []byte(fmt.Sprintf("%#v\n", jen.Type().Id(deln+"Delivery").Interface())))
	return err
}
func (du DomainUsecase) AddUsecaseToDomain(dn, un string) error {
	dn = correctLetterAndNumber(dn)
	un = correctLetterAndNumber(un)
	err := du.dr.WriteFile(un+".go", "app/"+dn+"/usecase/", []byte(fmt.Sprintf("%#v\n", jen.Type().Id(un+"Usecase").Struct())))
	err = du.addToDomain(dn, []byte(fmt.Sprintf("%#v\n", jen.Type().Id(un+"Usecase").Interface())))
	return err
}
func (du DomainUsecase) AddRepositoryToDomain(dn, rn string) error {
	dn = correctLetterAndNumber(dn)
	rn = correctLetterAndNumber(rn)
	err := du.dr.WriteFile(rn+".go", "app/"+dn+"/repository/", []byte(fmt.Sprintf("%#v\n", jen.Type().Id(rn+"Repository").Struct())))
	err = du.addToDomain(dn, []byte(fmt.Sprintf("%#v\n", jen.Type().Id(rn+"Repository").Interface())))
	return err
}
func correctLetterAndNumber(dn string) string {
	dn = strings.TrimSpace(dn)
	var correctedStr string
	makeItUpper := true
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
