package context

import (
	"testing"
	"github.com/stretchr/testify/mock"
)

// ContextMock is
type ContextMock struct {
	mock.Mock
}

// New is
func (m *ContextMock) New(port string, remoteHost string) *AppContext {
	args := m.Called(string)
  	return &AppContext{}
}
