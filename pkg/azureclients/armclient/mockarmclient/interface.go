// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/azureclients/armclient/interface.go

// Package mockarmclient is a generated GoMock package.
package mockarmclient

import (
	context "context"
	http "net/http"
	reflect "reflect"

	autorest "github.com/Azure/go-autorest/autorest"
	azure "github.com/Azure/go-autorest/autorest/azure"
	gomock "github.com/golang/mock/gomock"
	armclient "sigs.k8s.io/cloud-provider-azure/pkg/azureclients/armclient"
	retry "sigs.k8s.io/cloud-provider-azure/pkg/retry"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// CloseResponse mocks base method.
func (m *MockInterface) CloseResponse(ctx context.Context, response *http.Response) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloseResponse", ctx, response)
}

// CloseResponse indicates an expected call of CloseResponse.
func (mr *MockInterfaceMockRecorder) CloseResponse(ctx, response interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseResponse", reflect.TypeOf((*MockInterface)(nil).CloseResponse), ctx, response)
}

// DeleteResource mocks base method.
func (m *MockInterface) DeleteResource(ctx context.Context, resourceID string, decorators ...autorest.PrepareDecorator) *retry.Error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, resourceID}
	for _, a := range decorators {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteResource", varargs...)
	ret0, _ := ret[0].(*retry.Error)
	return ret0
}

// DeleteResource indicates an expected call of DeleteResource.
func (mr *MockInterfaceMockRecorder) DeleteResource(ctx, resourceID interface{}, decorators ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, resourceID}, decorators...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResource", reflect.TypeOf((*MockInterface)(nil).DeleteResource), varargs...)
}

// DeleteResourceAsync mocks base method.
func (m *MockInterface) DeleteResourceAsync(ctx context.Context, resourceID string, decorators ...autorest.PrepareDecorator) (*azure.Future, *retry.Error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, resourceID}
	for _, a := range decorators {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteResourceAsync", varargs...)
	ret0, _ := ret[0].(*azure.Future)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// DeleteResourceAsync indicates an expected call of DeleteResourceAsync.
func (mr *MockInterfaceMockRecorder) DeleteResourceAsync(ctx, resourceID interface{}, decorators ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, resourceID}, decorators...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResourceAsync", reflect.TypeOf((*MockInterface)(nil).DeleteResourceAsync), varargs...)
}

// GetResource mocks base method.
func (m *MockInterface) GetResource(ctx context.Context, resourceID string, decorators ...autorest.PrepareDecorator) (*http.Response, *retry.Error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, resourceID}
	for _, a := range decorators {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetResource", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// GetResource indicates an expected call of GetResource.
func (mr *MockInterfaceMockRecorder) GetResource(ctx, resourceID interface{}, decorators ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, resourceID}, decorators...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResource", reflect.TypeOf((*MockInterface)(nil).GetResource), varargs...)
}

// GetResourceWithExpandAPIVersionQuery mocks base method.
func (m *MockInterface) GetResourceWithExpandAPIVersionQuery(ctx context.Context, resourceID, expand, apiVersion string) (*http.Response, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceWithExpandAPIVersionQuery", ctx, resourceID, expand, apiVersion)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// GetResourceWithExpandAPIVersionQuery indicates an expected call of GetResourceWithExpandAPIVersionQuery.
func (mr *MockInterfaceMockRecorder) GetResourceWithExpandAPIVersionQuery(ctx, resourceID, expand, apiVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceWithExpandAPIVersionQuery", reflect.TypeOf((*MockInterface)(nil).GetResourceWithExpandAPIVersionQuery), ctx, resourceID, expand, apiVersion)
}

// GetResourceWithExpandQuery mocks base method.
func (m *MockInterface) GetResourceWithExpandQuery(ctx context.Context, resourceID, expand string) (*http.Response, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceWithExpandQuery", ctx, resourceID, expand)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// GetResourceWithExpandQuery indicates an expected call of GetResourceWithExpandQuery.
func (mr *MockInterfaceMockRecorder) GetResourceWithExpandQuery(ctx, resourceID, expand interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceWithExpandQuery", reflect.TypeOf((*MockInterface)(nil).GetResourceWithExpandQuery), ctx, resourceID, expand)
}

// HeadResource mocks base method.
func (m *MockInterface) HeadResource(ctx context.Context, resourceID string) (*http.Response, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeadResource", ctx, resourceID)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// HeadResource indicates an expected call of HeadResource.
func (mr *MockInterfaceMockRecorder) HeadResource(ctx, resourceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadResource", reflect.TypeOf((*MockInterface)(nil).HeadResource), ctx, resourceID)
}

// PatchResource mocks base method.
func (m *MockInterface) PatchResource(ctx context.Context, resourceID string, parameters interface{}, decorators ...autorest.PrepareDecorator) (*http.Response, *retry.Error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, resourceID, parameters}
	for _, a := range decorators {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchResource", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// PatchResource indicates an expected call of PatchResource.
func (mr *MockInterfaceMockRecorder) PatchResource(ctx, resourceID, parameters interface{}, decorators ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, resourceID, parameters}, decorators...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchResource", reflect.TypeOf((*MockInterface)(nil).PatchResource), varargs...)
}

// PatchResourceAsync mocks base method.
func (m *MockInterface) PatchResourceAsync(ctx context.Context, resourceID string, parameters interface{}, decorators ...autorest.PrepareDecorator) (*azure.Future, *retry.Error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, resourceID, parameters}
	for _, a := range decorators {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchResourceAsync", varargs...)
	ret0, _ := ret[0].(*azure.Future)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// PatchResourceAsync indicates an expected call of PatchResourceAsync.
func (mr *MockInterfaceMockRecorder) PatchResourceAsync(ctx, resourceID, parameters interface{}, decorators ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, resourceID, parameters}, decorators...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchResourceAsync", reflect.TypeOf((*MockInterface)(nil).PatchResourceAsync), varargs...)
}

// PostResource mocks base method.
func (m *MockInterface) PostResource(ctx context.Context, resourceID, action string, parameters interface{}, queryParameters map[string]interface{}) (*http.Response, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostResource", ctx, resourceID, action, parameters, queryParameters)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// PostResource indicates an expected call of PostResource.
func (mr *MockInterfaceMockRecorder) PostResource(ctx, resourceID, action, parameters, queryParameters interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostResource", reflect.TypeOf((*MockInterface)(nil).PostResource), ctx, resourceID, action, parameters, queryParameters)
}

// PrepareDeleteRequest mocks base method.
func (m *MockInterface) PrepareDeleteRequest(ctx context.Context, decorators ...autorest.PrepareDecorator) (*http.Request, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range decorators {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PrepareDeleteRequest", varargs...)
	ret0, _ := ret[0].(*http.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareDeleteRequest indicates an expected call of PrepareDeleteRequest.
func (mr *MockInterfaceMockRecorder) PrepareDeleteRequest(ctx interface{}, decorators ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, decorators...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareDeleteRequest", reflect.TypeOf((*MockInterface)(nil).PrepareDeleteRequest), varargs...)
}

// PrepareGetRequest mocks base method.
func (m *MockInterface) PrepareGetRequest(ctx context.Context, decorators ...autorest.PrepareDecorator) (*http.Request, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range decorators {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PrepareGetRequest", varargs...)
	ret0, _ := ret[0].(*http.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareGetRequest indicates an expected call of PrepareGetRequest.
func (mr *MockInterfaceMockRecorder) PrepareGetRequest(ctx interface{}, decorators ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, decorators...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareGetRequest", reflect.TypeOf((*MockInterface)(nil).PrepareGetRequest), varargs...)
}

// PrepareHeadRequest mocks base method.
func (m *MockInterface) PrepareHeadRequest(ctx context.Context, decorators ...autorest.PrepareDecorator) (*http.Request, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range decorators {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PrepareHeadRequest", varargs...)
	ret0, _ := ret[0].(*http.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareHeadRequest indicates an expected call of PrepareHeadRequest.
func (mr *MockInterfaceMockRecorder) PrepareHeadRequest(ctx interface{}, decorators ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, decorators...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareHeadRequest", reflect.TypeOf((*MockInterface)(nil).PrepareHeadRequest), varargs...)
}

// PreparePostRequest mocks base method.
func (m *MockInterface) PreparePostRequest(ctx context.Context, decorators ...autorest.PrepareDecorator) (*http.Request, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range decorators {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PreparePostRequest", varargs...)
	ret0, _ := ret[0].(*http.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PreparePostRequest indicates an expected call of PreparePostRequest.
func (mr *MockInterfaceMockRecorder) PreparePostRequest(ctx interface{}, decorators ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, decorators...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreparePostRequest", reflect.TypeOf((*MockInterface)(nil).PreparePostRequest), varargs...)
}

// PreparePutRequest mocks base method.
func (m *MockInterface) PreparePutRequest(ctx context.Context, decorators ...autorest.PrepareDecorator) (*http.Request, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range decorators {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PreparePutRequest", varargs...)
	ret0, _ := ret[0].(*http.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PreparePutRequest indicates an expected call of PreparePutRequest.
func (mr *MockInterfaceMockRecorder) PreparePutRequest(ctx interface{}, decorators ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, decorators...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreparePutRequest", reflect.TypeOf((*MockInterface)(nil).PreparePutRequest), varargs...)
}

// PutResource mocks base method.
func (m *MockInterface) PutResource(ctx context.Context, resourceID string, parameters interface{}, decorators ...autorest.PrepareDecorator) (*http.Response, *retry.Error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, resourceID, parameters}
	for _, a := range decorators {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutResource", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// PutResource indicates an expected call of PutResource.
func (mr *MockInterfaceMockRecorder) PutResource(ctx, resourceID, parameters interface{}, decorators ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, resourceID, parameters}, decorators...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutResource", reflect.TypeOf((*MockInterface)(nil).PutResource), varargs...)
}

// PutResourceAsync mocks base method.
func (m *MockInterface) PutResourceAsync(ctx context.Context, resourceID string, parameters interface{}, decorators ...autorest.PrepareDecorator) (*azure.Future, *retry.Error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, resourceID, parameters}
	for _, a := range decorators {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutResourceAsync", varargs...)
	ret0, _ := ret[0].(*azure.Future)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// PutResourceAsync indicates an expected call of PutResourceAsync.
func (mr *MockInterfaceMockRecorder) PutResourceAsync(ctx, resourceID, parameters interface{}, decorators ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, resourceID, parameters}, decorators...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutResourceAsync", reflect.TypeOf((*MockInterface)(nil).PutResourceAsync), varargs...)
}

// PutResourcesInBatches mocks base method.
func (m *MockInterface) PutResourcesInBatches(ctx context.Context, resources map[string]interface{}, batchSize int) map[string]*armclient.PutResourcesResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutResourcesInBatches", ctx, resources, batchSize)
	ret0, _ := ret[0].(map[string]*armclient.PutResourcesResponse)
	return ret0
}

// PutResourcesInBatches indicates an expected call of PutResourcesInBatches.
func (mr *MockInterfaceMockRecorder) PutResourcesInBatches(ctx, resources, batchSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutResourcesInBatches", reflect.TypeOf((*MockInterface)(nil).PutResourcesInBatches), ctx, resources, batchSize)
}

// Send mocks base method.
func (m *MockInterface) Send(ctx context.Context, request *http.Request, decorators ...autorest.SendDecorator) (*http.Response, *retry.Error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, request}
	for _, a := range decorators {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Send", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockInterfaceMockRecorder) Send(ctx, request interface{}, decorators ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, request}, decorators...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockInterface)(nil).Send), varargs...)
}

// SendAsync mocks base method.
func (m *MockInterface) SendAsync(ctx context.Context, request *http.Request) (*azure.Future, *http.Response, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAsync", ctx, request)
	ret0, _ := ret[0].(*azure.Future)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(*retry.Error)
	return ret0, ret1, ret2
}

// SendAsync indicates an expected call of SendAsync.
func (mr *MockInterfaceMockRecorder) SendAsync(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAsync", reflect.TypeOf((*MockInterface)(nil).SendAsync), ctx, request)
}

// WaitForAsyncOperationCompletion mocks base method.
func (m *MockInterface) WaitForAsyncOperationCompletion(ctx context.Context, future *azure.Future, asyncOperationName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForAsyncOperationCompletion", ctx, future, asyncOperationName)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForAsyncOperationCompletion indicates an expected call of WaitForAsyncOperationCompletion.
func (mr *MockInterfaceMockRecorder) WaitForAsyncOperationCompletion(ctx, future, asyncOperationName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForAsyncOperationCompletion", reflect.TypeOf((*MockInterface)(nil).WaitForAsyncOperationCompletion), ctx, future, asyncOperationName)
}

// WaitForAsyncOperationResult mocks base method.
func (m *MockInterface) WaitForAsyncOperationResult(ctx context.Context, future *azure.Future, asyncOperationName string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForAsyncOperationResult", ctx, future, asyncOperationName)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForAsyncOperationResult indicates an expected call of WaitForAsyncOperationResult.
func (mr *MockInterfaceMockRecorder) WaitForAsyncOperationResult(ctx, future, asyncOperationName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForAsyncOperationResult", reflect.TypeOf((*MockInterface)(nil).WaitForAsyncOperationResult), ctx, future, asyncOperationName)
}
