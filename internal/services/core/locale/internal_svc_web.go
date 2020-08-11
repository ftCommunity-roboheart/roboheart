package locale

import (
	"errors"
	"github.com/ftCommunity/roboheart/internal/service"
	"net/http"

	"github.com/ftCommunity/roboheart/internal/services/core/web"
	"github.com/ftCommunity/roboheart/package/api"
)

func (l *locale) initSvcWeb(svc service.Service) {
	var ok bool
	l.web, ok = svc.(web.Web)
	if !ok {
		l.error(errors.New("Type assertion error"))
	}
	l.mux = l.web.RegisterServiceAPI(l)
	l.mux.HandleFunc("/locale", func(w http.ResponseWriter, r *http.Request) {
		if locale, err := l.GetLocale(); err != nil {
			api.ErrorResponseWriter(w, 500, err)
		} else {
			api.ResponseWriter(w, locale)
		}
	}).Methods("GET")
	l.mux.HandleFunc("/locale", func(w http.ResponseWriter, r *http.Request) {
		data := &struct {
			api.TokenRequest
			Locale string `json:"locale"`
		}{}
		if !api.RequestLoader(r, w, data) {
			return
		}
		if err, uae := l.SetLocale(data.Token, data.Locale); err != nil {
			code := 500
			if uae {
				code = 403
			}
			api.ErrorResponseWriter(w, code, err)
		} else {
			api.ResponseWriter(w, nil)
		}
	}).Methods("POST")
	l.mux.HandleFunc("/allowed", func(w http.ResponseWriter, r *http.Request) {
		api.ResponseWriter(w, l.GetAllowedLocales())
	}).Methods("GET")
}
