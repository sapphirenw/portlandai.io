package routes

import (
	"fmt"
	"net/http"

	"github.com/sapphirenw/gotmpl"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := gotmpl.XT.ExecuteTemplate(w, "notFound.html", nil); err != nil {
		fmt.Println(err)
		http.Error(w, "<html><p>There was a critical error</p></html>", http.StatusInternalServerError)
	}
}
