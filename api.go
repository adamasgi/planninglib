package planning

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Sys) newSchedule(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
}
func (s *Sys) newItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
}
func (s *Sys) addItemSchedule(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
}
func (s *Sys) getItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := p.ByName("id")
	if id == "" {
		errMesg := struct{ Err string }{Err: "ID not set"}
		errJson := json.NewEncoder(w).Encode(errMesg)
		fmt.Fprint(w, errJson)

	} else {
		var res Item
		s.Db.First(&res, "id = ?", id)
		fmt.Printf("%#v", res)
		j := json.NewEncoder(w).Encode(res)
		fmt.Fprint(w, j)
	}
}

func (s *Sys) handler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "LONNIE!!! %s", s.Filename)
}

func Api(port string) {
	sys := NewSys("./test.db")
	i := sys.AddItem("TESTING")
	fmt.Printf("ID: %s", i.Id)
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

	sys.Router.GET("/item/:item", sys.getItem)
	sys.Router.GET("/", sys.handler)

	log.Fatal(http.ListenAndServe(":"+port, sys.Router))
}
