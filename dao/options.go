/*
--------------------------------------------------
 File Name: IotOptions.go
 Author: hanxu
 Created Time: 2019-12-30-上午10:38
---------------------说明--------------------------

---------------------------------------------------
*/

package base

import (
	"fmt"
)

type ProjectDao interface {
	Init(opts ...ProjectDaoOption)
}

type ProjectDaoOption func(opts *ProjectDaoOptions) error
type ProjectDaoOptions struct {
	Dao BaseDao
}

func WithBaseDao(dao BaseDao) ProjectDaoOption {
	if dao == nil {
		panic(fmt.Errorf("invlid basedao,is nil"))
	}
	return func(opts *ProjectDaoOptions) error {
		opts.Dao = dao
		return nil
	}
}
