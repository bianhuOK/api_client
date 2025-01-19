//go:build wireinject
// +build wireinject

package schemas

import (
	"github.com/bianhuOK/api_client/internal/domain"
	"github.com/bianhuOK/api_client/internal/infra"
	"github.com/google/wire"
)

var SequenceAppSet = wire.NewSet(
	domain.SequenceServiceSet,
	infra.SequenceGenerateInfraSet,
	NewSeqControlloer,
)

func InitializeSequenceApp() (*SequenceController, error) {
	wire.Build(SequenceAppSet)
	return &SequenceController{}, nil
}
