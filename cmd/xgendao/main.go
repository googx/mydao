/*
--------------------------------------------------
 File Name: main.go
 Author: hanxu

 Created Time: 2019-9-9-下午5:58
---------------------说明--------------------------
 按照模板自动dao的代码
---------------------------------------------------
*/

package main

//go:generate go install

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/googx/mydao/cmd/xgendao/assets"
	"github.com/googx/mydao/internal"
)

type BootArgs struct {
	// 项目名称
	xprojectName string
	// 实体目录的位置
	xentityPath string
	// 生成的dao文件目录位置
	xdao        string
	basedaoPath string
	// 需要生成dao层代码的类型名称
	xtype string
	// 包名
	xpackageName string
	//
	xprefix string

	xsuffix string
}

var (
	Xbootargs = NewBootArgs()
)

func NewBootArgs() *BootArgs {
	args := BootArgs{
		xentityPath:  ".",
		xdao:         "daos/",
		basedaoPath:  "",
		xtype:        "test",
		xpackageName: "dao",
		xprefix:      "",
		xsuffix:      "",
	}
	//
	flag.StringVar(&args.xentityPath, "entity", args.xentityPath, "")
	flag.StringVar(&args.xdao, "dao", args.xdao, "")
	flag.StringVar(&args.basedaoPath, "bdp", args.basedaoPath, "base dao path")
	flag.StringVar(&args.xtype, "type", args.xtype, "")
	flag.StringVar(&args.xpackageName, "pkg", args.xpackageName, "")
	flag.StringVar(&args.xprefix, "prefix", args.xprefix, "")
	flag.StringVar(&args.xsuffix, "suffix", args.xsuffix, "")
	flag.StringVar(&args.xprojectName, "project", args.xprojectName, "project 名称")
	flag.Parse()
	if len(args.basedaoPath) == 0 {
		args.basedaoPath = filepath.Join(args.xdao, "basedao")
	}
	//
	return &args
}

func main() {
	CreateBaseDao()
	CreateDao()
}

func CreateBaseDao() {
	model := CreateModel(Xbootargs.xtype)
	model.Fileinfo.Filename = fmt.Sprintf("%s%s", Xbootargs.xtype, "Dao")
	view, err := CreateTempl("basedao.tmpl")
	if err != nil {
		log.Fatalf("%v \n ", err)
	}
	mv := ModelAndView{
		data: *model,
		tmpl: view,
	}

	var fwriter io.Writer
	f := filepath.Join(Xbootargs.basedaoPath, fmt.Sprintf("%s.go", model.Fileinfo.Filename))
	os.Remove(f)
	fwriter = GetWriter(f)
	// fwriter=os.Stdout
	err = mv.Merge(fwriter)
	if err != nil {
		log.Fatalf("渲染模板失败: %v \n ", err)
	}
}

func CreateDao() {
	model := CreateModel(Xbootargs.xtype)
	model.Fileinfo.Filename = fmt.Sprintf("%s%s", Xbootargs.xtype, "Dao")
	view, err := CreateTempl("dao.tmpl")
	if err != nil {
		log.Fatalf("%v \n ", err)
	}
	mv := ModelAndView{
		data: *model,
		tmpl: view,
	}

	var fwriter io.Writer
	f := filepath.Join(Xbootargs.xdao, fmt.Sprintf("%s.go", model.Fileinfo.Filename))
	//
	_xf, _ := os.Stat(f)
	// 该文件已存在
	if _xf != nil {
		// 需要手动编写 dao过程的,不可以覆盖原有文件
		log.Printf("文件[%s]已存在,不覆盖.", f)
		return
	}
	//
	fwriter = GetWriter(f)
	// fwriter=os.Stdout
	err = mv.Merge(fwriter)
	if err != nil {
		log.Fatalf("渲染模板失败: %v \n ", err)
	}
}
func GetWriter(filename string) io.Writer {
	filename = filepath.Clean(filename)
	dir, _ := filepath.Split(filename)
	os.MkdirAll(dir, 0777)
	log.Printf("写入到文件:%v", filename)
	// 000 --- 0
	// 111 rwx 7
	writer, e := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0664)
	if e != nil {
		log.Fatalf("创建文件流失败: %v \n ", e)
	}
	return writer
}

// File Name: {{.Filename}}.go
// Author: {{.Author}}
// AuthorSite: {{.Domain}}
// GitSource: {{.Gitsource}}
// Created Time: {{.CreateTime}}
// ---------------------说明--------------------------
// {{.DescribeInfo}}
var fileinfo = FileInfos{
	Author:    "hanxu",
	Domain:    " ",
	Gitsource: " ",
}

type FileInfos struct {
	Filename     string
	Author       string
	Domain       string
	Gitsource    string
	CreateTime   string
	DescribeInfo string
}

type Commons struct {
	Prefix      string
	Suffix      string
	PkgName     string
	ProjectName string
	EntityPath  string
}

type TemplData struct {
	Fileinfo FileInfos
	Common   Commons
	DaoName  string
}

//
type ModelAndView struct {
	data TemplData
	tmpl *template.Template
}

func CreateTempl(name string) (tmpl *template.Template, err error) {
	templ := template.New(name)
	tmplfile := []string{name}
	for _, x := range tmplfile {
		b, err := resouces.TmplFile(x)
		if err != nil {
			return nil, err
		}
		_, err = templ.Parse(string(b))
		if err != nil {
			return nil, err
		}
	}
	return templ, nil
}

func CreateModel(typename string) *TemplData {
	var _fileinfo = fileinfo
	_fileinfo.CreateTime = time.Now().Format(internal.TimeFormatCommon)
	var _common = Commons{
		Prefix:      Xbootargs.xprefix,
		Suffix:      Xbootargs.xsuffix,
		PkgName:     Xbootargs.xpackageName,
		ProjectName: Xbootargs.xprojectName,
		EntityPath:  Xbootargs.xentityPath,
	}
	td := TemplData{
		Fileinfo: _fileinfo,
		Common:   _common,
		DaoName:  typename,
	}
	return &td
}

func (mv *ModelAndView) Merge(writer io.Writer) error {
	if mv.tmpl != nil {
		return mv.tmpl.Execute(writer, mv.data)
	}

	return errors.New("模板为空")
}
