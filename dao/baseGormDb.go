/*
--------------------------------------------------
 File Name: baseGormDb.go
 Author: hanxu
 Created Time: 2019-12-30-下午1:42
---------------------说明--------------------------

---------------------------------------------------
*/

package base

import (
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/googx/mydao"
	"github.com/googx/mydao/internal"

	"github.com/googx/mydao/model"
)

type BaseGormDB struct {
	dbGOrm *mydao.DbGorm
}

func (bgd *BaseGormDB) Name() string {
	return "BaseGormDB"
}

func NewBaseGormDB(dboptions ...mydao.Dboption) *BaseGormDB {
	dbGOrm := mydao.NewDbGorm(dboptions...)

	_db, e := dbGOrm.GetDb()
	if e != nil {
		panic(e)
	}
	dborm_callback(_db)
	if dbGOrm.Opts.Isdebug {
		dborm_debug_callback(_db)
		_db = _db.Debug()
	}

	dbGOrm.Init(
		mydao.WithGormDB(_db),
	)
	return &BaseGormDB{
		dbGOrm: dbGOrm,
	}
}

func (bgd *BaseGormDB) Table(entityScheml DbEntity) *BaseGormDB {

	if db, e := bgd.dbGOrm.GetInitDb(entityScheml); e != nil {
		panic(e)
		// return nil
	} else {
		dbGorm := bgd.dbGOrm.Copy()
		dbGorm.Init(
			mydao.WithGormDB(db),
		)
		return &BaseGormDB{
			dbGorm,
		}
	}
}

func (bgd *BaseGormDB) GetGormDB() (*gorm.DB, error) {
	return bgd.dbGOrm.GetDb()
}

// TODO 系统了解和学习下 gorm的callback机制, https://www.cnblogs.com/sgyBlog/p/10154424.html
func dborm_callback(db *gorm.DB) {
	// 原生定义的代码在这里: pkg/mod/github.com/jinzhu/gorm@v1.9.10/callback_(create|delete|query|save|).go:init()
	db = db.LogMode(false)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	// db.Callback().Delete().Replace("gorm:delete", deleteCallback)
}

func dborm_debug_callback(db *gorm.DB) {
	db.Callback().Create().Before("gorm:before_create").Register("gorm:create_json_field", jsonStringCallback)
}

// 改变 原有的创建时间的机制
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {

		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(model.NowJsonTime())
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(model.NowJsonTime())
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdatedAt", model.NowJsonTime())
	}
}

// 开启会耗性能,应该是debug模式开启,需要优化结构
// 当使用string来存储json,以后可以做些通用的优化操作, json检查,美化,解析扫描结构体
func jsonStringCallback(scope *gorm.Scope) {
	fields := scope.Fields()
	for _, v := range fields {
		// log.Printf("v.Name==>%s", v.Name)
		if s, b := v.TagSettingsGet(strings.ToUpper("format")); b {
			switch s {
			case "json":
				if !v.IsBlank {
					switch v.Field.Kind() {
					case reflect.String:
						js := v.Field.String()
						if prettyjs, e := internal.PrettyJson(js); e == nil {
							v.Field.SetString(prettyjs)
						}
					}
				}
			}
		}
	}
}
