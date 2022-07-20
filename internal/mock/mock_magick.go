// Code generated by MockGen. DO NOT EDIT.
// Source: magick.go

// Package mock is a generated GoMock package.
package mock

import (
	log "log"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMagickClientImpl is a mock of MagickClientImpl interface.
type MockMagickClientImpl struct {
	ctrl     *gomock.Controller
	recorder *MockMagickClientImplMockRecorder
}

// MockMagickClientImplMockRecorder is the mock recorder for MockMagickClientImpl.
type MockMagickClientImplMockRecorder struct {
	mock *MockMagickClientImpl
}

// NewMockMagickClientImpl creates a new mock instance.
func NewMockMagickClientImpl(ctrl *gomock.Controller) *MockMagickClientImpl {
	mock := &MockMagickClientImpl{ctrl: ctrl}
	mock.recorder = &MockMagickClientImplMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMagickClientImpl) EXPECT() *MockMagickClientImplMockRecorder {
	return m.recorder
}

// AddDescriptionToImage mocks base method.
func (m *MockMagickClientImpl) AddDescriptionToImage(logger *log.Logger, description, imagePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDescriptionToImage", logger, description, imagePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDescriptionToImage indicates an expected call of AddDescriptionToImage.
func (mr *MockMagickClientImplMockRecorder) AddDescriptionToImage(logger, description, imagePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDescriptionToImage", reflect.TypeOf((*MockMagickClientImpl)(nil).AddDescriptionToImage), logger, description, imagePath)
}

// AddLabelToImage mocks base method.
func (m *MockMagickClientImpl) AddLabelToImage(logger *log.Logger, label, imagePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLabelToImage", logger, label, imagePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddLabelToImage indicates an expected call of AddLabelToImage.
func (mr *MockMagickClientImplMockRecorder) AddLabelToImage(logger, label, imagePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLabelToImage", reflect.TypeOf((*MockMagickClientImpl)(nil).AddLabelToImage), logger, label, imagePath)
}

// CaptureWindow mocks base method.
func (m *MockMagickClientImpl) CaptureWindow(logger *log.Logger, buildDirectory, screenshotName, screenshotExtension string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CaptureWindow", logger, buildDirectory, screenshotName, screenshotExtension)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CaptureWindow indicates an expected call of CaptureWindow.
func (mr *MockMagickClientImplMockRecorder) CaptureWindow(logger, buildDirectory, screenshotName, screenshotExtension interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CaptureWindow", reflect.TypeOf((*MockMagickClientImpl)(nil).CaptureWindow), logger, buildDirectory, screenshotName, screenshotExtension)
}
