package API

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Response struct represents the Response structure
type Response struct {
	Status      string
	ErrMsg      string
	UserArrList []UserStruct
}

// UserStruct struct represents the UserStruct structure
type UserStruct struct {
	Username string `json: "username"`
	Age      int `json: "age"`
}

// GetUserList() writes the user list
func GetUserList(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")

	var lResp Response
	var arr []UserStruct

	lResp.Status = "S"
	for i:=0; i<5; i++ {
		arr=append(arr, UserStruct{Username: "User_"+fmt.Sprintf("%d", i), Age: 28})
	}

	lResp.UserArrList = arr
	lData, _ := json.Marshal(lResp)
	fmt.Fprintf(w, "%s", string(lData))
}

// Greet() greets with time
func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Gopher!\n Time: %s", time.Now())
}

func GetUserData(w http.ResponseWriter, r *http.Request) {
	var UserDetails UserStruct

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	fmt.Println("Raw JSON body: ", string(body))

	err = json.Unmarshal(body, &UserDetails)

	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusInternalServerError)
		return 
	}

	fmt.Println("user details: ", UserDetails)

	status := "S"

	response := Response{
		Status: status,
		UserArrList: []UserStruct{UserDetails},
	}

	resp, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

// func main() {
// 	fmt.Println("Starting main() function")
// 	log.Println("Server Started...")
// 	// entrypoints
// 	http.HandleFunc("/hi", Greet)
// 	http.HandleFunc("/getuserlist", GetUserList)
//  http.HandleFunc("/getuserdata", GetUserData)
// 	// port
// 	http.ListenAndServe(":29100", nil)
// 	log.Println("Server Ended...")
// }