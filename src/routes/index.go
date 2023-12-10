package routes

import (
	"fmt"
	"net/http"

	"github.com/sapphirenw/gotmpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if err := gotmpl.XT.ExecuteTemplate(w, "index.html", nil); err != nil {
		fmt.Println(err)
		Error(w, r)
	}
}
