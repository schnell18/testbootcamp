package persistence

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAverageScore(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockScoreDAO(ctrl)
	ivc1 := m.EXPECT().GetScore(gomock.Eq("Tom"), gomock.Eq("Matha")).Return(90, nil)
	ivc2 := m.EXPECT().GetScore(gomock.Eq("Tom"), gomock.Eq("Chinese")).Return(81, nil)
	ivc3 := m.EXPECT().GetScore(gomock.Eq("Tom"), gomock.Eq("English")).Return(90, nil)
	gomock.InOrder(ivc1, ivc2, ivc3)

	avg, err := GetAverageScore(m, "Tom")
	assert.Nil(err)
	assert.Equal(87, avg)
}
