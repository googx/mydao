/*
--------------------------------------------------
 File Name: common.go
 Author: hanxu

 Created Time: 2019-8-30-下午6:23
---------------------说明--------------------------

---------------------------------------------------
*/

package mydao

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/googx/mydao/model"
)

type IdModule interface {
	GetId() model.ModelId
	GetModelTime() model.ModelTime
}

type DbEntity interface {
	GetType() interface{}
}

func GetSqliteDb(dbfile string) *gorm.DB {
	db, err := gorm.Open("sqlite3", dbfile)
	if err != nil {
		panic(err)
	}

	return db
}

/*func GetInitDb(ent DbEntity) (*gorm.DB, error) {
	dbinstance := GetDb()
	tableSchema := ent.GetType()
	if tableSchema == nil {
		return nil, errors.New("IotEntity->GetType 未实现 ")
	}
	hastab := dbinstance.HasTable(tableSchema)

	if !hastab {
		// 表不存在
		// 创建表
		dbinstance.CreateTable(tableSchema)
	}
	// TODO 已存在, 查看配置 是否需要更新结构
	dbnew := dbinstance.AutoMigrate(tableSchema)
	return dbnew, nil

}*/

// 主键id
type ModelId struct {
	ID uint `gorm:"primary_key"`
}

// 创建和修改时间
type ModelTime struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 软删除
type ModelSofeDel struct {
	DeletedAt *time.Time `sql:"index"`
}
