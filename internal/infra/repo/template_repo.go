package repo

import (
	"sync"
	"time"

	domain "github.com/bianhuOK/api_client/internal/domain/sql_template"
)

type RemoteApiTemplateRepository struct {
	TemplateCache     domain.Cache
	TemplateRemoteApi domain.RemoteAPI
	mux               sync.Mutex
}

func NewRemoteApiTemplateRepository(c domain.Cache, remoteAPi domain.RemoteAPI) *RemoteApiTemplateRepository {
	return &RemoteApiTemplateRepository{
		TemplateCache:     c,
		TemplateRemoteApi: remoteAPi,
		mux:               sync.Mutex{},
	}
}

func (r *RemoteApiTemplateRepository) GetTemplateById(id string) (*domain.SqlTemplate, error) {
	// 检查缓存
	template, ok := r.TemplateCache.Get(id)
	if ok {
		return template, nil
	}

	// 加锁避免重复拉取
	r.mux.Lock()
	defer r.mux.Unlock()

	// 再次检查缓存
	template, ok = r.TemplateCache.Get(id)
	if ok {
		return template, nil
	}

	// 从远程API拉取模板
	fetchedTemplate, err := r.TemplateRemoteApi.FetchTemplate(id)
	if err != nil {
		return nil, err
	}

	// 缓存并设置TTL
	r.TemplateCache.Set(id, fetchedTemplate, 5*time.Minute)
	return fetchedTemplate, nil
}
