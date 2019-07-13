package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"github.com/smartcontractkit/bea-adapter/services"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Result struct {
	DataValue  string `json:"DataValue"`
	SeriesCode string `json:"SeriesCode"`
	TimePeriod string `json:"TimePeriod"`
}

type Request struct {
	JobId string `json:"id"`
}

type Response struct {
	JobRunID   string                 `json:"jobRunId"`
	StatusCode int                    `json:"statusCode"`
	Status     string                 `json:"status"`
	Data       map[string]interface{} `json:"data"`
	Error      interface{}            `json:"error"`
	Result     interface{}            `json:"result"`
}

func main() {
	if os.Getenv("API_KEY") == "" {
		fmt.Println("Error: missing API key")
		return
	}

	StartServer()
}

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		decoder := json.NewDecoder(r.Body)
		var t Request
		err := decoder.Decode(&t)
		if err != nil {
			writeError(w, t.JobId, err)
			return
		}

		avg, err := GetDpcergAvg(3)
		if err != nil {
			writeError(w, t.JobId, err)
			return
		}

		avgRounded, err := strconv.ParseFloat(fmt.Sprintf("%.3f", avg), 64)
		if err != nil {
			writeError(w, t.JobId, err)
			return
		}

		writeResult(w, t.JobId, avgRounded)
	})

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func writeResult(w http.ResponseWriter, jobRunId string, result interface{}) {
	encoder := json.NewEncoder(w)
	_ = encoder.Encode(Response{
		JobRunID:   jobRunId,
		StatusCode: 200,
		Status:     "success",
		Data: map[string]interface{}{
			"result": result,
		},
		Result: result,
	})
}

func writeError(w http.ResponseWriter, jobRunId string, err interface{}) {
	encoder := json.NewEncoder(w)
	_ = encoder.Encode(Response{
		JobRunID:   jobRunId,
		StatusCode: 500,
		Status:     "error",
		Error:      err,
	})
}

func getDpcergClient() (*services.Client, error) {
	client, err := services.NewClient()
	if err != nil {
		return nil, err
	}

	client.AddParam("DataSetName", "NIPA")
	client.AddParam("TableName", "T20804")
	client.AddParam("UserID", os.Getenv("API_KEY"))
	client.AddParam("ResultFormat", "json")
	client.AddParam("method", "getData")
	client.AddParam("Frequency", "M")
	client.AddParam("Year", "2019,2018")

	return client, nil
}

func getDpcergData() ([]interface{}, error) {
	client, err := getDpcergClient()
	if err != nil {
		return nil, err
	}

	return client.GetData()
}

func makePriorityQueue(items []map[string]interface{}) services.PriorityQueue {
	priorityQueue := services.PriorityQueue{}

	for _, item := range items {
		val, err := strconv.ParseFloat(item["DataValue"].(string), 64)
		if err != nil {
			continue
		}

		date := strings.Split(item["TimePeriod"].(string), "M")
		year, err := strconv.Atoi(date[0])
		if err != nil {
			continue
		}
		month, err := strconv.Atoi(date[1])
		if err != nil {
			continue
		}

		priorityQueue.Push(&services.Item{Value: val, Year: year, Month: month})
	}

	heap.Init(&priorityQueue)

	return priorityQueue
}

func GetDpcergAvg(months int) (float64, error) {
	results, err := getDpcergData()
	if err != nil {
		return 0, err
	}

	var items []map[string]interface{}

	for _, i := range results {
		item := i.(map[string]interface{})
		if item["SeriesCode"] == "DPCERG" {
			items = append(items, item)
		}
	}

	priorityQueue := makePriorityQueue(items)

	var sum float64
	var count int

	for priorityQueue.Len() > 0 {
		item := heap.Pop(&priorityQueue).(*services.Item)
		if count < months {
			sum += item.Value
			count++
		}
	}

	return sum / float64(months), nil
}
