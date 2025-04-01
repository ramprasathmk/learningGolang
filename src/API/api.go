package API

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Student Structure
type Employee struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	// Age     int    `json:"age"`
	Mobile  string `json:"mobile"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Dob     string `json:"dob"`
}

// type StudentList defines the student list structure
type EmployeesList []Employee

// type APIResponse defines the api response structure
type Response struct {
	Status    string        `json:"status"`
	ErrMsg    string        `json:"errMsg"`
	Employees EmployeesList `json:"empList"`
}

func ShowEmployeeData(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	log.Println("Get employee data (+)")

	if r.Method == http.MethodGet {

		_, lErr := io.ReadAll(r.Body)

		if lErr != nil {
			lResp.Status = lStatusE
			lResp.ErrMsg = lErr.Error()
			log.Println("AG001: " + lErr.Error())
		}

		var lResp Response
		lResp.Status = lStatusS
		lResp.Employees = empArr

		lData, lErr := json.Marshal(lResp)

		if lErr != nil {
			lResp.Status = lStatusE
			lResp.ErrMsg = lErr.Error()
			fmt.Fprintf(w, "AG002: %s", lErr.Error())
		} else {
			fmt.Fprintf(w, "%s", string(lData))
		}
	} else {
		log.Println("invalid request method")
	}
	log.Println("Get employee data (-)")
}

func RunServr() {
	// log start
	log.Println("Server starting...")

	// handle GET, POST, PUT and DELETE methods

	// listen to the port
	http.ListenAndServe(":29100", nil)

	// log end
	log.Println("Server stopping...")
}
