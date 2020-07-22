package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var elements map[string]int

func main() {
	handler := createHandler()
	elements = make(map[string]int)
	http.ListenAndServe(":8080", handler)
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
			Message: err.Error(),
			Code:    400,
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
				Message: "The result is " + strconv.Itoa(add),
				Code:    200,
			}
		} else {
			resp = APIResponse{
				Message: "element doesnt exist ",
				Code:    200,
			}
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
			Message: err.Error(),
			Code:    400,
		}
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
				Message: "The result is " + strconv.Itoa(add),
				Code:    200,
			}
			w.WriteHeader(http.StatusOK)
		} else {
			resp = APIResponse{
				Message: "element doesnt exist ",
				Code:    404,
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
			Message: err.Error(),
			Code:    400,
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		for elem := range temp {
			elements[elem] = temp[elem]
		}
		resp = APIResponse{
			Message: "Data Storage",
			Code:    200,
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
			Message: "Element " + element + " is " + strconv.Itoa(value),
			Code:    200,
		}
	} else {
		resp = APIResponse{
			Message: "Element " + element + " dont exist ",
			Code:    404,
		}
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
			Message: err.Error(),
			Code:    400,
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		if exist {
			elements[element] = value
			resp = APIResponse{
				Message: "Element " + element + " is " + strconv.Itoa(value),
				Code:    200,
			}
		} else {
			resp = APIResponse{
				Message: "Element " + element + " dont exist ",
				Code:    404,
			}
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
			Message: "Element " + element + " deleted ",
			Code:    200,
		}
	} else {
		resp = APIResponse{
			Message: "Element " + element + " dont exist ",
			Code:    404,
		}
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(resp)
}

//APIResponse response
type APIResponse struct {
	Message string
	Code    int
}
