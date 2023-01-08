package repository

import (
	"fmt"
	"os"
	"strings"

	"github.com/golover/bobgen/domain"
)

type DomainRepository struct {
}

func NewDomainRepository() DomainRepository {
	return DomainRepository{}
}
func (DomainRepository) CreateFile(name, pfb string) error {
	err := pathFromBaseNumberOfSlashesValidator(pfb)
	if err != nil {
		return err
	}

	if strings.TrimSpace(pfb) != "" {
		err := os.MkdirAll(pfb, os.ModePerm)
		if err != nil {
			return fmt.Errorf("%v, %w", domain.ErrCreatingFile, err)
		}
	}

	pfb = "./" + pfb
	_, isFileExist := os.Open(pfb + name)
	if isFileExist == nil {
		return domain.ErrFileAlreadyExist
	}

	_, err = os.Create(pfb + name)
	if err != nil {
		return fmt.Errorf("%v, %w", domain.ErrCreatingFile, err)
	}
	return err
}

func (DomainRepository) DeleteFile(pfb string) error {
	err := pathFromBaseNumberOfSlashesValidator(pfb)
	if err != nil {
		return err
	}
	pfb = "./" + pfb
	err = os.RemoveAll(pfb)
	if err != nil {
		return fmt.Errorf("%v, %w", domain.ErrRemovingFile, err)
	}
	return err
}

func pathFromBaseNumberOfSlashesValidator(pfb string) error {
	numberOfslashes := strings.Count(pfb, "/")
	if numberOfslashes%2 == 0 && numberOfslashes != 0 {
		return domain.ErrNumberOfSlashesMustBeOdd
	}
	return nil
}
