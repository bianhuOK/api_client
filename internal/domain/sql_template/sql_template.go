package sql_template

import (
	"errors"
	"time"

	"github.com/bianhuOK/api_client/internal/domain/model"
)

type TemplateCont string
type ApiId string
type ParamRule struct {
	Required bool
	DataType string
	// more
}

type SqlTemplate struct {
	ApiId           ApiId
	TemplateContent TemplateCont
	Params          []map[string][]ParamRule
	DbConfig        model.DbConfig
	Updated         time.Time
}

func (st *SqlTemplate) ValidateParams(params map[string]interface{}) bool {
	return true
}

type SqlTemplateService struct {
	TemplateRepo TemplateRepository
}

func NewSqlTemplateService(repo TemplateRepository) *SqlTemplateService {
	return &SqlTemplateService{TemplateRepo: repo}
}

func (s *SqlTemplateService) GetSqlTemplate(apiId string, params map[string]interface{}) (*SqlTemplate, error) {
	st, err := s.TemplateRepo.GetTemplateById(apiId)
	if err != nil {
		// todo
		return nil, errors.New("")
	}
	validate_res := st.ValidateParams(params)
	if !validate_res {
		// todo
		return nil, errors.New("")
	}

	// todo map variables
	return st, nil
}
