/*
Copyright 2021 The Dapr Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by mockery v2.9.4. DO NOT EDIT.

package testing

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	configuration "github.com/dapr/components-contrib/configuration"
)

// MockConfigurationStore is an autogenerated mock type for the Store type
type MockConfigurationStore struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, req
func (_m *MockConfigurationStore) Get(ctx context.Context, req *configuration.GetRequest) (*configuration.GetResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *configuration.GetResponse
	if rf, ok := ret.Get(0).(func(context.Context, *configuration.GetRequest) *configuration.GetResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*configuration.GetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *configuration.GetRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields: metadata
func (_m *MockConfigurationStore) Init(ctx context.Context, metadata configuration.Metadata) error {
	ret := _m.Called(metadata)

	var r0 error
	if rf, ok := ret.Get(0).(func(configuration.Metadata) error); ok {
		r0 = rf(metadata)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: ctx, req, handler
func (_m *MockConfigurationStore) Subscribe(ctx context.Context, req *configuration.SubscribeRequest, handler configuration.UpdateHandler) (string, error) {
	ret := _m.Called(ctx, req, handler)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *configuration.SubscribeRequest, configuration.UpdateHandler) string); ok {
		r0 = rf(ctx, req, handler)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *configuration.SubscribeRequest, configuration.UpdateHandler) error); ok {
		r1 = rf(ctx, req, handler)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unsubscribe provides a mock function with given fields: ctx, req
func (_m *MockConfigurationStore) Unsubscribe(ctx context.Context, req *configuration.UnsubscribeRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *configuration.UnsubscribeRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type FailingConfigurationStore struct {
	Failure Failure
}

func (f *FailingConfigurationStore) Get(ctx context.Context, req *configuration.GetRequest) (*configuration.GetResponse, error) {
	if err := f.Failure.PerformFailure(req.Metadata["key"]); err != nil {
		return nil, err
	}
	return &configuration.GetResponse{}, nil
}

func (f *FailingConfigurationStore) Init(ctx context.Context, metadata configuration.Metadata) error {
	return nil
}

func (f *FailingConfigurationStore) Subscribe(ctx context.Context, req *configuration.SubscribeRequest, handler configuration.UpdateHandler) (string, error) {
	if err := f.Failure.PerformFailure(req.Metadata["key"]); err != nil {
		return "", err
	}

	handler(ctx, &configuration.UpdateEvent{
		Items: map[string]*configuration.Item{
			req.Metadata["key"]: {
				Value: "testConfig",
			},
		},
	})
	
	return "subscribeID", nil
}

func (f *FailingConfigurationStore) Unsubscribe(ctx context.Context, req *configuration.UnsubscribeRequest) error {
	return f.Failure.PerformFailure(req.ID)
}
