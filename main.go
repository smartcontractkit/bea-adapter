package main

import (
	"container/heap"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/linkpoolio/bridges"
	"github.com/smartcontractkit/bea-adapter/services"
)

// Bea contains the data format for the response
type Bea struct {
	BEAAPI struct {
		Request struct {
			RequestParam []struct {
				ParameterName  string `json:"ParameterName"`
				ParameterValue string `json:"ParameterValue"`
			} `json:"RequestParam"`
		} `json:"Request"`
		Results struct {
			Statistic         string `json:"Statistic"`
			UTCProductionTime string `json:"UTCProductionTime"`
			Dimensions        []struct {
				Ordinal  string `json:"Ordinal"`
				Name     string `json:"Name"`
				DataType string `json:"DataType"`
				IsValue  string `json:"IsValue"`
			} `json:"Dimensions"`
			Data []struct {
				TableName       string `json:"TableName"`
				SeriesCode      string `json:"SeriesCode"`
				LineNumber      string `json:"LineNumber"`
				LineDescription string `json:"LineDescription"`
				TimePeriod      string `json:"TimePeriod"`
				METRICNAME      string `json:"METRIC_NAME"`
				CLUNIT          string `json:"CL_UNIT"`
				UNITMULT        string `json:"UNIT_MULT"`
				DataValue       string `json:"DataValue"`
				NoteRef         string `json:"NoteRef,omitempty"`
			} `json:"Data"`
			Notes []struct {
				NoteRef  string `json:"NoteRef"`
				NoteText string `json:"NoteText"`
			} `json:"Notes"`
		} `json:"Results"`
	} `json:"BEAAPI"`
}

var (
	// APIKey is value for the API_KEY environment variable
	APIKey = os.Getenv("API_KEY")
	// BaseURL is the URL to reach out to
	BaseURL = "https://apps.bea.gov/api/data/"
)

// Run calls the endpoint and returns the resulting data
func (b *Bea) Run(h *bridges.Helper) (interface{}, error) {
	bea := Bea{}
	err := h.HTTPCallWithOpts(
		http.MethodGet,
		BaseURL,
		&bea,
		bridges.CallOpts{
			Auth: bridges.NewAuth(bridges.AuthParam, "UserID", APIKey),
			Query: map[string]interface{}{
				"DataSetName":  "NIPA",
				"TableName":    "T20804",
				"ResultFormat": "json",
				"method":       "getData",
				"Frequency":    "M",
				"Year":         "2019,2018",
			},
		},
	)

	var dataValues []string
	var timePeriods []string

	for _, v := range bea.BEAAPI.Results.Data {
		if v.SeriesCode == "DPCERG" {
			dataValues = append(dataValues, v.DataValue)
			timePeriods = append(timePeriods, v.TimePeriod)
		}
	}

	priorityQueue := makePriorityQueue(dataValues, timePeriods)

	var sum float64
	var count int

	for priorityQueue.Len() > 0 && count < 3 {
		item, ok := heap.Pop(&priorityQueue).(*services.Item)
		if !ok {
			return 0, errors.New("unable to cast queue item type")
		}
		sum += item.Value
		count++
	}

	result := make(map[string]interface{})
	result["result"] = sum / float64(3)

	return result, err
}

// Opts is the bridges.Bridge implementation
func (b *Bea) Opts() *bridges.Opts {
	return &bridges.Opts{
		Name:   "BEA",
		Lambda: true,
	}
}

func main() {
	bridges.NewServer(&Bea{}).Start(8080)
}

func makePriorityQueue(dataValues []string, timePeriods []string) services.PriorityQueue {
	priorityQueue := services.PriorityQueue{}

	for k, item := range dataValues {
		val, err := strconv.ParseFloat(fmt.Sprint(item), 64)
		if err != nil {
			continue
		}

		date := strings.Split(fmt.Sprint(timePeriods[k]), "M")
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
