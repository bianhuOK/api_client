package service

import (
	"github.com/bianhuOK/api_client/internal/domain/model"
	repo "github.com/bianhuOK/api_client/internal/infra/iface"
	"github.com/bianhuOK/api_client/pkg/utils"
)

type SqlQueryService struct {
	DbFactory repo.DbFactory
}

func NewSqlQueryService(dbFactory repo.DbFactory) *SqlQueryService {
	return &SqlQueryService{
		DbFactory: dbFactory,
	}
}

func (s *SqlQueryService) ExecuteSql(dc model.DbConfig, sql string) (results *model.SqlQuery, err error) {
	log := utils.GetLogger()
	db, err := s.DbFactory.GetDbRepository(dc)
	if err != nil {
		log.Warn("GetDbRepository error %w", err)
		return nil, err
	}

	rows, err := db.ExecuteSql(sql)
	if err != nil {
		log.Warn("ExecuteSql error %w", err)
		return nil, err
	}
	results = &model.SqlQuery{
		Query:    sql,
		DbConfig: dc,
		Results:  rows,
	}
	return
}
