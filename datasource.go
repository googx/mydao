/*
--------------------------------------------------
 File Name: datasource.go
 Author: hanxu

 Created Time: 2019-9-5-下午3:32
---------------------说明--------------------------

---------------------------------------------------
*/

package mydao

import (
	"database/sql"
)

type Datasource interface {
	Name() string
	Dsn() (string, error)
	Connect() *sql.DB
}
