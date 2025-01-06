package repo

import (
	"fmt"

	"github.com/bianhuOK/api_client/internal/domain/model"
	repo "github.com/bianhuOK/api_client/internal/infra/iface"
	"github.com/bianhuOK/api_client/internal/infra/persistence"
	"github.com/bianhuOK/api_client/pkg/utils"
)

type DbFactoryImpl struct {
}

func NewDbFactoryImpl() *DbFactoryImpl {
	return &DbFactoryImpl{}
}

func (d *DbFactoryImpl) GetDbRepository(dc model.DbConfig) (repo.DbRepository, error) {
	log := utils.GetLogger()
	switch dc.DbType {
	case "mysql":
		dbRepo, err := persistence.NewMySQLDBRepository(dc.DSN)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return dbRepo, nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dc.DbType)
	}
}
