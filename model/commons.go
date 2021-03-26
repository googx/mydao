/*
--------------------------------------------------
 File Name: commons.go
 Author: hanxu

 Created Time: 2019-8-31-上午11:18
---------------------说明--------------------------

---------------------------------------------------
*/

package model

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/googx/mydao/internal"
)

type ImodelId interface {
	IsEmpty() bool
}

// 主键id
type ModelId struct {
	ID uint `gorm:"primary_key"`
}

func (m *ModelId) IsEmpty() bool {
	return m.ID == 0
}

// http://www.axiaoxin.com/article/241/
// 该时间使用 2006-01-02 15:04:05 格式来展示
type JsonTime struct {
	time.Time
}

func NowJsonTime() JsonTime {
	return JsonTime{
		time.Now(),
	}
}

// 所以我们只需定义一个内嵌time.Time的结构体，并重写MarshalJSON方法，然后在定义model的时候把time.Time类型替换为我们自己的类型即可。但是在gorm中只重写MarshalJSON是不够的，只写这个方法会在写数据库的时候会提示delete_at字段不存在，需要加上database/sql的Value和Scan方法 https://github.com/jinzhu/gorm/issues/1611#issuecomment-329654638。
func (jtime JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", jtime.Format(internal.TimeFormatCommon))), nil
}

func (jtime *JsonTime) UnmarshalJSON(data []byte) error {
	if data != nil && len(data) != 0 {
		jtimestr := string(data)
		// var e error
		if gotime, e := time.Parse(`"`+internal.TimeFormatCommon+`"`, jtimestr); e != nil {
			return e
		} else {
			*jtime = JsonTime{Time: gotime}
			return nil
		}
	}
	return fmt.Errorf("错误的时间:%v", data)
}

func (jtime JsonTime) String() string {
	return fmt.Sprintf("%s", jtime.Format(internal.TimeFormatCommon))
}

// Value insert timestamp into mysql need this function.
func (t JsonTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *JsonTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JsonTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// TODO 需要研究下 gorm中的回调机制 https://www.cnblogs.com/sgyBlog/p/10154424.html
// 创建和修改时间
type ModelTime struct {
	CreatedAt JsonTime
	UpdatedAt JsonTime
}

// 软删除
type ModelSofeDel struct {
	DeletedAt JsonTime
}
