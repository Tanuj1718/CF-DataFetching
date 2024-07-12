package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Problem struct {
	Name string `json:"name"`
}

type Member struct {
	Handle string `json:"handle"`
}

type Author struct {
	Members []Member `json:"members"`
}

type Cheater struct {
	CandidateId   int     `json:"id"`
	CandidateName Author  `json:"author"`
	ContestId     int     `json:"contestId"`
	ProblemName   Problem `json:"problem"`
	Verdict       string  `json:"verdict"`
}

type Apiresponse struct {
	Status string
	Result []Cheater
}

func main() {
	fmt.Println("CODEFORCES CONTEST STATUS...")

	r := mux.NewRouter()

	//router
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/data", data).Methods("GET")

	//CORS MIDDLEWARE
	r.Use(mux.CORSMethodMiddleware(r))

	//listen to the port
	log.Fatal(http.ListenAndServe(":3000", r))


}

func getresponse(id string) (string, error) {
	myurl := fmt.Sprintf("https://codeforces.com/api/user.status?handle=%s&from=1&count=16", id)

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	var apiresponse Apiresponse
	err = json.Unmarshal(data, &apiresponse)
	if err != nil {
		panic(err)
	}
	output := fmt.Sprintf("Candidate Name: %s\n", apiresponse.Result[1].CandidateName.Members[0].Handle)
	

	for _, it := range apiresponse.Result {
		if it.Verdict == "SKIPPED" {
			output += "Candidate is a cheater"
			break
		}

		output += fmt.Sprintf("\nCandidateID: %d\n ContestID: %d\n Problem: %s\n Verdict: %s\n", it.CandidateId, it.ContestId, it.ProblemName.Name, it.Verdict)
	}
	return output, nil

}

func data(w http.ResponseWriter, r *http.Request){
	//CORS HEADER
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	id := r.URL.Query().Get("id")
	if (id==""){
		http.Error(w, "CF id is required", http.StatusBadRequest)
		return
	}
	output, err := getresponse(id)
	if(err!=nil){
		panic(err)
	}
	fmt.Fprint(w, output)
}

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to Codeforces Problem Stats...</h1>"))
}

