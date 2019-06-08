package service
import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"go-cat-server/dbclient"
	"fmt"
)
var DBClient dbclient.IBoltClient

func GetAccount(w http.ResponseWriter, r *http.Request) {
	// Read the 'accountId' path parameter from the mux map
	var accountId = mux.Vars(r)["accountId"]

	// Read the account struct BoltDB
	account, err := DBClient.QueryAccount(accountId)

	if err != nil {
		fmt.Println("Some error occured serving " + accountId + ": " + err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}
// If found, marshal into JSON, write headers and content
  data, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	// w.Write([]byte("{\"result\":\"OK\"}"))
}
