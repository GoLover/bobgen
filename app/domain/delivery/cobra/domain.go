package cobra

import (
	"github.com/golover/bobgen/domain"
	"github.com/spf13/cobra"
)

type DomainDelivery struct {
	du      domain.DomainUseCase
	rootCmd *cobra.Command
}

func NewDomainDelivery(du domain.DomainUseCase) DomainDelivery {
	rootCmd := &cobra.Command{
		Use:   "bof-gen",
		Short: "A generator for Band of Blinds project structure",
		Long:  `Any team need a tool for generating bulk project structures, band of blinds isn't exception too.`,
	}
	return DomainDelivery{du: du, rootCmd: rootCmd}
}
func (dd DomainDelivery) Execute() {

	addDomainCmd := &cobra.Command{
		Use:   "domain [domainNames]",
		Short: "adding domain with delivery,usecase,repository",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, k := range args {
				err := dd.du.CreateDomain(k)
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	dd.rootCmd.AddCommand(addDomainCmd)
	dd.rootCmd.Execute()
}
