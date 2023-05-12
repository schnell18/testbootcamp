// Code generated by MockGen. DO NOT EDIT.
// Source: score.go

// Package persistence is a generated GoMock package.
package persistence

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockScoreDAO is a mock of ScoreDAO interface.
type MockScoreDAO struct {
	ctrl     *gomock.Controller
	recorder *MockScoreDAOMockRecorder
}

// MockScoreDAOMockRecorder is the mock recorder for MockScoreDAO.
type MockScoreDAOMockRecorder struct {
	mock *MockScoreDAO
}

// NewMockScoreDAO creates a new mock instance.
func NewMockScoreDAO(ctrl *gomock.Controller) *MockScoreDAO {
	mock := &MockScoreDAO{ctrl: ctrl}
	mock.recorder = &MockScoreDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScoreDAO) EXPECT() *MockScoreDAOMockRecorder {
	return m.recorder
}

// GetScore mocks base method.
func (m *MockScoreDAO) GetScore(student, subject string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScore", student, subject)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScore indicates an expected call of GetScore.
func (mr *MockScoreDAOMockRecorder) GetScore(student, subject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScore", reflect.TypeOf((*MockScoreDAO)(nil).GetScore), student, subject)
}
