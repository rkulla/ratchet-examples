package mypkg

import (
	"github.com/dailyburn/ratchet/data"
	"github.com/dailyburn/ratchet/util"
)

type myTransformer struct{}

// Expose our DataProcessor for clients to use
func NewMyTransformer() *myTransformer {
	return &myTransformer{}
}

func (t *myTransformer) ProcessData(d data.JSON, outputChan chan data.JSON, killChan chan error) {
	// Step 1: Unmarshal json into slice of ReceivedData structs
	var users []ReceivedData
	var transforms []TransformedData

	err := data.ParseJSON(d, &users)
	util.KillPipelineIfErr(err, killChan)

	// Step 2: Loop through slice and transform data
	for _, user := range users {
		transform := TransformedData{}
		transform.UserID = user.ID
		transform.SomeNewField = "whatever"
		transforms = append(transforms, transform)
	}

	// Step 3: Marshal transformed data and send to next stage
	if len(transforms) > 0 {
		dd, err := data.NewJSON(transforms)
		util.KillPipelineIfErr(err, killChan)
		outputChan <- dd
	}
}

func (t *myTransformer) Finish(outputChan chan data.JSON, killChan chan error) {}

func (t *myTransformer) String() string {
	return "myTransformer"
}
