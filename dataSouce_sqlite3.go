/*
--------------------------------------------------
 File Name: sqlite3DataSouces.go
 Author: hanxu

 Created Time: 2019-9-5-下午2:30
---------------------说明--------------------------

---------------------------------------------------
*/

package mydao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type sqliteDbs struct {
	dbfile string
}

func (*sqliteDbs) Name() string {
	return "sqlite3"
}

func (this *sqliteDbs) Dsn() (string, error) {
	return this.dbfile, nil
}

func (*sqliteDbs) Connect() *sql.DB {
	panic("implement me")
}

func (*sqliteDbs) GetDb() *sql.DB {
	panic("implement me")
}

func NewSqliteDbs(ops ...Option) *sqliteDbs {
	return newSqliteDbs(ops...)
}

func newSqliteDbs(ops ...Option) *sqliteDbs {
	dbopts := defaultOptions
	for _, x := range ops {
		x(&dbopts)
	}
	dbs := &sqliteDbs{
		dbfile: dbopts.dbfile,
	}
	log.Printf("[sqlite3] using db file==> %v  \n", dbs.dbfile)
	return dbs
}

// ----------start----------options----------start----------
var defaultOptions = Sqlite3Options{
	// 要分隔下os.args[0]
	// /tmp/sqlite3//tmp/___TestUserInfoDao_SaveUser_in_t4_gorm_c4_associated.dbinstance
	dbfile: setdbFile(fmt.Sprintf("/tmp/sqlite3/%s.dbinstance", getDefaultDbName())),
}

func getDefaultDbName() string {
	exename := os.Args[0]
	_, filename := filepath.Split(exename)
	if filename != "" {
		_ext := filepath.Ext(filename)
		if _ext != "" {
			return strings.Replace(filename, _ext, "", 1)
		}
		return filename
	}
	return "__dbsources"
}

func setdbFile(dbfile string) string {
	dir, file := filepath.Split(dbfile)
	// mkdirerr := os.MkdirAll(dir, 0700)
	mkdirerr := os.MkdirAll(dir, 0700)
	if mkdirerr != nil {
		log.Fatalf("创建数据库目录失败%v \n ", mkdirerr)
	}
	return filepath.Join(dir, file)
}

type Sqlite3Options struct {
	dbfile string
}
type Option func(_db *Sqlite3Options)

func WithDbFile(dbfile string) Option {
	return func(_db *Sqlite3Options) {
		if dbfile == "" {
			return
		}
		_db.dbfile = setdbFile(dbfile)
	}
}

// ----------end------------options----------end------------
