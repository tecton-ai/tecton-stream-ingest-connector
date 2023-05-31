package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type IngestAPInput struct {
	WorkspaceName string                   `json:"workspace_name"`
	DryRun        bool                     `json:"dry_run"`
	Records       map[string][]interface{} `json:"records"`
}
type RecordWrapper struct {
	Record interface{} `json:"record"`
}

var (
	apiInput       *IngestAPInput
	httpClient     *http.Client
	pushSourceName string
	apiURL         string
	authToken      string
)

func handler(ctx context.Context, kinesisEvent events.KinesisEvent) error {
	recordsInput := make(map[string][]interface{})
	records := make([]interface{}, 0)
	for _, record := range kinesisEvent.Records {
		dataBytes := record.Kinesis.Data
		out := map[string]interface{}{}
		json.Unmarshal(dataBytes, &out)
		recordWrapper := &RecordWrapper{
			Record: out,
		}
		records = append(records, recordWrapper)
	}
	recordsInput[pushSourceName] = records
	apiInput.Records = recordsInput
	apiInputJSON, _ := json.Marshal(apiInput)
	fmt.Println(string(apiInputJSON))
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(apiInputJSON))
	req.Header.Set("Authorization", authToken)
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Response Status: %v", resp.StatusCode))
	}
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))
	return nil
}

func main() {
	apiInput = &IngestAPInput{
		WorkspaceName: os.Getenv("WORKSPACE_NAME"),
		DryRun:        false,
	}
	httpClient = &http.Client{}
	pushSourceName = os.Getenv("PUSH_SOURCE_NAME")
	apiURL = os.Getenv("TECTON_API_URL")
	authToken = os.Getenv("TECTON_AUTH_TOKEN")
	lambda.Start(handler)
}
