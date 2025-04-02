package API

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
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

func ShowEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	log.Println("Get employee data (+)")

	if r.Method == http.MethodGet {

		_, lErr := io.ReadAll(r.Body)

		var lResp Response
		lResp.Status = lStatusS
		lResp.Employees = lEmpArr

		if lErr != nil {
			lResp.Status = lStatusE
			lResp.ErrMsg = lErr.Error()
			log.Println("AG001: " + lErr.Error())
		}

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

func InsertEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	log.Println("Insert Employee data (+)")

	if r.Method == http.MethodPost {
		var lResp Response
		lResp.Status = lStatusS

		lBody, lErr := io.ReadAll(r.Body)

		if lErr != nil {
			lResp.Status = lStatusE
			lResp.ErrMsg = lErr.Error()
			log.Println("AI001: " + lErr.Error())
		}

		var lNewEmp Employee

		lErr = json.Unmarshal(lBody, &lNewEmp)

		if lErr != nil {
			lResp.Status = lStatusE
			lResp.ErrMsg = lErr.Error()
			log.Println("AI002: " + lErr.Error())
		}

		lIsUnique := true

		for _, lEmp := range lEmpArr {
			if lEmp.Id == lNewEmp.Id {
				lIsUnique = false
				break
			}
		}

		if !lIsUnique {
			lResp.Status = lStatusE
			lResp.ErrMsg = lEmpExistErr
			log.Println("AI003: " + lEmpExistErr)
			return
		}

		lEmpArr = append(lEmpArr, lNewEmp)

		// lResp.EmployeeList = empArr

		lData, lErr := json.Marshal(lResp)

		if lErr != nil {
			lResp.Status = lStatusE
			lResp.ErrMsg = lErr.Error()
			fmt.Fprintf(w, "%s", lErr.Error())
		} else {
			log.Println("employee data inserted successfully")
			fmt.Fprintf(w, "%s", string(lData))
		}
	} else {
		log.Println("invalid request method")
	}
	log.Println("Insert Employee Data (-)")
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	log.Println("Update Employee Data (+)")

	switch r.Method {
	default:
		{
			log.Println("invalid request method")
		}
	case http.MethodPut:
		{
			var lResp Response
			lResp.Status = lStatusS

			lBody, lErr := io.ReadAll(r.Body)

			if lErr != nil {
				lResp.Status = lStatusE
				lResp.ErrMsg = lErr.Error()
				log.Println("AU001: " + lErr.Error())
			}

			var lUpdateEmp Employee

			lErr = json.Unmarshal(lBody, &lUpdateEmp)

			if lErr != nil {
				lResp.Status = lStatusE
				lResp.ErrMsg = lErr.Error()
				log.Println("AU002: " + lErr.Error())
			}

			lIsUpdated := false

			for lIndex, lEmp := range lEmpArr {
				if lEmp.Id == lUpdateEmp.Id {
					lEmpArr[lIndex] = lUpdateEmp
					lIsUpdated = true
					break
				}
			}

			if !lIsUpdated {
				lResp.Status = lStatusE
				lResp.ErrMsg = lEmpNotFoundErr
				log.Println("AU003: " + lEmpNotFoundErr)
				return
			}

			// lResp.Employees = lEmpArr

			log.Println("employee data updated successfully")

			lData, lErr := json.Marshal(lResp)

			if lErr != nil {
				lResp.Status = lStatusE
				lResp.ErrMsg = lErr.Error()
				fmt.Fprintf(w, "%s", lErr.Error())
			} else {
				fmt.Fprintf(w, "%s", string(lData))
			}
		}
	}
	log.Println("Update Employee Data (-)")
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	log.Println("Delete Employee Data (+)")

	if r.Method == http.MethodDelete {
		var lResp Response
		lResp.Status = lStatusS

		lBody, lErr := io.ReadAll(r.Body)

		if lErr != nil {
			lResp.Status = lStatusS
			lResp.ErrMsg = lErr.Error()
			log.Println("AD001: " + lErr.Error())
		}

		var lDeleteEmp struct {
			Id int `json:"id"`
		}

		lErr = json.Unmarshal(lBody, &lDeleteEmp)

		if lErr != nil {
			lResp.Status = lStatusE
			lResp.ErrMsg = lErr.Error()
			log.Println("AD002: " + lErr.Error())
		}

		lIndex := -1

		for lEmpIndex, lEmp := range lEmpArr {
			if lEmp.Id == lDeleteEmp.Id {
				lIndex = lEmpIndex
				break
			}
		}

		if lIndex == -1 {
			lResp.Status = lStatusE
			lResp.ErrMsg = lEmpNotFoundErr
			log.Println("AD003: " + lEmpNotFoundErr)
			return
		}

		// lEmpArr = slices.Delete(lEmpArr, lIndex, lIndex+1) // using built-in
		lEmpArr = append(lEmpArr[:lIndex], lEmpArr[lIndex+1:]...) // using logic

		log.Println("employee data deleted successfully")

		lData, lErr := json.Marshal(lResp)

		if lErr != nil {
			lResp.Status = lStatusE
			lResp.ErrMsg = lErr.Error()
			fmt.Fprintf(w, "%s", lErr.Error())
		} else {
			fmt.Fprintf(w, "%s", string(lData))
		}
	} else {
		log.Println("invalid request method")
	}
	log.Println("Delete Employee Data (-)")
}

func RunAPIServer() {
	// log start
	log.Println("Server starting...")

	// ! create "log/" dir before running the file.
	lLogFile, lErr := os.OpenFile("./log/logfile"+time.Now().Format("02012006.15.04.05.000000000")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if lErr != nil {
		log.Fatalf("error opening logfile: %v", lErr.Error())
	}

	defer func() {
		log.Println("log file closed successfully")
		lLogFile.Close()
	}()

	log.SetOutput(lLogFile)

	// handle GET, POST, PUT and DELETE methods
	http.HandleFunc("/get", ShowEmployee)
	http.HandleFunc("/post", InsertEmployee)
	http.HandleFunc("/update", UpdateEmployee)
	http.HandleFunc("/delete", DeleteEmployee)

	// listen to the port
	http.ListenAndServe(":29100", nil)

	// log end
	log.Println("Server stopping...")
}
