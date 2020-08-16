package ledger

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
)

// Run is the web server
func Run() error {
	mux := makeMuxRouter()

	httpPort := os.Getenv("PORT")
	log.Println("Listening on ", os.Getenv("PORT"))
	s := &http.Server{
		Addr:           ":" + httpPort,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

// makeMuxRouter defined our handlers
func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("Get")
	muxRouter.HandleFunc("/", handleWritetBlockchain).Methods("POST")

	return muxRouter
}

// handleGetBlockchain it's a handler to get a blockchain
func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Blockchain, "", " ")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	io.WriteString(w, string(bytes))
}

// Message serves to write data
type Message struct {
	BPM int
}

// handleWritetBlockchain it's a handler to write a blockchain
func handleWritetBlockchain(w http.ResponseWriter, r *http.Request) {

	var m Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
	}

	defer r.Body.Close()

	newB, err := CreateBlock(Blockchain[len(Blockchain)-1], m.BPM)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, m)
	}

	if isValidBlock(newB, Blockchain[len(Blockchain)-1]) {
		newBlockchain := append(Blockchain, newB)
		ReplaceChain(newBlockchain)
		//pretty print our structs into the console
		spew.Dump(Blockchain)
	}

	respondWithJSON(w, r, http.StatusCreated, newB)

}

//responsdWithJson is a wrapper function which triggers status of a request
func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {

	response, err := json.MarshalIndent(payload, "", " ")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(response)
}
