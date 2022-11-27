package planning

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func newSchedule(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}
func newItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}
func addItemSchedule(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}
func getItem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}

func Api(port int) {
	router := httprouter.New()
	router.POST("/item", newItem)
	router.GET("/:item", getItem)

	log.Fatal(http.ListenAndServe(port, router))
}
