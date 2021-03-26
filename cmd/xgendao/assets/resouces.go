/*
--------------------------------------------------
 File Name: generate.go
 Author: hanxu

 Created Time: 2019-9-10-下午1:48
---------------------说明--------------------------

---------------------------------------------------
*/

package resouces

import (
	"path/filepath"

	assetsFile "github.com/googx/mydao/cmd/xgendao/assets/assetsfile"
)

//go:generate go-bindata -pkg assetsFile -o assetsfile/tmplFile.go resources/tmplFile/

// 取得模板文件内容
func TmplFile(tmplfileName string) ([]byte, error) {
	f := filepath.Join("resources/tmplFile/", tmplfileName)
	return assetsFile.Asset(f)
}
