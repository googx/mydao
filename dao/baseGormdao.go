/*
--------------------------------------------------
 File Name: basedao.go
 Author: hanxu

 Created Time: 2019-9-2-上午9:35
---------------------说明--------------------------

---------------------------------------------------
*/

package base

import (
	"github.com/googx/mydao"
)

type BaseGormDao struct {
	opts   *BaseDaoOptions
	rootDb *BaseGormDB
}

func NewBaseOrmDao(opts ...BaseDaoOption) *BaseGormDao {
	defOpts := &BaseDaoOptions{}
	for _, x := range opts {
		if e := x(defOpts); e != nil {
			panic(e)
		}
	}
	basedao := &BaseGormDao{
		opts: defOpts,
	}
	basedao.init()
	return basedao
}
func (bd *BaseGormDao) init() {
	dbormOpts := []mydao.Dboption{
		mydao.WithPrintSql(),
	}
	if bd.opts.Debug {
		dbormOpts = append(dbormOpts,
			mydao.WithDebug(),
		)
	}
	if bd.opts.DataSouce != nil {
		dbormOpts = append(dbormOpts, mydao.WithDataSource(bd.opts.DataSouce))
	}

	//
	bd.rootDb = NewBaseGormDB(dbormOpts...)
}

func (this *BaseGormDao) InitTableSchema(ent DbEntity) (BaseDB, error) {
	return this.rootDb.Table(ent), nil
}

/*func getDataSouces() mydao.Datasource {

	dbfilepath := "/data/gfyt/sqlite3db"
	dbfilename := "iot.db"
	if strings.EqualFold(consts.GFIOT_Default_VERSION, consts.GFIOT_VERSION)==false {
		dbfilename = fmt.Sprintf("iot.%s.db", consts.GFIOT_VERSION)
	}
	// 返回nil默认是选择sqlite3
	return mydao.NewSqliteDbs(
		mydao.WithDbFile(filepath.Join(dbfilepath, dbfilename)),
	)
	// return nil
}*/
