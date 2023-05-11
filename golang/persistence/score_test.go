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
	m.EXPECT().GetScore(gomock.Eq("Tom"), gomock.Eq("Math")).Return(90, nil)
	m.EXPECT().GetScore(gomock.Eq("Tom"), gomock.Eq("Chinese")).Return(81, nil)
	m.EXPECT().GetScore(gomock.Eq("Tom"), gomock.Eq("English")).Return(90, nil)

	avg, err := GetAverageScore(m, "Tom")
	assert.Nil(err)
	assert.Equal(87, avg)
}
