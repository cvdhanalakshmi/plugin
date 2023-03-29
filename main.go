package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Request struct {
	Standard      string     `json:"standard"`
	Code          string     `json:"code"`
	Name          string     `json:"name"`
	Importance    string     `json:"importance"`
	DetailHeaders []string   `json:"detailHeaders"`
	DetailTypes   []string   `json:"detailTypes"`
	Failures      []Failures `json:"failures"`
}
type Asset struct {
	Type       string `json:"type"`
	SubType    string `json:"subType"`
	Identifier string `json:"identifier"`
}
type Details struct {
	Data []string `json:"data"`
}
type Failures struct {
	Asset       Asset     `json:"asset"`
	AssetUUID   string    `json:"assetUuid"`
	ProfileUUID string    `json:"profileUuid"`
	Details     []Details `json:"details"`
}

func main() {
	data := ReadJson()
	fmt.Println(data)

}

func ReadJson() Response {
	filename := `D:\Users\vdhanalakshmi\Documents\checkmarxResponse.json`
	plan, err1 := ioutil.ReadFile(filename) // filename is the JSON file to read
	if err1 != nil {
		fmt.Println(err1)
	}
	// fmt.Println(err1,plan)
	var data []Request

	err := json.Unmarshal(plan, &data)
	if err != nil {
		fmt.Println("Cannot unmarshal the json ", err)
	}
	var res Response
	for _, val := range data {
		res.FailureCount = len(val.Failures[0].Details)
		res.Type = val.Failures[0].Asset.Type
		res.SubType = val.Failures[0].Asset.SubType
		res.Identifier = val.Failures[0].Asset.Identifier
		var failres []Failure

		for _, failure := range val.Failures {
			var fail Failure
			fail.Details.Headers = val.DetailHeaders
			fail.Details.Types = val.DetailTypes
			for _, details := range failure.Details {
				fail.Details.Data = append(fail.Details.Data, details.Data)
			}
			fail.Name = val.Name
			fail.Importance = val.Importance
			fail.AssetType = failure.Asset.Type
			failres = append(failres, fail)
		}
		res.Failures = failres

	}
	return res
}

type Response struct {
	UUID         string    `json:"uuid"`
	Account      Account   `json:"account"`
	Type         string    `json:"type"`
	SubType      string    `json:"subType"`
	Identifier   string    `json:"Identifier"`
	LastEvalTime time.Time `json:"lastEvalTime"`
	Riskscore    struct {
	} `json:"riskscore"`
	FailureCount int        `json:"failureCount"`
	Failures     []Failure  `json:"failures"`
	PassCount    int        `json:"passCount"`
	Passes       []any      `json:"passes"`
	Profiles     []Profiles `json:"profiles"`
}

type Account struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}
type Failure struct {
	EvalTime    time.Time `json:"evalTime"`
	ControlUUID string    `json:"controlUuid"`
	TestUUID    string    `json:"testUuid"`
	AssetType   string    `json:"assetType"`
	Name        string    `json:"name"`
	Importance  string    `json:"importance"`
	Details     Detail    `json:"details"`
}

type Profiles struct {
	UUID       string `json:"uuid"`
	Identifier string `json:"identifier"`
}
type Detail struct {
	Headers []string   `json:"headers"`
	Types   []string   `json:"types"`
	Data    [][]string `json:"data"`
}
