// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import core "github.com/projectriff/riff/pkg/core"
import duckv1alpha1 "github.com/knative/pkg/apis/duck/v1alpha1"
import io "io"
import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/api/core/v1"
import v1alpha1 "github.com/knative/serving/pkg/apis/serving/v1alpha1"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// BuildFunction provides a mock function with given fields: builder, options, log
func (_m *Client) BuildFunction(builder core.Builder, options core.BuildFunctionOptions, log io.Writer) error {
	ret := _m.Called(builder, options, log)

	var r0 error
	if rf, ok := ret.Get(0).(func(core.Builder, core.BuildFunctionOptions, io.Writer) error); ok {
		r0 = rf(builder, options, log)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateFunction provides a mock function with given fields: builder, options, log
func (_m *Client) CreateFunction(builder core.Builder, options core.CreateFunctionOptions, log io.Writer) (*v1alpha1.Service, *v1alpha1.Revision, *v1.PersistentVolumeClaim, error) {
	ret := _m.Called(builder, options, log)

	var r0 *v1alpha1.Service
	if rf, ok := ret.Get(0).(func(core.Builder, core.CreateFunctionOptions, io.Writer) *v1alpha1.Service); ok {
		r0 = rf(builder, options, log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Service)
		}
	}

	var r1 *v1alpha1.Revision
	if rf, ok := ret.Get(1).(func(core.Builder, core.CreateFunctionOptions, io.Writer) *v1alpha1.Revision); ok {
		r1 = rf(builder, options, log)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*v1alpha1.Revision)
		}
	}

	var r2 *v1.PersistentVolumeClaim
	if rf, ok := ret.Get(2).(func(core.Builder, core.CreateFunctionOptions, io.Writer) *v1.PersistentVolumeClaim); ok {
		r2 = rf(builder, options, log)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*v1.PersistentVolumeClaim)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(core.Builder, core.CreateFunctionOptions, io.Writer) error); ok {
		r3 = rf(builder, options, log)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// CreateService provides a mock function with given fields: options
func (_m *Client) CreateService(options core.CreateOrUpdateServiceOptions) (*v1alpha1.Service, error) {
	ret := _m.Called(options)

	var r0 *v1alpha1.Service
	if rf, ok := ret.Get(0).(func(core.CreateOrUpdateServiceOptions) *v1alpha1.Service); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.CreateOrUpdateServiceOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DefaultBuildImagePrefix provides a mock function with given fields: namespace
func (_m *Client) DefaultBuildImagePrefix(namespace string) (string, error) {
	ret := _m.Called(namespace)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(namespace)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteService provides a mock function with given fields: options
func (_m *Client) DeleteService(options core.DeleteServiceOptions) error {
	ret := _m.Called(options)

	var r0 error
	if rf, ok := ret.Get(0).(func(core.DeleteServiceOptions) error); ok {
		r0 = rf(options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchPackConfig provides a mock function with given fields:
func (_m *Client) FetchPackConfig() (*core.PackConfig, error) {
	ret := _m.Called()

	var r0 *core.PackConfig
	if rf, ok := ret.Get(0).(func() *core.PackConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.PackConfig)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListServices provides a mock function with given fields: options
func (_m *Client) ListServices(options core.ListServiceOptions) (*v1alpha1.ServiceList, error) {
	ret := _m.Called(options)

	var r0 *v1alpha1.ServiceList
	if rf, ok := ret.Get(0).(func(core.ListServiceOptions) *v1alpha1.ServiceList); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.ServiceList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.ListServiceOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamespaceCleanup provides a mock function with given fields: options
func (_m *Client) NamespaceCleanup(options core.NamespaceCleanupOptions) error {
	ret := _m.Called(options)

	var r0 error
	if rf, ok := ret.Get(0).(func(core.NamespaceCleanupOptions) error); ok {
		r0 = rf(options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NamespaceInit provides a mock function with given fields: manifests, options
func (_m *Client) NamespaceInit(manifests map[string]*core.Manifest, options core.NamespaceInitOptions) error {
	ret := _m.Called(manifests, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]*core.Manifest, core.NamespaceInitOptions) error); ok {
		r0 = rf(manifests, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServiceCoordinates provides a mock function with given fields: options
func (_m *Client) ServiceCoordinates(options core.ServiceInvokeOptions) (string, string, error) {
	ret := _m.Called(options)

	var r0 string
	if rf, ok := ret.Get(0).(func(core.ServiceInvokeOptions) string); ok {
		r0 = rf(options)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(core.ServiceInvokeOptions) string); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(core.ServiceInvokeOptions) error); ok {
		r2 = rf(options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ServiceStatus provides a mock function with given fields: options
func (_m *Client) ServiceStatus(options core.ServiceStatusOptions) (*duckv1alpha1.Condition, error) {
	ret := _m.Called(options)

	var r0 *duckv1alpha1.Condition
	if rf, ok := ret.Get(0).(func(core.ServiceStatusOptions) *duckv1alpha1.Condition); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*duckv1alpha1.Condition)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.ServiceStatusOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetDefaultBuildImagePrefix provides a mock function with given fields: namespace, prefix
func (_m *Client) SetDefaultBuildImagePrefix(namespace string, prefix string) error {
	ret := _m.Called(namespace, prefix)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(namespace, prefix)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SystemInstall provides a mock function with given fields: manifests, options
func (_m *Client) SystemInstall(manifests map[string]*core.Manifest, options core.SystemInstallOptions) (bool, error) {
	ret := _m.Called(manifests, options)

	var r0 bool
	if rf, ok := ret.Get(0).(func(map[string]*core.Manifest, core.SystemInstallOptions) bool); ok {
		r0 = rf(manifests, options)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]*core.Manifest, core.SystemInstallOptions) error); ok {
		r1 = rf(manifests, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemUninstall provides a mock function with given fields: options
func (_m *Client) SystemUninstall(options core.SystemUninstallOptions) (bool, error) {
	ret := _m.Called(options)

	var r0 bool
	if rf, ok := ret.Get(0).(func(core.SystemUninstallOptions) bool); ok {
		r0 = rf(options)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.SystemUninstallOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFunction provides a mock function with given fields: builder, options, log
func (_m *Client) UpdateFunction(builder core.Builder, options core.UpdateFunctionOptions, log io.Writer) error {
	ret := _m.Called(builder, options, log)

	var r0 error
	if rf, ok := ret.Get(0).(func(core.Builder, core.UpdateFunctionOptions, io.Writer) error); ok {
		r0 = rf(builder, options, log)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateService provides a mock function with given fields: options
func (_m *Client) UpdateService(options core.CreateOrUpdateServiceOptions) (*v1alpha1.Service, error) {
	ret := _m.Called(options)

	var r0 *v1alpha1.Service
	if rf, ok := ret.Get(0).(func(core.CreateOrUpdateServiceOptions) *v1alpha1.Service); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.CreateOrUpdateServiceOptions) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
