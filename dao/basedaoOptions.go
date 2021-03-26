/*
--------------------------------------------------
 File Name: basedaoOptions.go
 Author: hanxu
 Created Time: 2019-12-30-上午9:54
---------------------说明--------------------------

---------------------------------------------------
*/

package base

import (
	"github.com/googx/mydao"
)

type BaseDaoOption func(opts *BaseDaoOptions) error
type BaseDaoOptions struct {
	DataSouce dbSources.Datasource
	Debug     bool
}

func WithDataSources(ds dbSources.Datasource) BaseDaoOption {
	return func(opts *BaseDaoOptions) error {
		opts.DataSouce = ds
		return nil
	}
}

func WithDebug() BaseDaoOption {
	return func(opts *BaseDaoOptions) error {
		opts.Debug = true
		return nil
	}
}
