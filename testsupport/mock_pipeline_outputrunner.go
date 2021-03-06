// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/mozilla-services/heka/pipeline (interfaces: OutputRunner)

package testsupport

import (
	sync "sync"
	pipeline "github.com/mozilla-services/heka/pipeline"
	gomock "code.google.com/p/gomock/gomock"
	time "time"
)

// Mock of OutputRunner interface
type MockOutputRunner struct {
	ctrl     *gomock.Controller
	recorder *_MockOutputRunnerRecorder
}

// Recorder for MockOutputRunner (not exported)
type _MockOutputRunnerRecorder struct {
	mock *MockOutputRunner
}

func NewMockOutputRunner(ctrl *gomock.Controller) *MockOutputRunner {
	mock := &MockOutputRunner{ctrl: ctrl}
	mock.recorder = &_MockOutputRunnerRecorder{mock}
	return mock
}

func (_m *MockOutputRunner) EXPECT() *_MockOutputRunnerRecorder {
	return _m.recorder
}

func (_m *MockOutputRunner) Deliver(_param0 *pipeline.PipelinePack) {
	_m.ctrl.Call(_m, "Deliver", _param0)
}

func (_mr *_MockOutputRunnerRecorder) Deliver(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Deliver", arg0)
}

func (_m *MockOutputRunner) InChan() chan *pipeline.PipelineCapture {
	ret := _m.ctrl.Call(_m, "InChan")
	ret0, _ := ret[0].(chan *pipeline.PipelineCapture)
	return ret0
}

func (_mr *_MockOutputRunnerRecorder) InChan() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InChan")
}

func (_m *MockOutputRunner) LogError(_param0 error) {
	_m.ctrl.Call(_m, "LogError", _param0)
}

func (_mr *_MockOutputRunnerRecorder) LogError(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LogError", arg0)
}

func (_m *MockOutputRunner) LogMessage(_param0 string) {
	_m.ctrl.Call(_m, "LogMessage", _param0)
}

func (_mr *_MockOutputRunnerRecorder) LogMessage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LogMessage", arg0)
}

func (_m *MockOutputRunner) Name() string {
	ret := _m.ctrl.Call(_m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockOutputRunnerRecorder) Name() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Name")
}

func (_m *MockOutputRunner) Output() pipeline.Output {
	ret := _m.ctrl.Call(_m, "Output")
	ret0, _ := ret[0].(pipeline.Output)
	return ret0
}

func (_mr *_MockOutputRunnerRecorder) Output() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Output")
}

func (_m *MockOutputRunner) Plugin() pipeline.Plugin {
	ret := _m.ctrl.Call(_m, "Plugin")
	ret0, _ := ret[0].(pipeline.Plugin)
	return ret0
}

func (_mr *_MockOutputRunnerRecorder) Plugin() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Plugin")
}

func (_m *MockOutputRunner) SetName(_param0 string) {
	_m.ctrl.Call(_m, "SetName", _param0)
}

func (_mr *_MockOutputRunnerRecorder) SetName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetName", arg0)
}

func (_m *MockOutputRunner) Start(_param0 pipeline.PluginHelper, _param1 *sync.WaitGroup) error {
	ret := _m.ctrl.Call(_m, "Start", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockOutputRunnerRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start", arg0, arg1)
}

func (_m *MockOutputRunner) Ticker() <-chan time.Time {
	ret := _m.ctrl.Call(_m, "Ticker")
	ret0, _ := ret[0].(<-chan time.Time)
	return ret0
}

func (_mr *_MockOutputRunnerRecorder) Ticker() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Ticker")
}
