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

import (
	"testing"
)

func TestCreateDao(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "f"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateDao()
		})
	}
}
