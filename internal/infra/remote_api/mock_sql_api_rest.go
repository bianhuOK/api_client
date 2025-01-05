package remoteapi

import (
	"github.com/bianhuOK/api_client/internal/domain/sql_template"
	"github.com/stretchr/testify/mock"
)

// MockRemoteAPI is a mock implementation of the RemoteAPI interface
type MockRemoteAPI struct {
	mock.Mock
}

func NewMockRemoteAPI() *MockRemoteAPI {
	m := MockRemoteAPI{}
	// Use mock's Method to simulate the behavior
	expectedTemplate := &sql_template.SqlTemplate{
		ApiId:           "id123",
		TemplateContent: "SELECT * FROM users WHERE id = ?",
	}

	// Setup expectations
	m.On("FetchTemplate", "123").Return(expectedTemplate, nil)
	return &m
}

// FetchTemplate is the mock implementation of the FetchTemplate method
func (m *MockRemoteAPI) FetchTemplate(id string) (*sql_template.SqlTemplate, error) {
	args := m.Called(id)
	return args.Get(0).(*sql_template.SqlTemplate), args.Error(1)
}
