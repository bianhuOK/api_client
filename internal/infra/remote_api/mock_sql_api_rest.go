package remoteapi

import (
	"github.com/bianhuOK/api_client/internal/domain/model"
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
		ApiId:           "zzz",
		TemplateContent: "SELECT * FROM employees WHERE employee_id = 1",
		DbConfig: model.DbConfig{
			Host:     "49.233.45.238",
			Port:     3306,
			User:     "yep",
			PassWord: "A_bcd0932",
			DbName:   "mydatabase",
			Charset:  "utf8mb4",
			DbType:   "mysql",
			DSN:      "yep:A_bcd0932@tcp(49.233.45.238:3306)/mydatabase?charset=utf8mb4",
		},
	}

	// Setup expectations
	m.On("FetchTemplate", mock.Anything).Return(expectedTemplate, nil)
	return &m
}

// FetchTemplate is the mock implementation of the FetchTemplate method
func (m *MockRemoteAPI) FetchTemplate(id string) (*sql_template.SqlTemplate, error) {
	args := m.Called(id)
	return args.Get(0).(*sql_template.SqlTemplate), args.Error(1)
}
