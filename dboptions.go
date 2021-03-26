/*
--------------------------------------------------
 File Name: dboptions.go
 Author: hanxu

 Created Time: 2019-9-5-下午3:47
---------------------说明--------------------------

---------------------------------------------------
*/

package mydao

import (
	"errors"
)

// ----------start----------dbinstance option impl----------start----------
type Dboption func(opts *dboptions) error

/*type Dboptions interface {
	apply(opts *dboptions)
}*/
type dboptions struct {
	Dbsource Datasource
	Isdebug  bool
	PrintSql bool
}

func WithDebug() Dboption {
	return func(opts *dboptions) error {
		opts.Isdebug = true
		return nil
	}
}

func WithPrintSql() Dboption {
	return func(opts *dboptions) error {
		opts.PrintSql = true
		return nil
	}
}

func WithDataSource(dbs Datasource) Dboption {
	return func(opts *dboptions) error {
		if dbs != nil {
			opts.Dbsource = dbs
			return nil
		}
		return errors.New("this params of Datasouce is empty.")
	}
}

// ----------end------------dbinstance option impl----------end------------
