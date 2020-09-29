package endpoints

import(
	"fmt"
    "encoding/json"
    "net/http"
)

func ReturnTest(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnTest")
    json.NewEncoder(w).Encode("test")
}

func HomePage(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: homePage")
   
    json.NewEncoder(w).Encode("home")
}


func ReturnTestWithArgument(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnTestWithArgument")

	vars := mux.Vars(r)
    key := vars["id"]

    json.NewEncoder(w).Encode("test "+key)

}