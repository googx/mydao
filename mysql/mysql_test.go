/*
--------------------------------------------------
 File Name: mysql.go
 Author: hanxu

 Created Time: 2019-9-5-下午4:10
---------------------说明--------------------------

---------------------------------------------------
*/

package mysql

import (
	"testing"

	"github.com/googx/mydao"
)

func _rightDsOpts() *mydao.DsOptions {
	return &mydao.DsOptions{
		User:   "gorm",
		Passwd: "gorm",
		//
		Host: "localhost",
		Port: 3306,
		//
		DbName: "gormdb",
	}
}

func _errorDsOpts() *mydao.DsOptions {
	return &mydao.DsOptions{}
}

func Test_dnsBuilder_dsn(t *testing.T) {
	type fields struct {
		dbsOpts *mydao.DsOptions
	}
	tests := []struct {
		name    string
		want    string
		f       fields
		wantErr bool
	}{
		{name: "right dsn", f: fields{_rightDsOpts()}, want: "gorm:gorm@tcp(localhost:3306)/gormdb", wantErr: false},
		{name: "error dsn", f: fields{_errorDsOpts()}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &dnsBuilder{
				dbsOpts: tt.f.dbsOpts,
			}
			got, err := b.dsn()
			if (err != nil) != tt.wantErr {
				t.Errorf("dnsBuilder.dsn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("dnsBuilder.dsn() = %v, want %v", got, tt.want)
			}
		})
	}
}
