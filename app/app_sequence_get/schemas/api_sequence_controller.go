package schemas

import (
	"runtime/debug"

	"github.com/bianhuOK/api_client/internal/domain/service"
	"github.com/bianhuOK/api_client/pkg/utils"
	"github.com/go-chassis/go-chassis/v2/pkg/metrics"
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
)

type SequenceController struct {
	SequenceGenerator *service.SeqGenerator
}

func NewSeqControlloer(seqGenerator *service.SeqGenerator) *SequenceController {
	return &SequenceController{
		SequenceGenerator: seqGenerator,
	}
}

func (c *SequenceController) GetSequence(b *rf.Context) {
	logger := utils.GetLogger()
	logger.Info("GetSequence Begin")
	// 记录请求
	metrics.CounterAdd("request_counter", 1, map[string]string{
		"method":   b.ReadRequest().Method,
		"endpoint": b.ReadRequest().URL.Path,
	})
	defer func() {
		if err := recover(); err != nil {
			logger.WithFields(map[string]interface{}{
				"panic": err,
				"stack": string(debug.Stack()),
			}).Error("handle request panic")
			b.WriteJSON(struct {
				Error string `json:"error"`
			}{Error: "Internal server error"}, "application/json")
		}
	}()

	logger.Info("NextValue")
	id, err := c.SequenceGenerator.NextValue(b.Ctx)
	if err != nil {
		logger.Errorf("generate sequence err: %v", err)
		b.WriteJSON(struct {
			Error string `json:"error"`
		}{Error: "generate sequence err"}, "application/json")
		return
	}

	b.WriteJSON(struct {
		ID int64 `json:"id"`
	}{ID: id}, "application/json")
}

func (c *SequenceController) URLPatterns() []rf.Route {
	return []rf.Route{
		{Method: "POST", Path: "/sequence_generate", ResourceFunc: c.GetSequence,
			Returns: []*rf.Returns{{Code: 200}}},
	}
}
