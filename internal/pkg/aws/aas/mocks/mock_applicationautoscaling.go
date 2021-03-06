// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/aws/applicationautoscaling/applicationautoscaling.go

// Package mocks is a generated GoMock package.
package mocks

import (
	applicationautoscaling "github.com/aws/aws-sdk-go/service/applicationautoscaling"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Mockapi is a mock of api interface
type Mockapi struct {
	ctrl     *gomock.Controller
	recorder *MockapiMockRecorder
}

// MockapiMockRecorder is the mock recorder for Mockapi
type MockapiMockRecorder struct {
	mock *Mockapi
}

// NewMockapi creates a new mock instance
func NewMockapi(ctrl *gomock.Controller) *Mockapi {
	mock := &Mockapi{ctrl: ctrl}
	mock.recorder = &MockapiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockapi) EXPECT() *MockapiMockRecorder {
	return m.recorder
}

// DescribeScalingPolicies mocks base method
func (m *Mockapi) DescribeScalingPolicies(input *applicationautoscaling.DescribeScalingPoliciesInput) (*applicationautoscaling.DescribeScalingPoliciesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeScalingPolicies", input)
	ret0, _ := ret[0].(*applicationautoscaling.DescribeScalingPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeScalingPolicies indicates an expected call of DescribeScalingPolicies
func (mr *MockapiMockRecorder) DescribeScalingPolicies(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScalingPolicies", reflect.TypeOf((*Mockapi)(nil).DescribeScalingPolicies), input)
}
