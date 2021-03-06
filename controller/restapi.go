package controller

import (
	"context"
	"encoding/json"
	"github.com/sycki/mknote/logger"
	"net/http"
)

func (m *Manager) ArticleNavigation(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == get {
		articleIndex, err := m.storage.GetTags()
		if err != nil {
			panic(err)
		}
		result, _ := json.Marshal(articleIndex)
		w.Write(result)
	}
}

func (m *Manager) Articles(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	uri := r.URL.Path[len("/api/v1/articles/"):]
	if method == get {
		articleTag, err := m.storage.GetArticle(uri)
		if err != nil {
			panic(err)
		}
		result, _ := json.Marshal(articleTag)
		w.Write(result)
	}
}

func (m *Manager) Like(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	artID := r.URL.Path[len("/v1/like/"):]

	if method == post {
		m.l.Lock()
		defer m.l.Unlock()
		art, err := m.storage.GetArticle(artID)
		if err != nil {
			panic(err)
		}
		art.Like_count += 1
		if _, err2 := m.storage.UpdateArtcile(art); err != nil {
			panic(err2)
		}
	}
}

func (m *Manager) Visit(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	artID := r.URL.Path[len("/v1/visit/articles/"):]

	if method == post {
		m.l.Lock()
		defer m.l.Unlock()
		art, err := m.storage.GetArticle(artID)
		if err != nil {
			panic(err)
		}
		art.Viewer_count += 1
		if _, err2 := m.storage.UpdateArtcile(art); err != nil {
			panic(err2)
		}
	}
}

func (m *Manager) Pprof(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	action := r.URL.Path[len("/v1/manage/pprof/"):]

	if method == post {
		if action == "open" {
			if m.pprof == nil {
				m.pprof = &http.Server{Addr: ":" + m.config.DebugPort, Handler: nil, ErrorLog: logger.GetLogger()}
				go func() {
					m.pprof.ListenAndServe()
				}()
				logger.Info("pprofile server is started on: ", m.config.DebugPort)
				w.Write([]byte("open"))
			}
		} else if action == "close" {
			if m.pprof != nil {
				m.pprof.Shutdown(context.Background())
				m.pprof = nil
				logger.Info("pprofile server is stopped gracefully")
				w.Write([]byte("close"))
			}
		}
	}
}
