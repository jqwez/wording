package routes

import (
	"net/http"

	"github.com/jqwez/wording/finder"
	"github.com/jqwez/wording/games"
	"github.com/jqwez/wording/templates/pages"
)

type BlossomRoutes struct {
	Mux *http.ServeMux
}

func NewBlossomRoutes(mux *http.ServeMux) *BlossomRoutes {
	blr := &BlossomRoutes{Mux: mux}
	blr.ApplyRoutes()
	return blr
}

func (blr *BlossomRoutes) ApplyRoutes() {
	blr.Mux.HandleFunc("/games/blossom", blr.BasePage)
	blr.Mux.HandleFunc("/games/blossom/answers", blr.GetAnswers)
}

func (blr *BlossomRoutes) BasePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}
	pages.BlossomMainPage(false).Render(r.Context(), w)
}

func (blr *BlossomRoutes) GetAnswers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}
	err := r.ParseForm()
	if err != nil {
		pages.BlossomMainErrorPage(err.Error()).Render(r.Context(), w)
		return
	}
	center := r.FormValue("center")
	petals := r.FormValue("petals")
	blossom, err := games.NewBlossom(center, petals)
	if err != nil {
		pages.BlossomMainErrorPage(err.Error()).Render(r.Context(), w)
		return
	}
	dictionary, err := finder.NewDictionary()
	if err != nil {
		pages.BlossomMainErrorPage(err.Error()).Render(r.Context(), w)
		return
	}
	words := blossom.FindWords(dictionary)
	data := blossom.WordsWithInfo(words)

	pages.BlossomAnswersPage(data).Render(r.Context(), w)
}
