package testcases

import (
	"context"
	"testing"

	"github.com/bianhuOK/api_client/app/app_sequence_get/schemas"
	"github.com/bianhuOK/api_client/internal/domain/service"
	"github.com/bianhuOK/api_client/internal/infra/configs"
	"github.com/bianhuOK/api_client/internal/infra/persistence"
	"github.com/bianhuOK/api_client/internal/infra/repo"
	"github.com/stretchr/testify/assert"
)

func InitializeSequenceApp() (*schemas.SequenceController, error) {
	rangeConfig, err := configs.LoadRangeConfig()
	if err != nil {
		return nil, err
	}
	databaseOptionConfig, err := configs.LoadDbOptionConfig()
	if err != nil {
		return nil, err
	}
	dbProvider := persistence.NewDBProvider(databaseOptionConfig)
	sequenceConfig, err := configs.LoadSequenceConfig()
	if err != nil {
		return nil, err
	}
	sequenceRepository, err := repo.NewSequenceRepository(dbProvider, sequenceConfig)
	if err != nil {
		return nil, err
	}
	sequenceDbTransactionManager, err := repo.NewSequenceDbTransactionManager(dbProvider, sequenceConfig)
	if err != nil {
		return nil, err
	}
	rangeManager := service.NewRangeManager(sequenceRepository, sequenceDbTransactionManager, rangeConfig)
	preloadManager := service.NewPreloadManager(rangeManager, rangeConfig)
	seqGenerator := service.NewSeqGenerator(rangeConfig, rangeManager, preloadManager)
	sequenceController := schemas.NewSeqControlloer(seqGenerator)
	return sequenceController, nil
}

func TestSequenceGenerate(t *testing.T) {
	cl, err := InitializeSequenceApp()
	assert.NoError(t, err)
	id, err := cl.SequenceGenerator.NextValue(context.Background())
	assert.NoError(t, err)
	assert.NotZero(t, id)
}
