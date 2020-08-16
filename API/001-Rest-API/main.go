package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var elements map[string]int
var client string

func main() {
	handler := createHandler()
	elements = make(map[string]int)
	go http.ListenAndServe(":8080", handler)
	log.Println("Server Ready in 8080")
	for {
		time.Sleep(500)
	}

}

func createHandler() *mux.Router {
	handler := mux.NewRouter()
	//ENDPOINT
	handler.HandleFunc("/add", apiAdd).Methods(http.MethodGet)
	handler.HandleFunc("/sub", apiSub).Methods(http.MethodGet)

	handler.HandleFunc("/element", apiSetElement).Methods(http.MethodPost)
	handler.HandleFunc("/element/{name}", apiGetElement).Methods(http.MethodGet)
	handler.HandleFunc("/element/{name}", apiUpdateElement).Methods(http.MethodPut)
	handler.HandleFunc("/element/{name}", apiDeleteElement).Methods(http.MethodDelete)
	return handler
}

//FUNCTION FOR THE API
func apiAdd(w http.ResponseWriter, r *http.Request) {
	resp := APIResponse{}
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	temp := make(map[string]int)
	err := json.Unmarshal(reqBody, &temp)
	if err != nil {
		resp = APIResponse{
			createHeader(),
			BodyResponse{
				Message: err.Error(),
			},
			//Code:    400,
		}
	} else {
		present := false
		add := 0
		for elem := range temp {
			val, present := elements[elem]
			if !present {
				break
			}
			add += val

		}
		if present {
			resp = APIResponse{
				createHeader(),
				BodyResponse{
					Message: "The result is " + strconv.Itoa(add),
				},
			}
			w.WriteHeader(http.StatusOK)
		} else {
			resp = APIResponse{
				createHeader(),
				BodyResponse{
					Message: "element doesnt exist ",
				},
			}
			w.WriteHeader(http.StatusNotFound)
		}

	}
	json.NewEncoder(w).Encode(resp)
}
func apiSub(w http.ResponseWriter, r *http.Request) {
	resp := APIResponse{}
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	temp := make(map[string]int)
	err := json.Unmarshal(reqBody, &temp)
	if err != nil {
		resp = APIResponse{
			createHeader(),
			BodyResponse{
				Message: err.Error(),
			},
		}
		w.WriteHeader(http.StatusNotFound)
	} else {
		present := false
		add := 0
		for elem := range temp {
			val, present := elements[elem]
			if !present {
				break
			}
			add -= val

		}
		if present {
			resp = APIResponse{
				createHeader(),
				BodyResponse{
					Message: "The result is " + strconv.Itoa(add),
				},
			}
			w.WriteHeader(http.StatusOK)
		} else {

			resp = APIResponse{
				createHeader(),
				BodyResponse{
					Message: "element doesnt exist ",
				},
			}
			w.WriteHeader(http.StatusNotFound)

		}

	}
	json.NewEncoder(w).Encode(resp)
}
func apiSetElement(w http.ResponseWriter, r *http.Request) {
	resp := APIResponse{}
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	temp := make(map[string]int)
	err := json.Unmarshal(reqBody, &temp)
	if err != nil {

		resp = APIResponse{
			createHeader(),
			BodyResponse{
				Message: err.Error(),
			},
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		for elem := range temp {
			elements[elem] = temp[elem]
		}

		resp = APIResponse{
			createHeader(),
			BodyResponse{
				Message: "Data Storage",
			},
		}
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(resp)
}
func apiGetElement(w http.ResponseWriter, r *http.Request) {
	resp := APIResponse{}
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	element := params["name"]
	value, exist := elements[element]
	if exist {

		resp = APIResponse{
			createHeader(),
			BodyResponse{
				Message: "Element " + element + " is " + strconv.Itoa(value),
			},
		}
		w.WriteHeader(http.StatusOK)
	} else {

		resp = APIResponse{
			createHeader(),
			BodyResponse{
				Message: "Element " + element + " dont exist ",
			},
		}
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(resp)

}
func apiUpdateElement(w http.ResponseWriter, r *http.Request) {
	resp := APIResponse{}
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	element := params["name"]
	value, exist := elements[element]
	reqBody, _ := ioutil.ReadAll(r.Body)
	temp := make(map[string]int)
	err := json.Unmarshal(reqBody, &temp)
	if err != nil {

		resp = APIResponse{
			createHeader(),
			BodyResponse{
				Message: err.Error(),
			},
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		if exist {
			elements[element] = value

			resp = APIResponse{
				createHeader(),
				BodyResponse{
					Message: "Element " + element + " is " + strconv.Itoa(value),
				},
			}
			w.WriteHeader(http.StatusOK)
		} else {

			resp = APIResponse{
				createHeader(),
				BodyResponse{
					Message: "Element " + element + " dont exist ",
				},
			}
			w.WriteHeader(http.StatusNotFound)
		}
	}

	json.NewEncoder(w).Encode(resp)

}
func apiDeleteElement(w http.ResponseWriter, r *http.Request) {
	resp := APIResponse{}
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	element := params["name"]
	_, exist := elements[element]
	if exist {
		delete(elements, element)
		w.WriteHeader(http.StatusOK)

		resp = APIResponse{
			createHeader(),
			BodyResponse{
				Message: "Element " + element + " deleted ",
			},
		}
		w.WriteHeader(http.StatusOK)
	} else {

		resp = APIResponse{
			createHeader(),
			BodyResponse{
				Message: "Element " + element + " dont exist ",
			},
		}
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(resp)
}

func createHeader() HeaderResponse {
	return HeaderResponse{
		Client: client,
		Time:   time.Now().UTC().String(),
	}
}

//APIResponse response
type APIResponse struct {
	Header HeaderResponse
	BodyResponse
}

//HeaderResponse header of the response
type HeaderResponse struct {
	Client string
	Time   string
}

//BodyResponse body of the response
type BodyResponse struct {
	Message string
}
