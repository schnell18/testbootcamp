package persistence

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetFromDB(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDB(ctrl)
	ivc1 := m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))
	ivc2 := m.EXPECT().Get(gomock.Eq("Sam")).Return(1, nil).Times(1)
	ivc3 := m.EXPECT().Get(gomock.Eq("Jake")).Return(0, nil).Times(2)
	gomock.InOrder(ivc2, ivc3, ivc1)
	// gomock.InOrder(ivc1, ivc3, ivc2)

	assert.Equal(1, GetFromDB(m, "Sam"))
	assert.Equal(0, GetFromDB(m, "Jake"))
	assert.Equal(1, GetFromDB(m, "Jake"))
	assert.Equal(-2, GetFromDB(m, "Tom"))
}
