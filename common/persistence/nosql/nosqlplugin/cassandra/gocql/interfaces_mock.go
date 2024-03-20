// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package gocql is a generated GoMock package.
package gocql

import (
	context "context"
	reflect "reflect"

	gocql "github.com/gocql/gocql"
	gomock "github.com/golang/mock/gomock"
)

// MockSession is a mock of Session interface.
type MockSession struct {
	ctrl     *gomock.Controller
	recorder *MockSessionMockRecorder
}

// MockSessionMockRecorder is the mock recorder for MockSession.
type MockSessionMockRecorder struct {
	mock *MockSession
}

// NewMockSession creates a new mock instance.
func NewMockSession(ctrl *gomock.Controller) *MockSession {
	mock := &MockSession{ctrl: ctrl}
	mock.recorder = &MockSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSession) EXPECT() *MockSessionMockRecorder {
	return m.recorder
}

// AwaitSchemaAgreement mocks base method.
func (m *MockSession) AwaitSchemaAgreement(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AwaitSchemaAgreement", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// AwaitSchemaAgreement indicates an expected call of AwaitSchemaAgreement.
func (mr *MockSessionMockRecorder) AwaitSchemaAgreement(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AwaitSchemaAgreement", reflect.TypeOf((*MockSession)(nil).AwaitSchemaAgreement), ctx)
}

// Close mocks base method.
func (m *MockSession) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockSessionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSession)(nil).Close))
}

// ExecuteBatch mocks base method.
func (m *MockSession) ExecuteBatch(arg0 Batch) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteBatch", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteBatch indicates an expected call of ExecuteBatch.
func (mr *MockSessionMockRecorder) ExecuteBatch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteBatch", reflect.TypeOf((*MockSession)(nil).ExecuteBatch), arg0)
}

// MapExecuteBatchCAS mocks base method.
func (m *MockSession) MapExecuteBatchCAS(arg0 Batch, arg1 map[string]interface{}) (bool, Iter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapExecuteBatchCAS", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(Iter)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// MapExecuteBatchCAS indicates an expected call of MapExecuteBatchCAS.
func (mr *MockSessionMockRecorder) MapExecuteBatchCAS(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapExecuteBatchCAS", reflect.TypeOf((*MockSession)(nil).MapExecuteBatchCAS), arg0, arg1)
}

// NewBatch mocks base method.
func (m *MockSession) NewBatch(arg0 BatchType) Batch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewBatch", arg0)
	ret0, _ := ret[0].(Batch)
	return ret0
}

// NewBatch indicates an expected call of NewBatch.
func (mr *MockSessionMockRecorder) NewBatch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewBatch", reflect.TypeOf((*MockSession)(nil).NewBatch), arg0)
}

// Query mocks base method.
func (m *MockSession) Query(arg0 string, arg1 ...interface{}) Query {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(Query)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockSessionMockRecorder) Query(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockSession)(nil).Query), varargs...)
}

// MockGocqlSession is a mock of GocqlSession interface.
type MockGocqlSession struct {
	ctrl     *gomock.Controller
	recorder *MockGocqlSessionMockRecorder
}

// MockGocqlSessionMockRecorder is the mock recorder for MockGocqlSession.
type MockGocqlSessionMockRecorder struct {
	mock *MockGocqlSession
}

// NewMockGocqlSession creates a new mock instance.
func NewMockGocqlSession(ctrl *gomock.Controller) *MockGocqlSession {
	mock := &MockGocqlSession{ctrl: ctrl}
	mock.recorder = &MockGocqlSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGocqlSession) EXPECT() *MockGocqlSessionMockRecorder {
	return m.recorder
}

// AwaitSchemaAgreement mocks base method.
func (m *MockGocqlSession) AwaitSchemaAgreement(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AwaitSchemaAgreement", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// AwaitSchemaAgreement indicates an expected call of AwaitSchemaAgreement.
func (mr *MockGocqlSessionMockRecorder) AwaitSchemaAgreement(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AwaitSchemaAgreement", reflect.TypeOf((*MockGocqlSession)(nil).AwaitSchemaAgreement), ctx)
}

// Close mocks base method.
func (m *MockGocqlSession) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockGocqlSessionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockGocqlSession)(nil).Close))
}

// ExecuteBatch mocks base method.
func (m *MockGocqlSession) ExecuteBatch(batch *gocql.Batch) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteBatch", batch)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteBatch indicates an expected call of ExecuteBatch.
func (mr *MockGocqlSessionMockRecorder) ExecuteBatch(batch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteBatch", reflect.TypeOf((*MockGocqlSession)(nil).ExecuteBatch), batch)
}

// MapExecuteBatchCAS mocks base method.
func (m *MockGocqlSession) MapExecuteBatchCAS(batch *gocql.Batch, dest map[string]interface{}) (bool, *gocql.Iter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapExecuteBatchCAS", batch, dest)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*gocql.Iter)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// MapExecuteBatchCAS indicates an expected call of MapExecuteBatchCAS.
func (mr *MockGocqlSessionMockRecorder) MapExecuteBatchCAS(batch, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapExecuteBatchCAS", reflect.TypeOf((*MockGocqlSession)(nil).MapExecuteBatchCAS), batch, dest)
}

// NewBatch mocks base method.
func (m *MockGocqlSession) NewBatch(typ gocql.BatchType) *gocql.Batch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewBatch", typ)
	ret0, _ := ret[0].(*gocql.Batch)
	return ret0
}

// NewBatch indicates an expected call of NewBatch.
func (mr *MockGocqlSessionMockRecorder) NewBatch(typ interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewBatch", reflect.TypeOf((*MockGocqlSession)(nil).NewBatch), typ)
}

// Query mocks base method.
func (m *MockGocqlSession) Query(stmt string, values ...interface{}) *gocql.Query {
	m.ctrl.T.Helper()
	varargs := []interface{}{stmt}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(*gocql.Query)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockGocqlSessionMockRecorder) Query(stmt interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{stmt}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockGocqlSession)(nil).Query), varargs...)
}

// MockQuery is a mock of Query interface.
type MockQuery struct {
	ctrl     *gomock.Controller
	recorder *MockQueryMockRecorder
}

// MockQueryMockRecorder is the mock recorder for MockQuery.
type MockQueryMockRecorder struct {
	mock *MockQuery
}

// NewMockQuery creates a new mock instance.
func NewMockQuery(ctrl *gomock.Controller) *MockQuery {
	mock := &MockQuery{ctrl: ctrl}
	mock.recorder = &MockQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuery) EXPECT() *MockQueryMockRecorder {
	return m.recorder
}

// Bind mocks base method.
func (m *MockQuery) Bind(arg0 ...interface{}) Query {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Bind", varargs...)
	ret0, _ := ret[0].(Query)
	return ret0
}

// Bind indicates an expected call of Bind.
func (mr *MockQueryMockRecorder) Bind(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bind", reflect.TypeOf((*MockQuery)(nil).Bind), arg0...)
}

// Consistency mocks base method.
func (m *MockQuery) Consistency(arg0 Consistency) Query {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consistency", arg0)
	ret0, _ := ret[0].(Query)
	return ret0
}

// Consistency indicates an expected call of Consistency.
func (mr *MockQueryMockRecorder) Consistency(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consistency", reflect.TypeOf((*MockQuery)(nil).Consistency), arg0)
}

// Exec mocks base method.
func (m *MockQuery) Exec() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exec")
	ret0, _ := ret[0].(error)
	return ret0
}

// Exec indicates an expected call of Exec.
func (mr *MockQueryMockRecorder) Exec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockQuery)(nil).Exec))
}

// Iter mocks base method.
func (m *MockQuery) Iter() Iter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iter")
	ret0, _ := ret[0].(Iter)
	return ret0
}

// Iter indicates an expected call of Iter.
func (mr *MockQueryMockRecorder) Iter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iter", reflect.TypeOf((*MockQuery)(nil).Iter))
}

// MapScan mocks base method.
func (m *MockQuery) MapScan(arg0 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapScan", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MapScan indicates an expected call of MapScan.
func (mr *MockQueryMockRecorder) MapScan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapScan", reflect.TypeOf((*MockQuery)(nil).MapScan), arg0)
}

// MapScanCAS mocks base method.
func (m *MockQuery) MapScanCAS(arg0 map[string]interface{}) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapScanCAS", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MapScanCAS indicates an expected call of MapScanCAS.
func (mr *MockQueryMockRecorder) MapScanCAS(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapScanCAS", reflect.TypeOf((*MockQuery)(nil).MapScanCAS), arg0)
}

// PageSize mocks base method.
func (m *MockQuery) PageSize(arg0 int) Query {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PageSize", arg0)
	ret0, _ := ret[0].(Query)
	return ret0
}

// PageSize indicates an expected call of PageSize.
func (mr *MockQueryMockRecorder) PageSize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PageSize", reflect.TypeOf((*MockQuery)(nil).PageSize), arg0)
}

// PageState mocks base method.
func (m *MockQuery) PageState(arg0 []byte) Query {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PageState", arg0)
	ret0, _ := ret[0].(Query)
	return ret0
}

// PageState indicates an expected call of PageState.
func (mr *MockQueryMockRecorder) PageState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PageState", reflect.TypeOf((*MockQuery)(nil).PageState), arg0)
}

// Scan mocks base method.
func (m *MockQuery) Scan(arg0 ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scan", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockQueryMockRecorder) Scan(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockQuery)(nil).Scan), arg0...)
}

// ScanCAS mocks base method.
func (m *MockQuery) ScanCAS(arg0 ...interface{}) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScanCAS", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScanCAS indicates an expected call of ScanCAS.
func (mr *MockQueryMockRecorder) ScanCAS(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanCAS", reflect.TypeOf((*MockQuery)(nil).ScanCAS), arg0...)
}

// WithContext mocks base method.
func (m *MockQuery) WithContext(arg0 context.Context) Query {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithContext", arg0)
	ret0, _ := ret[0].(Query)
	return ret0
}

// WithContext indicates an expected call of WithContext.
func (mr *MockQueryMockRecorder) WithContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithContext", reflect.TypeOf((*MockQuery)(nil).WithContext), arg0)
}

// WithTimestamp mocks base method.
func (m *MockQuery) WithTimestamp(arg0 int64) Query {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTimestamp", arg0)
	ret0, _ := ret[0].(Query)
	return ret0
}

// WithTimestamp indicates an expected call of WithTimestamp.
func (mr *MockQueryMockRecorder) WithTimestamp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTimestamp", reflect.TypeOf((*MockQuery)(nil).WithTimestamp), arg0)
}

// MockBatch is a mock of Batch interface.
type MockBatch struct {
	ctrl     *gomock.Controller
	recorder *MockBatchMockRecorder
}

// MockBatchMockRecorder is the mock recorder for MockBatch.
type MockBatchMockRecorder struct {
	mock *MockBatch
}

// NewMockBatch creates a new mock instance.
func NewMockBatch(ctrl *gomock.Controller) *MockBatch {
	mock := &MockBatch{ctrl: ctrl}
	mock.recorder = &MockBatchMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBatch) EXPECT() *MockBatchMockRecorder {
	return m.recorder
}

// Query mocks base method.
func (m *MockBatch) Query(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Query", varargs...)
}

// Query indicates an expected call of Query.
func (mr *MockBatchMockRecorder) Query(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockBatch)(nil).Query), varargs...)
}

// WithContext mocks base method.
func (m *MockBatch) WithContext(arg0 context.Context) Batch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithContext", arg0)
	ret0, _ := ret[0].(Batch)
	return ret0
}

// WithContext indicates an expected call of WithContext.
func (mr *MockBatchMockRecorder) WithContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithContext", reflect.TypeOf((*MockBatch)(nil).WithContext), arg0)
}

// WithTimestamp mocks base method.
func (m *MockBatch) WithTimestamp(arg0 int64) Batch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTimestamp", arg0)
	ret0, _ := ret[0].(Batch)
	return ret0
}

// WithTimestamp indicates an expected call of WithTimestamp.
func (mr *MockBatchMockRecorder) WithTimestamp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTimestamp", reflect.TypeOf((*MockBatch)(nil).WithTimestamp), arg0)
}

// MockIter is a mock of Iter interface.
type MockIter struct {
	ctrl     *gomock.Controller
	recorder *MockIterMockRecorder
}

// MockIterMockRecorder is the mock recorder for MockIter.
type MockIterMockRecorder struct {
	mock *MockIter
}

// NewMockIter creates a new mock instance.
func NewMockIter(ctrl *gomock.Controller) *MockIter {
	mock := &MockIter{ctrl: ctrl}
	mock.recorder = &MockIterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIter) EXPECT() *MockIterMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIter)(nil).Close))
}

// MapScan mocks base method.
func (m *MockIter) MapScan(arg0 map[string]interface{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapScan", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// MapScan indicates an expected call of MapScan.
func (mr *MockIterMockRecorder) MapScan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapScan", reflect.TypeOf((*MockIter)(nil).MapScan), arg0)
}

// PageState mocks base method.
func (m *MockIter) PageState() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PageState")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// PageState indicates an expected call of PageState.
func (mr *MockIterMockRecorder) PageState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PageState", reflect.TypeOf((*MockIter)(nil).PageState))
}

// Scan mocks base method.
func (m *MockIter) Scan(arg0 ...interface{}) bool {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scan", varargs...)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockIterMockRecorder) Scan(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockIter)(nil).Scan), arg0...)
}
