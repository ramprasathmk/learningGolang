package API

import (
	"encoding/json"
	"fmt"
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
	Username string
	Age      int
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
