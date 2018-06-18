package mocks

import event "flamingo.me/flamingo/framework/event"
import http "net/http"
import mock "github.com/stretchr/testify/mock"
import profiler "flamingo.me/flamingo/framework/profiler"
import sessions "github.com/gorilla/sessions"
import time "time"
import web "flamingo.me/flamingo/framework/web"

// Context is an autogenerated mock type for the Context type
type Context struct {
	mock.Mock
}

// Deadline provides a mock function with given fields:
func (_m *Context) Deadline() (time.Time, bool) {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Done provides a mock function with given fields:
func (_m *Context) Done() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Err provides a mock function with given fields:
func (_m *Context) Err() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EventRouter provides a mock function with given fields:
func (_m *Context) EventRouter() event.Router {
	ret := _m.Called()

	var r0 event.Router
	if rf, ok := ret.Get(0).(func() event.Router); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Router)
		}
	}

	return r0
}

// Form provides a mock function with given fields: _a0
func (_m *Context) Form(_a0 string) ([]string, error) {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Form1 provides a mock function with given fields: _a0
func (_m *Context) Form1(_a0 string) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FormAll provides a mock function with given fields:
func (_m *Context) FormAll() map[string][]string {
	ret := _m.Called()

	var r0 map[string][]string
	if rf, ok := ret.Get(0).(func() map[string][]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]string)
		}
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *Context) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// LoadParams provides a mock function with given fields: p
func (_m *Context) LoadParams(p map[string]string) {
	_m.Called(p)
}

// MustForm provides a mock function with given fields: _a0
func (_m *Context) MustForm(_a0 string) []string {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MustForm1 provides a mock function with given fields: _a0
func (_m *Context) MustForm1(_a0 string) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MustParam1 provides a mock function with given fields: _a0
func (_m *Context) MustParam1(_a0 string) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MustQuery provides a mock function with given fields: _a0
func (_m *Context) MustQuery(_a0 string) []string {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MustQuery1 provides a mock function with given fields: _a0
func (_m *Context) MustQuery1(_a0 string) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Param1 provides a mock function with given fields: _a0
func (_m *Context) Param1(_a0 string) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParamAll provides a mock function with given fields:
func (_m *Context) ParamAll() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// Profile provides a mock function with given fields: key, msg
func (_m *Context) Profile(key string, msg string) profiler.ProfileFinishFunc {
	ret := _m.Called(key, msg)

	var r0 profiler.ProfileFinishFunc
	if rf, ok := ret.Get(0).(func(string, string) profiler.ProfileFinishFunc); ok {
		r0 = rf(key, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(profiler.ProfileFinishFunc)
		}
	}

	return r0
}

// Profiler provides a mock function with given fields:
func (_m *Context) Profiler() profiler.Profiler {
	ret := _m.Called()

	var r0 profiler.Profiler
	if rf, ok := ret.Get(0).(func() profiler.Profiler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(profiler.Profiler)
		}
	}

	return r0
}

// Push provides a mock function with given fields: target, opts
func (_m *Context) Push(target string, opts *http.PushOptions) error {
	ret := _m.Called(target, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *http.PushOptions) error); ok {
		r0 = rf(target, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Query provides a mock function with given fields: _a0
func (_m *Context) Query(_a0 string) ([]string, error) {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Query1 provides a mock function with given fields: _a0
func (_m *Context) Query1(_a0 string) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryAll provides a mock function with given fields:
func (_m *Context) QueryAll() map[string][]string {
	ret := _m.Called()

	var r0 map[string][]string
	if rf, ok := ret.Get(0).(func() map[string][]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]string)
		}
	}

	return r0
}

// Request provides a mock function with given fields:
func (_m *Context) Request() *http.Request {
	ret := _m.Called()

	var r0 *http.Request
	if rf, ok := ret.Get(0).(func() *http.Request); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Request)
		}
	}

	return r0
}

// Session provides a mock function with given fields:
func (_m *Context) Session() *sessions.Session {
	ret := _m.Called()

	var r0 *sessions.Session
	if rf, ok := ret.Get(0).(func() *sessions.Session); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sessions.Session)
		}
	}

	return r0
}

// Value provides a mock function with given fields: key
func (_m *Context) Value(key interface{}) interface{} {
	ret := _m.Called(key)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// WithValue provides a mock function with given fields: key, value
func (_m *Context) WithValue(key interface{}, value interface{}) web.Context {
	ret := _m.Called(key, value)

	var r0 web.Context
	if rf, ok := ret.Get(0).(func(interface{}, interface{}) web.Context); ok {
		r0 = rf(key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(web.Context)
		}
	}

	return r0
}

// WithVars provides a mock function with given fields: vars
func (_m *Context) WithVars(vars map[string]string) web.Context {
	ret := _m.Called(vars)

	var r0 web.Context
	if rf, ok := ret.Get(0).(func(map[string]string) web.Context); ok {
		r0 = rf(vars)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(web.Context)
		}
	}

	return r0
}
