/*
--------------------------------------------------
 File Name: basedao.go
 Author: hanxu
 Created Time: 2019-12-30-上午11:19
---------------------说明--------------------------

---------------------------------------------------
*/

package base

type BaseDao interface {
	InitTableSchema(ent DbEntity) (BaseDB, error)
}

// db接口应该实现基本的单表crud和支持sql复杂查询以及事物.
type BaseDB interface {
	Name() string
}
