package planning

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func newSchedule(s *Sys, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
}
func newItem(s *Sys, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
}
func addItemSchedule(s *Sys, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
}
func getItem(s *Sys, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := p.ByName("id")
	var res Item
	s.Db.First(&res, "id = ?", id)
	j := json.NewEncoder(w).Encode(res)
	fmt.Fprint(w, j)
}

func (s *Sys) handler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "LONNIE!!! %s", s.Filename)
}

func Api(port string) {
	sys := NewSys("./test.db")
	sys.Router = httprouter.New()
	sys.Router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	sys.Router.POST("/item", newItem)
	sys.Router.GET("/:item", getItem)

	log.Fatal(http.ListenAndServe(":"+port, sys.Router))
}
