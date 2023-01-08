package cmd

import (
	"github.com/golover/bobgen/app/domain/delivery/cobra"
	"github.com/golover/bobgen/app/domain/repository"
	"github.com/golover/bobgen/app/domain/usecase"
)

func Boot() {
	r := repository.NewDomainRepository()
	u := usecase.NewDomainUsecase(r)
	d := cobra.NewDomainDelivery(u)
	d.Execute()
}
