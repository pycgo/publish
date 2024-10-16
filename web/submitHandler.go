package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"publish/internal/getTag"
	"publish/util"
	"reflect"
	"time"
)

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "POST" {
		// Read the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var formData struct {
			Options []string `json:"options"`
		}

		// Unmarshal the request body into the formData struct
		err = json.Unmarshal(body, &formData)
		if err != nil {
			http.Error(w, "Failed to unmarshal request body", http.StatusBadRequest)
			return
		}

		// Log the options received from the frontend
		fmt.Println("Received options:", formData.Options, reflect.TypeOf(formData.Options))

		now := time.Now().Format("2006-01-02 15:04:05")

		for _, option := range formData.Options {

			if option == "upgrade" {
				fmt.Println(option)
				returnMessage := util.CmdUpgrade()
				response := fmt.Sprintf("\n%v\nReceived options: %v\n%v", now, option, returnMessage)
				// Write the response back to the client
				w.Header().Set("Content-Type", "text/plain")
				w.Write([]byte(response))
			} else {
				end, _ := getTag.GetTag("jdwl-"+option, "jdwl-dev")
				returnMessage := util.Cmd(option, end)
				response := fmt.Sprintf("\n%v\nReceived options: %v\n%vcommit-id is:%v\n", now, option, returnMessage, end)
				// Write the response back to the client
				w.Header().Set("Content-Type", "text/plain")
				w.Write([]byte(response))
			}

		}

		// Process the options and prepare a response

	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
