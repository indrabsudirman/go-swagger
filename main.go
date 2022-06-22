package main

import (
	"encoding/json"
	"fmt"
	"go-swagger/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// User represents the user for this application
// swagger:model
type User struct {
	// the name for this user
	// required: true
	// min length: 5
	Name string `json:"name"`

	// The birth year for the user
	// min: 1900
	// max: 2022
	BirthYear int `json:"birth_year"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// swagger:operation GET /users getUsers
		//
		// Insert documentation
		//
		// ---
		// produces:
		// - application/json
		// responses:
		//   '200':
		//     description: user response
		//     schema:
		//       type: array
		//       items:
		//         "$ref": "#/definitions/User"

		users := []model.User{
			{"Indra", 1984},
			{"Lubna", 2016},
		}

		result, _ := json.Marshal(&users)

		_, _ = w.Write(result)
	}).Methods(http.MethodGet)

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// XXX: Imagine validation is implemented here

		// swagger:operation POST /users postUser
		//
		// Include documentation
		//
		// ---
		// produces:
		// - application/json
		// parameters:
		//   - name: Body
		//     in: body
		//     schema:
		//       "$ref": "#/definitions/User"
		// responses:
		//   '200':
		//     description: user response

		var user model.User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			_, _ = w.Write([]byte("decoding failed"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write([]byte(fmt.Sprintf("created %+v", user)))
	}).Methods(http.MethodPost)

	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(s.ListenAndServe())

}
