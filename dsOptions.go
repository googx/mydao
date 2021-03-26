/*
--------------------------------------------------
 File Name: dsOptions.go
 Author: hanxu

 Created Time: 2019-9-5-下午5:47
---------------------说明--------------------------
 数据源选项
---------------------------------------------------
*/

package mydao

type DsOption func(opt *DsOptions)
type DsOptions struct {
	Host    string
	Port    int
	Charset string
	DbName  string
	User    string
	Passwd  string
}

// ----------start----------option functions----------start----------
func WithHost(host string) DsOption {
	return func(opt *DsOptions) {
		opt.Host = host
	}
}
func WithPort(port int) DsOption {
	return func(opt *DsOptions) {
		opt.Port = port
	}
}

func WithCharset(charset string) DsOption {
	return func(opt *DsOptions) {
		opt.Charset = charset
	}
}

func WithDbName(dbname string) DsOption {
	return func(opt *DsOptions) {
		opt.DbName = dbname
	}
}

func WithUser(user string) DsOption {
	return func(opt *DsOptions) {
		opt.User = user
	}
}

func WithPasswd(passwd string) DsOption {
	return func(opt *DsOptions) {
		opt.Passwd = passwd
	}
}

// ----------end------------option functions----------end------------
