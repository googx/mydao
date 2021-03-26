/*
--------------------------------------------------
 File Name: baseEntity.go
 Author: hanxu

 Created Time: 2019-9-27-下午5:32
---------------------说明--------------------------

---------------------------------------------------
*/

package base

import (
	"fmt"

	"github.com/googx/mydao"
	"github.com/googx/mydao/model"
)

type SimpleEntity struct {
	// 记录id
	model.ModelId

	// 记录操作时间
	model.ModelTime
}

func (se SimpleEntity) GetId() model.ModelId {
	return se.ModelId
}

func (se SimpleEntity) GetModelTime() model.ModelTime {
	return se.ModelTime
}

type DbEntity interface {
	fmt.Stringer
	mydao.DbEntity
}

/*func GetDb(ds dbSources.Datasource, ent DbEntity) (*gorm.DB, error) {
	return dao.GetDb(ds, ent)
}*/

func ToString(e DbEntity) string {
	return fmt.Sprintf("%#v", e)
}

type Jsoner interface {
	ToJson() string
}
