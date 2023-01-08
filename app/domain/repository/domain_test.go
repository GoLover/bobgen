package repository

import (
	"testing"
)

func TestDomainRepository_CreateFile(t *testing.T) {
	type args struct {
		name string
		pfb  string
	}
	tests := []struct {
		name    string
		d       DomainRepository
		args    args
		wantErr bool
	}{
		{name: "OK File", d: DomainRepository{}, args: args{name: "testfile"}, wantErr: false},
		{name: "Duplicate File", d: DomainRepository{}, args: args{name: "testfile"}, wantErr: true},
		{name: "File with strange name", d: DomainRepository{}, args: args{name: "%&#%%@# .% . @&#$"}, wantErr: false},
		{name: "File in a directory without pfb", d: DomainRepository{}, args: args{name: "testdir/testfile"}, wantErr: true},
		{name: "File in a directory with pfb", d: DomainRepository{}, args: args{name: "testfile", pfb: "testdir/"}, wantErr: false},
		{name: "File exist in a directory with pfb", d: DomainRepository{}, args: args{name: "testfilea", pfb: "testdir/"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DomainRepository{}
			if err := d.CreateFile(tt.args.name, tt.args.pfb); (err != nil) != tt.wantErr {
				t.Errorf("DomainRepository.CreateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDomainRepository_DeleteFile(t *testing.T) {
	type args struct {
		pfb string
	}
	tests := []struct {
		name    string
		d       DomainRepository
		args    args
		wantErr bool
	}{
		{name: "delete exist directory", d: DomainRepository{}, args: args{pfb: "testdir/"}, wantErr: false},
		{name: "delete not exist directory", d: DomainRepository{}, args: args{pfb: "disthirektoridaznutegzist/"}, wantErr: false},
		{name: "delete exist file", d: DomainRepository{}, args: args{pfb: "testfile"}, wantErr: false},
		{name: "delete exist strange file", d: DomainRepository{}, args: args{pfb: "%&#%%@# .% . @&#$"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DomainRepository{}
			if err := d.DeleteFile(tt.args.pfb); (err != nil) != tt.wantErr {
				t.Errorf("DomainRepository.DeleteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pathFromBaseNumberOfSlashesValidator(t *testing.T) {
	type args struct {
		pfb string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "correct one", args: args{pfb: "testfile/"}, wantErr: false},
		{name: "wrong one", args: args{pfb: "/testfile/"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := pathFromBaseNumberOfSlashesValidator(tt.args.pfb); (err != nil) != tt.wantErr {
				t.Errorf("pathFromBaseNumberOfSlashesValidator() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
