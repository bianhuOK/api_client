package remoteapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	sqltemplate "github.com/bianhuOK/api_client/internal/domain/sql_template"
)

type SqlApiRest struct {
	BaseUrl string
}

func NewSqlApiRest(baseUrl string) *SqlApiRest {
	return &SqlApiRest{
		BaseUrl: baseUrl,
	}
}

func (s *SqlApiRest) FetchTemplate(id string) (*sqltemplate.SqlTemplate, error) {
	url := fmt.Sprintf("%s/templates/%s", s.BaseUrl, id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch template: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var template sqltemplate.SqlTemplate
	if err := json.Unmarshal(body, &template); err != nil {
		return nil, fmt.Errorf("failed to unmarshal template: %w", err)
	}

	return &template, nil
}
