// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	installation "github.com/kyma-incubator/hydroform/install/installation"

	mock "github.com/stretchr/testify/mock"

	model "github.com/kyma-project/control-plane/components/provisioner/internal/model"

	rest "k8s.io/client-go/rest"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CheckInstallationState provides a mock function with given fields: kubeconfig
func (_m *Service) CheckInstallationState(kubeconfig *rest.Config) (installation.InstallationState, error) {
	ret := _m.Called(kubeconfig)

	var r0 installation.InstallationState
	if rf, ok := ret.Get(0).(func(*rest.Config) installation.InstallationState); ok {
		r0 = rf(kubeconfig)
	} else {
		r0 = ret.Get(0).(installation.InstallationState)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*rest.Config) error); ok {
		r1 = rf(kubeconfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InstallKyma provides a mock function with given fields: runtimeId, kubeconfigRaw, release, globalConfig, componentsConfig
func (_m *Service) InstallKyma(runtimeId string, kubeconfigRaw string, release model.Release, globalConfig model.Configuration, componentsConfig []model.KymaComponentConfig) error {
	ret := _m.Called(runtimeId, kubeconfigRaw, release, globalConfig, componentsConfig)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, model.Release, model.Configuration, []model.KymaComponentConfig) error); ok {
		r0 = rf(runtimeId, kubeconfigRaw, release, globalConfig, componentsConfig)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PerformCleanup provides a mock function with given fields: kubeconfig
func (_m *Service) PerformCleanup(kubeconfig *rest.Config) error {
	ret := _m.Called(kubeconfig)

	var r0 error
	if rf, ok := ret.Get(0).(func(*rest.Config) error); ok {
		r0 = rf(kubeconfig)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TriggerInstallation provides a mock function with given fields: kubeconfigRaw, release, globalConfig, componentsConfig
func (_m *Service) TriggerInstallation(kubeconfigRaw *rest.Config, release model.Release, globalConfig model.Configuration, componentsConfig []model.KymaComponentConfig) error {
	ret := _m.Called(kubeconfigRaw, release, globalConfig, componentsConfig)

	var r0 error
	if rf, ok := ret.Get(0).(func(*rest.Config, model.Release, model.Configuration, []model.KymaComponentConfig) error); ok {
		r0 = rf(kubeconfigRaw, release, globalConfig, componentsConfig)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TriggerUninstall provides a mock function with given fields: kubeconfig
func (_m *Service) TriggerUninstall(kubeconfig *rest.Config) error {
	ret := _m.Called(kubeconfig)

	var r0 error
	if rf, ok := ret.Get(0).(func(*rest.Config) error); ok {
		r0 = rf(kubeconfig)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TriggerUpgrade provides a mock function with given fields: kubeconfigRaw, release, globalConfig, componentsConfig
func (_m *Service) TriggerUpgrade(kubeconfigRaw *rest.Config, release model.Release, globalConfig model.Configuration, componentsConfig []model.KymaComponentConfig) error {
	ret := _m.Called(kubeconfigRaw, release, globalConfig, componentsConfig)

	var r0 error
	if rf, ok := ret.Get(0).(func(*rest.Config, model.Release, model.Configuration, []model.KymaComponentConfig) error); ok {
		r0 = rf(kubeconfigRaw, release, globalConfig, componentsConfig)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
