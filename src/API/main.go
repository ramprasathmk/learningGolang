package API

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Student Structure
type Student struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Mobile  string `json:"mobile"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Dob     string `json:"dob"`
}

// type StudentList defines the student list structure
type StudentList []Student

// type APIResponse defines the api response structure
type APIResponse struct {
	Status   string      `json:"status"`
	ErrMsg   string      `json:"errMsg"`
	Students StudentList `json:"students"`
}

// globally declared student list data
var data StudentList = StudentList{
	{Id: 101, Name: "manoj", Mobile: "99xxxxxx01", Email: "manoj@acb.com", Address: "Chennai", Dob: "01/01/2001"},
	{Id: 102, Name: "vimal", Mobile: "99xxxxxx02", Email: "vimal@acb.com", Address: "Kerala", Dob: "02/02/2000"},
	{Id: 103, Name: "kiran", Mobile: "99xxxxxx03", Email: "kiran@acb.com", Address: "Bangalore", Dob: "10/02/2002"},
}

// globally declared api response data
var lResp APIResponse = APIResponse{
	Status:   "S",
	ErrMsg:   "",
	Students: data,
}

// Implementation of Get Method
func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")

	if r.Method == "GET" {
		lData, err := json.Marshal(data)

		if err != nil {
			http.Error(w, "Error occured marshalling", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(lData)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// Implementation of Post method
func PostStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")

	// Check the method is POST
	if r.Method == "POST" {
		log.Println("Post Student Data")
		body, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		// instance of the type student
		var newStudent Student
		err = json.Unmarshal(body, &newStudent)

		if err != nil {
			http.Error(w, "Error occurred in unmarshalling", http.StatusInternalServerError)
			return
		}

		log.Printf("Student %v details received for addition", newStudent.Name)

		// Check if ID is unique
		isUnique := true

		for _, stu := range data {
			if stu.Id == newStudent.Id {
				isUnique = false
				break // Exit loop if duplicate ID is found
			}
		}

		if isUnique {
			data = append(data, newStudent)
			log.Printf("Student %v details added successfully\n", newStudent.Name)
		} else {
			log.Printf("Student %v already exists\n", newStudent.Name)
			http.Error(w, "Student already exists", http.StatusConflict)
			return
		}

		// Update API response
		lResp.Students = data

		lData, err := json.Marshal(lResp)
		if err != nil {
			http.Error(w, "Error marshalling response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(lData)

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// Implementation of PUT method
func PutStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")

	// Check if method is PUT
	if r.Method == "PUT" {
		log.Println("Update Student Data")
		body, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		var updatedStudent Student
		err = json.Unmarshal(body, &updatedStudent)

		if err != nil {
			http.Error(w, "Error occurred in unmarshalling", http.StatusInternalServerError)
			return
		}

		// Find and update the student
		updated := false

		for i, stu := range data {
			if stu.Id == updatedStudent.Id {
				data[i] = updatedStudent
				updated = true
				break
			}
		}

		if !updated {
			http.Error(w, "Student not found", http.StatusNotFound)
			return
		}

		// Update response data
		lResp.Students = data
		lData, err := json.Marshal(lResp)

		if err != nil {
			http.Error(w, "Error marshalling response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(lData)
		log.Printf("Student ID %d updated successfully", updatedStudent.Id)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// Implementation of Delete Method
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")

	// Check if method is DELETE
	if r.Method == "DELETE" {
		log.Println("Delete Student Data")
		body, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		// logic 1: delete the student with student id
		var studentToDelete struct {
			Id string
		}

		// logic 2: delete the student with student instance
		// instance of the type student
		var studentToDelete Student
		err = json.Unmarshal(body, &studentToDelete)

		if err != nil {
			http.Error(w, "Error occurred in unmarshalling", http.StatusInternalServerError)
			return
		}

		// Find and delete student
		index := -1

		for i, stu := range data {
			if stu.Id == studentToDelete.Id {
				index = i
				break
			}
		}

		if index == -1 {
			http.Error(w, "Student not found", http.StatusNotFound)
			return
		}

		// Remove student from the slice
		data = append(data[:index], data[index+1:]...)

		// Update response data
		lResp.Students = data
		lData, err := json.Marshal(lResp)

		if err != nil {
			http.Error(w, "Error marshalling response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(lData)
		log.Printf("Student ID %d deleted successfully", studentToDelete.Id)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	log.Println("Go Server Started...")

	// entrypoints
	// GET method (Get all Students)
	http.HandleFunc("/get", GetStudent)
	// POST method (Add student)
	http.HandleFunc("/post", PostStudent)
	// PUT method (Update student)
	http.HandleFunc("/put", PutStudent)
	// DELETE method (Delete student)
	http.HandleFunc("/delete", DeleteStudent)

	// listen to the port 5000
	http.ListenAndServe(":5000", nil)
}
