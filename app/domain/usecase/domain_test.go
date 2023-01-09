package usecase

import (
	"reflect"
	"testing"

	"github.com/golover/bobgen/app/domain/repository"
	"github.com/golover/bobgen/domain"
)

func TestDomainUsecase_CreateDomain(t *testing.T) {
	type fields struct {
		dr domain.GenericRepository
	}
	type args struct {
		dn string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "create domain", fields: fields{dr: repository.NewDomainRepository()}, args: args{dn: "test"}, wantErr: false},
		{name: "create another domain", fields: fields{dr: repository.NewDomainRepository()}, args: args{dn: "anothertest"}, wantErr: false},
		{name: "create wierd domain", fields: fields{dr: repository.NewDomainRepository()}, args: args{dn: "$%&YHDF%#GFJ%"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			du := DomainUsecase{
				dr: tt.fields.dr,
			}
			if err := du.CreateDomain(tt.args.dn); (err != nil) != tt.wantErr {
				t.Errorf("DomainUsecase.CreateDomain() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewDomainUsecase(t *testing.T) {
	tests := []struct {
		name string
		want DomainUsecase
	}{
		{name: "add bulik domain", want: NewDomainUsecase(repository.NewDomainRepository())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDomainUsecase(repository.NewDomainRepository()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDomainUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDomainUsecase_AddDURToDomain(t *testing.T) {
	type fields struct {
		dr domain.GenericRepository
	}
	type args struct {
		dn string
		un string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "add bulik domain to exist domain", fields: fields{dr: repository.NewDomainRepository()}, args: args{dn: "test", un: "domainbydur"}, wantErr: false},
		{name: "add bulik domain to not exist domain", fields: fields{dr: repository.NewDomainRepository()}, args: args{dn: "testor", un: "domainbydur"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			du := DomainUsecase{
				dr: tt.fields.dr,
			}
			if err := du.AddDURToDomain(tt.args.dn, tt.args.un); (err != nil) != tt.wantErr {
				t.Errorf("DomainUsecase.AddDURToDomain() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDomainUsecase_AddDeliveryToDomain(t *testing.T) {
	type fields struct {
		dr domain.GenericRepository
	}
	type args struct {
		dn string
		un string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "add delivery to domain", fields: fields{dr: repository.NewDomainRepository()}, args: args{un: "DeliveryA", dn: "test"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			du := DomainUsecase{
				dr: tt.fields.dr,
			}
			if err := du.AddDeliveryToDomain(tt.args.dn, tt.args.un); (err != nil) != tt.wantErr {
				t.Errorf("DomainUsecase.AddDeliveryToDomain() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDomainUsecase_AddUsecaseToDomain(t *testing.T) {
	type fields struct {
		dr domain.GenericRepository
	}
	type args struct {
		dn string
		un string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "add usecase to domain", fields: fields{dr: repository.NewDomainRepository()}, args: args{un: "UsecaseA", dn: "test"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			du := DomainUsecase{
				dr: tt.fields.dr,
			}
			if err := du.AddUsecaseToDomain(tt.args.dn, tt.args.un); (err != nil) != tt.wantErr {
				t.Errorf("DomainUsecase.AddUsecaseToDomain() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDomainUsecase_AddRepositoryToDomain(t *testing.T) {
	type fields struct {
		dr domain.GenericRepository
	}
	type args struct {
		dn string
		un string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "add repository to domain", fields: fields{dr: repository.NewDomainRepository()}, args: args{un: "RepoA", dn: "test"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			du := DomainUsecase{
				dr: tt.fields.dr,
			}
			if err := du.AddRepositoryToDomain(tt.args.dn, tt.args.un); (err != nil) != tt.wantErr {
				t.Errorf("DomainUsecase.AddRepositoryToDomain() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
