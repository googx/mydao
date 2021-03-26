/*
--------------------------------------------------
 File Name: baseErr.go
 Author: hanxu
 AuthorSite: http://www.googx.top/
 GitSource: https://github.com/googx/linuxShell
 Created Time: 2020-5-9-上午10:01
---------------------说明--------------------------

---------------------------------------------------
*/

package base

type TransactionError struct {
	DbErr     error
	HandleErr error
}

func (te *TransactionError) Error() string {
	panic("implement me")
}
