package endpoints

import(
	"fmt"
    "encoding/json"
    "net/http"
)

func GetUserList(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnTest")
    json.NewEncoder(w).Encode("get users list")
}


func GetUser(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnTest")
    json.NewEncoder(w).Encode("get specific user")
}

func PostUser(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnTest")
    json.NewEncoder(w).Encode("post user")
}

func PatchUser(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnTest")
    json.NewEncoder(w).Encode("patch user")
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnTest")
    json.NewEncoder(w).Encode("delete user")
}
