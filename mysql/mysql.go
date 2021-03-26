/*
--------------------------------------------------
 File Name: mysql.go
 Author: hanxu

 Created Time: 2019-9-5-下午4:10
---------------------说明--------------------------

---------------------------------------------------
*/

package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"text/template"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/googx/mydao"
)

func NewDataSource(opts ...mydao.DsOption) mydao.Datasource {
	defaultopts := &mydao.DsOptions{
		Host:    "localhost",
		Port:    3306,
		Charset: "utf8mb4",
	}
	for _, x := range opts {
		x(defaultopts)
	}
	return &mysqlDbs{
		dbsOpts: defaultopts,
	}
}

type mysqlDbs struct {
	dbsOpts *mydao.DsOptions
}

func (this *mysqlDbs) Name() string {
	return "mysql"
}

func (this *mysqlDbs) Dsn() (string, error) {
	// "user:password@tcp(127.0.0.1:3306)/test?maxconnsize=10",
	// "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local")
	b := dnsBuilder{this.dbsOpts}
	return b.dsn()
}

func (this *mysqlDbs) Connect() *sql.DB {
	panic("implement me")
}

type dnsBuilder struct {
	dbsOpts *mydao.DsOptions
}

//go:generate go-bindata -nometadata -pkg mysql -o assets.go mysqldsn.tmpl
func (b *dnsBuilder) dsn() (string, error) {
	// 使用golang的模板来实现,简洁又高效.
	//
	/*mycfg := mysql.NewConfig()
	  mycfg.User = "gosql"
	  mycfg.Passwd = "xxgosqlxx"
	  mycfg.Addr = "localhost:3306"
	  mycfg.DBName = "gosql"
	  mycfg.Net = "tcp"
	  dsninfo := mycfg.FormatDSN()*/
	templ, err := getTemplate("mysqldsn")
	if err != nil {
		return "", err
	}
	_data := map[string]string{
		"user":   b.dbsOpts.User,
		"passwd": b.dbsOpts.Passwd,
		"host":   b.dbsOpts.Host,
		"port":   strconv.Itoa(b.dbsOpts.Port),
		"dbName": b.dbsOpts.DbName,
	}
	sb := strings.Builder{}

	err = templ.Execute(&sb, _data)
	if err != nil {
		return "", err
	}
	if sb.Len() == 0 {
		return "", errors.New("empty dns")
	}
	log.Printf("%s", sb.String())
	return sb.String(), nil
}

func getTemplate(name string) (templ *template.Template, err error) {
	templ = template.New(name)
	i := fmt.Sprintf("%s.tmpl", name)
	textb, _ := Asset(i)
	templ, e := templ.Parse(string(textb))
	if e != nil {
		return nil, errors.New(fmt.Sprintf("解析模板失败%s,", e))
	}
	return templ, nil
}

func (b *dnsBuilder) userPwdHandler(dsn string) (string, error) {
	userpwd := fmt.Sprintf("%s:%s", b.dbsOpts.User, b.dbsOpts.Passwd)
	return userpwd, nil
}
