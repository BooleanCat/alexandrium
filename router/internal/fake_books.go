// Code generated by counterfeiter. DO NOT EDIT.
package internal

import (
	"sync"

	"github.com/BooleanCat/alexandrium/books"
)

type FakeBooks struct {
	ByIDStub        func(string) (books.Book, error)
	byIDMutex       sync.RWMutex
	byIDArgsForCall []struct {
		arg1 string
	}
	byIDReturns struct {
		result1 books.Book
		result2 error
	}
	byIDReturnsOnCall map[int]struct {
		result1 books.Book
		result2 error
	}
	ByISBNStub        func(string) (books.Book, error)
	byISBNMutex       sync.RWMutex
	byISBNArgsForCall []struct {
		arg1 string
	}
	byISBNReturns struct {
		result1 books.Book
		result2 error
	}
	byISBNReturnsOnCall map[int]struct {
		result1 books.Book
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBooks) ByID(arg1 string) (books.Book, error) {
	fake.byIDMutex.Lock()
	ret, specificReturn := fake.byIDReturnsOnCall[len(fake.byIDArgsForCall)]
	fake.byIDArgsForCall = append(fake.byIDArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ByID", []interface{}{arg1})
	fake.byIDMutex.Unlock()
	if fake.ByIDStub != nil {
		return fake.ByIDStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.byIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBooks) ByIDCallCount() int {
	fake.byIDMutex.RLock()
	defer fake.byIDMutex.RUnlock()
	return len(fake.byIDArgsForCall)
}

func (fake *FakeBooks) ByIDCalls(stub func(string) (books.Book, error)) {
	fake.byIDMutex.Lock()
	defer fake.byIDMutex.Unlock()
	fake.ByIDStub = stub
}

func (fake *FakeBooks) ByIDArgsForCall(i int) string {
	fake.byIDMutex.RLock()
	defer fake.byIDMutex.RUnlock()
	argsForCall := fake.byIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBooks) ByIDReturns(result1 books.Book, result2 error) {
	fake.byIDMutex.Lock()
	defer fake.byIDMutex.Unlock()
	fake.ByIDStub = nil
	fake.byIDReturns = struct {
		result1 books.Book
		result2 error
	}{result1, result2}
}

func (fake *FakeBooks) ByIDReturnsOnCall(i int, result1 books.Book, result2 error) {
	fake.byIDMutex.Lock()
	defer fake.byIDMutex.Unlock()
	fake.ByIDStub = nil
	if fake.byIDReturnsOnCall == nil {
		fake.byIDReturnsOnCall = make(map[int]struct {
			result1 books.Book
			result2 error
		})
	}
	fake.byIDReturnsOnCall[i] = struct {
		result1 books.Book
		result2 error
	}{result1, result2}
}

func (fake *FakeBooks) ByISBN(arg1 string) (books.Book, error) {
	fake.byISBNMutex.Lock()
	ret, specificReturn := fake.byISBNReturnsOnCall[len(fake.byISBNArgsForCall)]
	fake.byISBNArgsForCall = append(fake.byISBNArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ByISBN", []interface{}{arg1})
	fake.byISBNMutex.Unlock()
	if fake.ByISBNStub != nil {
		return fake.ByISBNStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.byISBNReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBooks) ByISBNCallCount() int {
	fake.byISBNMutex.RLock()
	defer fake.byISBNMutex.RUnlock()
	return len(fake.byISBNArgsForCall)
}

func (fake *FakeBooks) ByISBNCalls(stub func(string) (books.Book, error)) {
	fake.byISBNMutex.Lock()
	defer fake.byISBNMutex.Unlock()
	fake.ByISBNStub = stub
}

func (fake *FakeBooks) ByISBNArgsForCall(i int) string {
	fake.byISBNMutex.RLock()
	defer fake.byISBNMutex.RUnlock()
	argsForCall := fake.byISBNArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBooks) ByISBNReturns(result1 books.Book, result2 error) {
	fake.byISBNMutex.Lock()
	defer fake.byISBNMutex.Unlock()
	fake.ByISBNStub = nil
	fake.byISBNReturns = struct {
		result1 books.Book
		result2 error
	}{result1, result2}
}

func (fake *FakeBooks) ByISBNReturnsOnCall(i int, result1 books.Book, result2 error) {
	fake.byISBNMutex.Lock()
	defer fake.byISBNMutex.Unlock()
	fake.ByISBNStub = nil
	if fake.byISBNReturnsOnCall == nil {
		fake.byISBNReturnsOnCall = make(map[int]struct {
			result1 books.Book
			result2 error
		})
	}
	fake.byISBNReturnsOnCall[i] = struct {
		result1 books.Book
		result2 error
	}{result1, result2}
}

func (fake *FakeBooks) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.byIDMutex.RLock()
	defer fake.byIDMutex.RUnlock()
	fake.byISBNMutex.RLock()
	defer fake.byISBNMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBooks) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ books.Books = new(FakeBooks)
