/*
--------------------------------------------------
 File Name: dbconfig.go
 Author: hanxu

 Created Time: 2019-8-31-下午3:49
---------------------说明--------------------------
 得到db对象
---------------------------------------------------
*/

package mydao

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

type DbGormOption func(opts *dbGormOptions) error
type dbGormOptions struct {
	db      *gorm.DB
	IsDebug bool
}
type DbGorm struct {
	Opts     *dboptions
	gormOpts *dbGormOptions
	//
	dbinstance *gorm.DB
}

// TODO 这些单例获取db 数据源这种操作 可以参考下, 使用sync.Once包实现
func (this *DbGorm) GetDb() (*gorm.DB, error) {
	if this.dbinstance != nil {
		return this.dbinstance, nil
	}
	dbs := this.Opts.Dbsource
	dsn, err := dbs.Dsn()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error dsn:%v", err))
	}
	db, err := gorm.Open(dbs.Name(), dsn)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("open database has failed.%v", err))
	}
	this.dbinstance = db
	return db, err
}
func (this *DbGorm) Init(opts ...DbGormOption) {
	var defopts *dbGormOptions
	if this.gormOpts != nil {
		defopts = this.gormOpts
	} else {
		defopts = &dbGormOptions{
			db:      this.dbinstance,
			IsDebug: false,
		}
	}
	//
	for _, x := range opts {
		x(defopts)
	}
	this.gormOpts = defopts
	this.dbinstance = defopts.db
}

func (this *DbGorm) Copy() *DbGorm {
	dbGorm := DbGorm{
		Opts:     this.Opts,
		gormOpts: this.gormOpts,
	}
	return &dbGorm
}

func WithGormDB(db *gorm.DB) DbGormOption {
	return func(opts *dbGormOptions) error {
		opts.db = db
		return nil
	}
}

func (this *DbGorm) GetInitDb(tableSchema interface{}) (*gorm.DB, error) {
	db, e := this.GetDb()
	if e != nil {
		return nil, e
	} else {
		// 只适合用来做一些前缀和后缀.无法用来实现自定义表名
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			// log.Printf("设置表名==>%v  \n", defaultTableName)
			return defaultTableName
		}
		hastab := db.HasTable(tableSchema)

		if !hastab {
			db.CreateTable(tableSchema)
		}
		// TODO 已存在, 查看配置 是否需要更新结构
		dbnew := db.AutoMigrate(tableSchema)
		if this.gormOpts.IsDebug {
			dbnew = dbnew.Debug()
		}
		return dbnew, nil
	}
}

func getDefaultOption() dboptions {
	defaultDbOption := dboptions{
		//Dbsource: NewSqliteDbs(),
		Isdebug: false,
	}
	return defaultDbOption
}

func NewDbGorm(opt ...Dboption) *DbGorm {
	defopts := getDefaultOption()
	for _, x := range opt {
		/*x.apply(&defopts)*/
		x(&defopts)
	}
	dborm := new(DbGorm)
	dborm.Opts = &defopts
	// check
	return dborm
}
