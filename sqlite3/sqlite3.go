/*
--------------------------------------------------
 File Name: sqlite3.go
 Author: hanxu

 Created Time: 2019-9-5-下午4:04
---------------------说明--------------------------

---------------------------------------------------
*/

package sqlite3

import (
	"github.com/googx/mydao"
)

func NewDataSource(opt ...mydao.Option) mydao.Datasource {
	return mydao.NewSqliteDbs(opt...)
}
