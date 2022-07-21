// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock is a generated GoMock package.
package mock

import (
	log "log"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockImageProcessingImpl is a mock of ImageProcessingImpl interface.
type MockImageProcessingImpl struct {
	ctrl     *gomock.Controller
	recorder *MockImageProcessingImplMockRecorder
}

// MockImageProcessingImplMockRecorder is the mock recorder for MockImageProcessingImpl.
type MockImageProcessingImplMockRecorder struct {
	mock *MockImageProcessingImpl
}

// NewMockImageProcessingImpl creates a new mock instance.
func NewMockImageProcessingImpl(ctrl *gomock.Controller) *MockImageProcessingImpl {
	mock := &MockImageProcessingImpl{ctrl: ctrl}
	mock.recorder = &MockImageProcessingImplMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageProcessingImpl) EXPECT() *MockImageProcessingImplMockRecorder {
	return m.recorder
}

// AddDescriptionToImage mocks base method.
func (m *MockImageProcessingImpl) AddDescriptionToImage(logger *log.Logger, description, imagePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDescriptionToImage", logger, description, imagePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDescriptionToImage indicates an expected call of AddDescriptionToImage.
func (mr *MockImageProcessingImplMockRecorder) AddDescriptionToImage(logger, description, imagePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDescriptionToImage", reflect.TypeOf((*MockImageProcessingImpl)(nil).AddDescriptionToImage), logger, description, imagePath)
}

// AddLabelToImage mocks base method.
func (m *MockImageProcessingImpl) AddLabelToImage(logger *log.Logger, label, imagePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLabelToImage", logger, label, imagePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddLabelToImage indicates an expected call of AddLabelToImage.
func (mr *MockImageProcessingImplMockRecorder) AddLabelToImage(logger, label, imagePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLabelToImage", reflect.TypeOf((*MockImageProcessingImpl)(nil).AddLabelToImage), logger, label, imagePath)
}

// CaptureWindow mocks base method.
func (m *MockImageProcessingImpl) CaptureWindow(logger *log.Logger, buildDirectory, screenshotName, screenshotExtension string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CaptureWindow", logger, buildDirectory, screenshotName, screenshotExtension)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CaptureWindow indicates an expected call of CaptureWindow.
func (mr *MockImageProcessingImplMockRecorder) CaptureWindow(logger, buildDirectory, screenshotName, screenshotExtension interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CaptureWindow", reflect.TypeOf((*MockImageProcessingImpl)(nil).CaptureWindow), logger, buildDirectory, screenshotName, screenshotExtension)
}
