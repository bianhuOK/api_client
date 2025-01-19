package testcases

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/bianhuOK/api_client/internal/infra/configs"
	"github.com/bianhuOK/api_client/internal/infra/persistence"
	"github.com/bianhuOK/api_client/internal/infra/repo"
	"github.com/stretchr/testify/assert"
)

func TestSeqRepoSelect(t *testing.T) {
	os.Setenv("SEQUENCE_CONFIG_PATH", "/Users/bianniu/GolandProjects/api_client/internal/infra/configs/sequence.local.yaml")
	seqConfig, err := configs.LoadSequenceConfig()
	assert.NoError(t, err)
	provider := persistence.NewDBProvider(&seqConfig.DatabaseOptionConfig)
	seqRepo, err := repo.NewSequenceRepository(provider, seqConfig)
	seq, err := seqRepo.GetSequenceForUpdate(context.Background(), "test_seq")
	assert.NoError(t, err)
	data, err := json.Marshal(seq)
	assert.NoError(t, err)
	fmt.Println(string(data))

	// update
	// seq.CurrentValue = 0
	// err = seqRepo.UpdateSequence(context.Background(), seq)
	// assert.NoError(t, err)
}

func TestTransaction(t *testing.T) {
	os.Setenv("SEQUENCE_CONFIG_PATH", "/Users/bianniu/GolandProjects/api_client/internal/infra/configs/sequence.local.yaml")
	seqConfig, err := configs.LoadSequenceConfig()
	assert.NoError(t, err)
	provider := persistence.NewDBProvider(&seqConfig.DatabaseOptionConfig)
	seqRepo, err := repo.NewSequenceRepository(provider, seqConfig)
	assert.NoError(t, err)
	seqTxManager, err := repo.NewSequenceDbTransactionManager(provider, seqConfig)
	assert.NoError(t, err)
	err = seqTxManager.ExecTx(context.Background(), func(ctx context.Context) error {
		seq, err := seqRepo.GetSequenceForUpdate(context.Background(), "test_seq")
		if err != nil {
			return err
		}
		data, err := json.Marshal(seq)
		if err != nil {
			return err
		}
		fmt.Println(string(data))
		return nil
	})
	assert.NoError(t, err)
}
