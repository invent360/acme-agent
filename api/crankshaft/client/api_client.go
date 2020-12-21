package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/acme-agent/api/crankshaft/models"
	"net/http"
	"os"
)

func PostObservation(observation *models.APIObservation)  {
	httpClient := &http.Client{}

	payload, err := json.Marshal(observation)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(payload))

	postReq, err := http.NewRequest("POST", "https://crankshaft.app-sec.eti-team6-demo.com/api/v1/orgs/demoat/apps/webexat/versions/5f90d62c-f71c-45db-984e-6fab58877f09/observations", bytes.NewReader(payload))

	if err != nil {
		os.Exit(1)
	}

	postReq.Header.Add("X-Access-Token", " what you talkin bout Willis?")
	resp, err := httpClient.Do(postReq)

	defer resp.Body.Close()
	fmt.Println(resp)
}