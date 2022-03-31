/*
Author:ydy
Date:
Desc:
*/
package testx

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestAdd_UseGoMock(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))

	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}
}
