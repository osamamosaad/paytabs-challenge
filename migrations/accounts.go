package migrations

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/osamamosaad/paytabs/entities"
	"github.com/osamamosaad/paytabs/storage"
)

func LoadAccounts() {
	url := "https://gist.githubusercontent.com/paytabs-engineering/c470210ebb19511a4e744aefc871974f/raw/6296df58428c89b8f852a6a83b0a5d0ac38289b6/accounts-mock.json"

	response, err := http.Get(url)
	if err != nil {
		log.Printf("couldn't open this url " + url)
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var accounts entities.Accounts
	json.Unmarshal(responseData, &accounts)

	// Store in DB
	storage.DBMemory = make(map[string]map[string]interface{})
	storage.DBMemory["Account"] = make(map[string]interface{})

	for _, account := range accounts {
		storage.DBMemory["Account"][account.ID] = account
	}
	log.Println("Accounts migrated and the system is ready now!")
}
