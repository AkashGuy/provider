// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "k8s.io/api/authentication/v1alpha1"
)

// SelfSubjectReviewInterface is an autogenerated mock type for the SelfSubjectReviewInterface type
type SelfSubjectReviewInterface struct {
	mock.Mock
}

type SelfSubjectReviewInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *SelfSubjectReviewInterface) EXPECT() *SelfSubjectReviewInterface_Expecter {
	return &SelfSubjectReviewInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, selfSubjectReview, opts
func (_m *SelfSubjectReviewInterface) Create(ctx context.Context, selfSubjectReview *v1alpha1.SelfSubjectReview, opts v1.CreateOptions) (*v1alpha1.SelfSubjectReview, error) {
	ret := _m.Called(ctx, selfSubjectReview, opts)

	var r0 *v1alpha1.SelfSubjectReview
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.SelfSubjectReview, v1.CreateOptions) (*v1alpha1.SelfSubjectReview, error)); ok {
		return rf(ctx, selfSubjectReview, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.SelfSubjectReview, v1.CreateOptions) *v1alpha1.SelfSubjectReview); ok {
		r0 = rf(ctx, selfSubjectReview, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.SelfSubjectReview)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1alpha1.SelfSubjectReview, v1.CreateOptions) error); ok {
		r1 = rf(ctx, selfSubjectReview, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelfSubjectReviewInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type SelfSubjectReviewInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - selfSubjectReview *v1alpha1.SelfSubjectReview
//   - opts v1.CreateOptions
func (_e *SelfSubjectReviewInterface_Expecter) Create(ctx interface{}, selfSubjectReview interface{}, opts interface{}) *SelfSubjectReviewInterface_Create_Call {
	return &SelfSubjectReviewInterface_Create_Call{Call: _e.mock.On("Create", ctx, selfSubjectReview, opts)}
}

func (_c *SelfSubjectReviewInterface_Create_Call) Run(run func(ctx context.Context, selfSubjectReview *v1alpha1.SelfSubjectReview, opts v1.CreateOptions)) *SelfSubjectReviewInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1alpha1.SelfSubjectReview), args[2].(v1.CreateOptions))
	})
	return _c
}

func (_c *SelfSubjectReviewInterface_Create_Call) Return(_a0 *v1alpha1.SelfSubjectReview, _a1 error) *SelfSubjectReviewInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SelfSubjectReviewInterface_Create_Call) RunAndReturn(run func(context.Context, *v1alpha1.SelfSubjectReview, v1.CreateOptions) (*v1alpha1.SelfSubjectReview, error)) *SelfSubjectReviewInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewSelfSubjectReviewInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewSelfSubjectReviewInterface creates a new instance of SelfSubjectReviewInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSelfSubjectReviewInterface(t mockConstructorTestingTNewSelfSubjectReviewInterface) *SelfSubjectReviewInterface {
	mock := &SelfSubjectReviewInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}