package httpdelivery

import (
	"fmt"
	"net/http"
)

func (u *Handler) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := u.TemplateCache[name]
	if !ok {
		http.Error(w, "no template", http.StatusInternalServerError)
		return
	}

	err := ts.Execute(w, td)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}

}

func (u *Handler) renderHTML(w http.ResponseWriter, r *http.Request, code int, name string, obj interface{}) {
	ts, ok := u.TemplateCache[name]
	if !ok {
		http.Error(w, "no template", http.StatusInternalServerError)
		return
	}

	// w.WriteHeader(code)

	err := ts.Execute(w, obj)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}

}
