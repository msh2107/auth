package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/msh2107/auth/internal/service.UserService -o ./mocks/user_service_minimock.go -n UserServiceMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/msh2107/auth/internal/model"
)

// UserServiceMock implements service.UserService
type UserServiceMock struct {
	t minimock.Tester

	funcCreate          func(ctx context.Context, info *model.UserInfo) (i1 int64, err error)
	inspectFuncCreate   func(ctx context.Context, info *model.UserInfo)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mUserServiceMockCreate

	funcDelete          func(ctx context.Context, id int64) (err error)
	inspectFuncDelete   func(ctx context.Context, id int64)
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mUserServiceMockDelete

	funcGet          func(ctx context.Context, id int64) (up1 *model.User, err error)
	inspectFuncGet   func(ctx context.Context, id int64)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mUserServiceMockGet

	funcUpdate          func(ctx context.Context, id int64, info *model.UserInfo) (err error)
	inspectFuncUpdate   func(ctx context.Context, id int64, info *model.UserInfo)
	afterUpdateCounter  uint64
	beforeUpdateCounter uint64
	UpdateMock          mUserServiceMockUpdate
}

// NewUserServiceMock returns a mock for service.UserService
func NewUserServiceMock(t minimock.Tester) *UserServiceMock {
	m := &UserServiceMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mUserServiceMockCreate{mock: m}
	m.CreateMock.callArgs = []*UserServiceMockCreateParams{}

	m.DeleteMock = mUserServiceMockDelete{mock: m}
	m.DeleteMock.callArgs = []*UserServiceMockDeleteParams{}

	m.GetMock = mUserServiceMockGet{mock: m}
	m.GetMock.callArgs = []*UserServiceMockGetParams{}

	m.UpdateMock = mUserServiceMockUpdate{mock: m}
	m.UpdateMock.callArgs = []*UserServiceMockUpdateParams{}

	return m
}

type mUserServiceMockCreate struct {
	mock               *UserServiceMock
	defaultExpectation *UserServiceMockCreateExpectation
	expectations       []*UserServiceMockCreateExpectation

	callArgs []*UserServiceMockCreateParams
	mutex    sync.RWMutex
}

// UserServiceMockCreateExpectation specifies expectation struct of the UserService.Create
type UserServiceMockCreateExpectation struct {
	mock    *UserServiceMock
	params  *UserServiceMockCreateParams
	results *UserServiceMockCreateResults
	Counter uint64
}

// UserServiceMockCreateParams contains parameters of the UserService.Create
type UserServiceMockCreateParams struct {
	ctx  context.Context
	info *model.UserInfo
}

// UserServiceMockCreateResults contains results of the UserService.Create
type UserServiceMockCreateResults struct {
	i1  int64
	err error
}

// Expect sets up expected params for UserService.Create
func (mmCreate *mUserServiceMockCreate) Expect(ctx context.Context, info *model.UserInfo) *mUserServiceMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserServiceMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UserServiceMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &UserServiceMockCreateParams{ctx, info}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the UserService.Create
func (mmCreate *mUserServiceMockCreate) Inspect(f func(ctx context.Context, info *model.UserInfo)) *mUserServiceMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for UserServiceMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by UserService.Create
func (mmCreate *mUserServiceMockCreate) Return(i1 int64, err error) *UserServiceMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserServiceMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UserServiceMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &UserServiceMockCreateResults{i1, err}
	return mmCreate.mock
}

// Set uses given function f to mock the UserService.Create method
func (mmCreate *mUserServiceMockCreate) Set(f func(ctx context.Context, info *model.UserInfo) (i1 int64, err error)) *UserServiceMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the UserService.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the UserService.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the UserService.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mUserServiceMockCreate) When(ctx context.Context, info *model.UserInfo) *UserServiceMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserServiceMock.Create mock is already set by Set")
	}

	expectation := &UserServiceMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &UserServiceMockCreateParams{ctx, info},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up UserService.Create return parameters for the expectation previously defined by the When method
func (e *UserServiceMockCreateExpectation) Then(i1 int64, err error) *UserServiceMock {
	e.results = &UserServiceMockCreateResults{i1, err}
	return e.mock
}

// Create implements service.UserService
func (mmCreate *UserServiceMock) Create(ctx context.Context, info *model.UserInfo) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, info)
	}

	mm_params := &UserServiceMockCreateParams{ctx, info}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := UserServiceMockCreateParams{ctx, info}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("UserServiceMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the UserServiceMock.Create")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, info)
	}
	mmCreate.t.Fatalf("Unexpected call to UserServiceMock.Create. %v %v", ctx, info)
	return
}

// CreateAfterCounter returns a count of finished UserServiceMock.Create invocations
func (mmCreate *UserServiceMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of UserServiceMock.Create invocations
func (mmCreate *UserServiceMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to UserServiceMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mUserServiceMockCreate) Calls() []*UserServiceMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*UserServiceMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *UserServiceMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *UserServiceMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserServiceMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserServiceMock.Create")
		} else {
			m.t.Errorf("Expected call to UserServiceMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to UserServiceMock.Create")
	}
}

type mUserServiceMockDelete struct {
	mock               *UserServiceMock
	defaultExpectation *UserServiceMockDeleteExpectation
	expectations       []*UserServiceMockDeleteExpectation

	callArgs []*UserServiceMockDeleteParams
	mutex    sync.RWMutex
}

// UserServiceMockDeleteExpectation specifies expectation struct of the UserService.Delete
type UserServiceMockDeleteExpectation struct {
	mock    *UserServiceMock
	params  *UserServiceMockDeleteParams
	results *UserServiceMockDeleteResults
	Counter uint64
}

// UserServiceMockDeleteParams contains parameters of the UserService.Delete
type UserServiceMockDeleteParams struct {
	ctx context.Context
	id  int64
}

// UserServiceMockDeleteResults contains results of the UserService.Delete
type UserServiceMockDeleteResults struct {
	err error
}

// Expect sets up expected params for UserService.Delete
func (mmDelete *mUserServiceMockDelete) Expect(ctx context.Context, id int64) *mUserServiceMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("UserServiceMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &UserServiceMockDeleteExpectation{}
	}

	mmDelete.defaultExpectation.params = &UserServiceMockDeleteParams{ctx, id}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the UserService.Delete
func (mmDelete *mUserServiceMockDelete) Inspect(f func(ctx context.Context, id int64)) *mUserServiceMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for UserServiceMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by UserService.Delete
func (mmDelete *mUserServiceMockDelete) Return(err error) *UserServiceMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("UserServiceMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &UserServiceMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &UserServiceMockDeleteResults{err}
	return mmDelete.mock
}

// Set uses given function f to mock the UserService.Delete method
func (mmDelete *mUserServiceMockDelete) Set(f func(ctx context.Context, id int64) (err error)) *UserServiceMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the UserService.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the UserService.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the UserService.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mUserServiceMockDelete) When(ctx context.Context, id int64) *UserServiceMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("UserServiceMock.Delete mock is already set by Set")
	}

	expectation := &UserServiceMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &UserServiceMockDeleteParams{ctx, id},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up UserService.Delete return parameters for the expectation previously defined by the When method
func (e *UserServiceMockDeleteExpectation) Then(err error) *UserServiceMock {
	e.results = &UserServiceMockDeleteResults{err}
	return e.mock
}

// Delete implements service.UserService
func (mmDelete *UserServiceMock) Delete(ctx context.Context, id int64) (err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(ctx, id)
	}

	mm_params := &UserServiceMockDeleteParams{ctx, id}

	// Record call args
	mmDelete.DeleteMock.mutex.Lock()
	mmDelete.DeleteMock.callArgs = append(mmDelete.DeleteMock.callArgs, mm_params)
	mmDelete.DeleteMock.mutex.Unlock()

	for _, e := range mmDelete.DeleteMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDelete.DeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDelete.DeleteMock.defaultExpectation.Counter, 1)
		mm_want := mmDelete.DeleteMock.defaultExpectation.params
		mm_got := UserServiceMockDeleteParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDelete.t.Errorf("UserServiceMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDelete.DeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmDelete.t.Fatal("No results are set for the UserServiceMock.Delete")
		}
		return (*mm_results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(ctx, id)
	}
	mmDelete.t.Fatalf("Unexpected call to UserServiceMock.Delete. %v %v", ctx, id)
	return
}

// DeleteAfterCounter returns a count of finished UserServiceMock.Delete invocations
func (mmDelete *UserServiceMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of UserServiceMock.Delete invocations
func (mmDelete *UserServiceMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to UserServiceMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mUserServiceMockDelete) Calls() []*UserServiceMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*UserServiceMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *UserServiceMock) MinimockDeleteDone() bool {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteInspect logs each unmet expectation
func (m *UserServiceMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserServiceMock.Delete with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserServiceMock.Delete")
		} else {
			m.t.Errorf("Expected call to UserServiceMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		m.t.Error("Expected call to UserServiceMock.Delete")
	}
}

type mUserServiceMockGet struct {
	mock               *UserServiceMock
	defaultExpectation *UserServiceMockGetExpectation
	expectations       []*UserServiceMockGetExpectation

	callArgs []*UserServiceMockGetParams
	mutex    sync.RWMutex
}

// UserServiceMockGetExpectation specifies expectation struct of the UserService.Get
type UserServiceMockGetExpectation struct {
	mock    *UserServiceMock
	params  *UserServiceMockGetParams
	results *UserServiceMockGetResults
	Counter uint64
}

// UserServiceMockGetParams contains parameters of the UserService.Get
type UserServiceMockGetParams struct {
	ctx context.Context
	id  int64
}

// UserServiceMockGetResults contains results of the UserService.Get
type UserServiceMockGetResults struct {
	up1 *model.User
	err error
}

// Expect sets up expected params for UserService.Get
func (mmGet *mUserServiceMockGet) Expect(ctx context.Context, id int64) *mUserServiceMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserServiceMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserServiceMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &UserServiceMockGetParams{ctx, id}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the UserService.Get
func (mmGet *mUserServiceMockGet) Inspect(f func(ctx context.Context, id int64)) *mUserServiceMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for UserServiceMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by UserService.Get
func (mmGet *mUserServiceMockGet) Return(up1 *model.User, err error) *UserServiceMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserServiceMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserServiceMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &UserServiceMockGetResults{up1, err}
	return mmGet.mock
}

// Set uses given function f to mock the UserService.Get method
func (mmGet *mUserServiceMockGet) Set(f func(ctx context.Context, id int64) (up1 *model.User, err error)) *UserServiceMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the UserService.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the UserService.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the UserService.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mUserServiceMockGet) When(ctx context.Context, id int64) *UserServiceMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserServiceMock.Get mock is already set by Set")
	}

	expectation := &UserServiceMockGetExpectation{
		mock:   mmGet.mock,
		params: &UserServiceMockGetParams{ctx, id},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up UserService.Get return parameters for the expectation previously defined by the When method
func (e *UserServiceMockGetExpectation) Then(up1 *model.User, err error) *UserServiceMock {
	e.results = &UserServiceMockGetResults{up1, err}
	return e.mock
}

// Get implements service.UserService
func (mmGet *UserServiceMock) Get(ctx context.Context, id int64) (up1 *model.User, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(ctx, id)
	}

	mm_params := &UserServiceMockGetParams{ctx, id}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.up1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_got := UserServiceMockGetParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("UserServiceMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the UserServiceMock.Get")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(ctx, id)
	}
	mmGet.t.Fatalf("Unexpected call to UserServiceMock.Get. %v %v", ctx, id)
	return
}

// GetAfterCounter returns a count of finished UserServiceMock.Get invocations
func (mmGet *UserServiceMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of UserServiceMock.Get invocations
func (mmGet *UserServiceMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to UserServiceMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mUserServiceMockGet) Calls() []*UserServiceMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*UserServiceMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *UserServiceMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *UserServiceMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserServiceMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserServiceMock.Get")
		} else {
			m.t.Errorf("Expected call to UserServiceMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to UserServiceMock.Get")
	}
}

type mUserServiceMockUpdate struct {
	mock               *UserServiceMock
	defaultExpectation *UserServiceMockUpdateExpectation
	expectations       []*UserServiceMockUpdateExpectation

	callArgs []*UserServiceMockUpdateParams
	mutex    sync.RWMutex
}

// UserServiceMockUpdateExpectation specifies expectation struct of the UserService.Update
type UserServiceMockUpdateExpectation struct {
	mock    *UserServiceMock
	params  *UserServiceMockUpdateParams
	results *UserServiceMockUpdateResults
	Counter uint64
}

// UserServiceMockUpdateParams contains parameters of the UserService.Update
type UserServiceMockUpdateParams struct {
	ctx  context.Context
	id   int64
	info *model.UserInfo
}

// UserServiceMockUpdateResults contains results of the UserService.Update
type UserServiceMockUpdateResults struct {
	err error
}

// Expect sets up expected params for UserService.Update
func (mmUpdate *mUserServiceMockUpdate) Expect(ctx context.Context, id int64, info *model.UserInfo) *mUserServiceMockUpdate {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("UserServiceMock.Update mock is already set by Set")
	}

	if mmUpdate.defaultExpectation == nil {
		mmUpdate.defaultExpectation = &UserServiceMockUpdateExpectation{}
	}

	mmUpdate.defaultExpectation.params = &UserServiceMockUpdateParams{ctx, id, info}
	for _, e := range mmUpdate.expectations {
		if minimock.Equal(e.params, mmUpdate.defaultExpectation.params) {
			mmUpdate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmUpdate.defaultExpectation.params)
		}
	}

	return mmUpdate
}

// Inspect accepts an inspector function that has same arguments as the UserService.Update
func (mmUpdate *mUserServiceMockUpdate) Inspect(f func(ctx context.Context, id int64, info *model.UserInfo)) *mUserServiceMockUpdate {
	if mmUpdate.mock.inspectFuncUpdate != nil {
		mmUpdate.mock.t.Fatalf("Inspect function is already set for UserServiceMock.Update")
	}

	mmUpdate.mock.inspectFuncUpdate = f

	return mmUpdate
}

// Return sets up results that will be returned by UserService.Update
func (mmUpdate *mUserServiceMockUpdate) Return(err error) *UserServiceMock {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("UserServiceMock.Update mock is already set by Set")
	}

	if mmUpdate.defaultExpectation == nil {
		mmUpdate.defaultExpectation = &UserServiceMockUpdateExpectation{mock: mmUpdate.mock}
	}
	mmUpdate.defaultExpectation.results = &UserServiceMockUpdateResults{err}
	return mmUpdate.mock
}

// Set uses given function f to mock the UserService.Update method
func (mmUpdate *mUserServiceMockUpdate) Set(f func(ctx context.Context, id int64, info *model.UserInfo) (err error)) *UserServiceMock {
	if mmUpdate.defaultExpectation != nil {
		mmUpdate.mock.t.Fatalf("Default expectation is already set for the UserService.Update method")
	}

	if len(mmUpdate.expectations) > 0 {
		mmUpdate.mock.t.Fatalf("Some expectations are already set for the UserService.Update method")
	}

	mmUpdate.mock.funcUpdate = f
	return mmUpdate.mock
}

// When sets expectation for the UserService.Update which will trigger the result defined by the following
// Then helper
func (mmUpdate *mUserServiceMockUpdate) When(ctx context.Context, id int64, info *model.UserInfo) *UserServiceMockUpdateExpectation {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("UserServiceMock.Update mock is already set by Set")
	}

	expectation := &UserServiceMockUpdateExpectation{
		mock:   mmUpdate.mock,
		params: &UserServiceMockUpdateParams{ctx, id, info},
	}
	mmUpdate.expectations = append(mmUpdate.expectations, expectation)
	return expectation
}

// Then sets up UserService.Update return parameters for the expectation previously defined by the When method
func (e *UserServiceMockUpdateExpectation) Then(err error) *UserServiceMock {
	e.results = &UserServiceMockUpdateResults{err}
	return e.mock
}

// Update implements service.UserService
func (mmUpdate *UserServiceMock) Update(ctx context.Context, id int64, info *model.UserInfo) (err error) {
	mm_atomic.AddUint64(&mmUpdate.beforeUpdateCounter, 1)
	defer mm_atomic.AddUint64(&mmUpdate.afterUpdateCounter, 1)

	if mmUpdate.inspectFuncUpdate != nil {
		mmUpdate.inspectFuncUpdate(ctx, id, info)
	}

	mm_params := &UserServiceMockUpdateParams{ctx, id, info}

	// Record call args
	mmUpdate.UpdateMock.mutex.Lock()
	mmUpdate.UpdateMock.callArgs = append(mmUpdate.UpdateMock.callArgs, mm_params)
	mmUpdate.UpdateMock.mutex.Unlock()

	for _, e := range mmUpdate.UpdateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmUpdate.UpdateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmUpdate.UpdateMock.defaultExpectation.Counter, 1)
		mm_want := mmUpdate.UpdateMock.defaultExpectation.params
		mm_got := UserServiceMockUpdateParams{ctx, id, info}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmUpdate.t.Errorf("UserServiceMock.Update got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmUpdate.UpdateMock.defaultExpectation.results
		if mm_results == nil {
			mmUpdate.t.Fatal("No results are set for the UserServiceMock.Update")
		}
		return (*mm_results).err
	}
	if mmUpdate.funcUpdate != nil {
		return mmUpdate.funcUpdate(ctx, id, info)
	}
	mmUpdate.t.Fatalf("Unexpected call to UserServiceMock.Update. %v %v %v", ctx, id, info)
	return
}

// UpdateAfterCounter returns a count of finished UserServiceMock.Update invocations
func (mmUpdate *UserServiceMock) UpdateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdate.afterUpdateCounter)
}

// UpdateBeforeCounter returns a count of UserServiceMock.Update invocations
func (mmUpdate *UserServiceMock) UpdateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdate.beforeUpdateCounter)
}

// Calls returns a list of arguments used in each call to UserServiceMock.Update.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmUpdate *mUserServiceMockUpdate) Calls() []*UserServiceMockUpdateParams {
	mmUpdate.mutex.RLock()

	argCopy := make([]*UserServiceMockUpdateParams, len(mmUpdate.callArgs))
	copy(argCopy, mmUpdate.callArgs)

	mmUpdate.mutex.RUnlock()

	return argCopy
}

// MinimockUpdateDone returns true if the count of the Update invocations corresponds
// the number of defined expectations
func (m *UserServiceMock) MinimockUpdateDone() bool {
	for _, e := range m.UpdateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdate != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		return false
	}
	return true
}

// MinimockUpdateInspect logs each unmet expectation
func (m *UserServiceMock) MinimockUpdateInspect() {
	for _, e := range m.UpdateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserServiceMock.Update with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		if m.UpdateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserServiceMock.Update")
		} else {
			m.t.Errorf("Expected call to UserServiceMock.Update with params: %#v", *m.UpdateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdate != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		m.t.Error("Expected call to UserServiceMock.Update")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UserServiceMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreateInspect()

		m.MinimockDeleteInspect()

		m.MinimockGetInspect()

		m.MinimockUpdateInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UserServiceMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *UserServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockDeleteDone() &&
		m.MinimockGetDone() &&
		m.MinimockUpdateDone()
}
