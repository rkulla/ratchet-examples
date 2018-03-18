package mypkg

import (
	"github.com/dailyburn/ratchet/data"
	"github.com/dailyburn/ratchet/logger"
	"github.com/dailyburn/ratchet/util"
)

type myTransformer struct {
	processedUserData    []ReceivedData1
	processedAddressData []ReceivedData2
}

// Expose our DataProcessor for clients to use
func NewMyTransformer() *myTransformer {
	return &myTransformer{}
}

func (t *myTransformer) ProcessData(d data.JSON, outputChan chan data.JSON, killChan chan error) {
	// Step 1: Unmarshal json into ReceivedData1 or ReceivedData2
	var users []ReceivedData1
	var addresses []ReceivedData2
	if t.validateAndParseUsers(d, &users) {
		logger.Debug("data parsed into []ReceivedData1")
		t.processedUserData = append(t.processedUserData, users...)
	} else if t.validateAndParseAddresses(d, &addresses) {
		logger.Debug("data parsed into []ReceivedData2")
		t.processedAddressData = append(t.processedAddressData, addresses...)
	}
}

func (t *myTransformer) Finish(outputChan chan data.JSON, killChan chan error) {
	// Step 2: Loop through batched data and transform it into one
	var transforms []TransformedData
	// Note consider using a map instead, keyed off of ID
	for _, user := range t.processedUserData {
		for _, address := range t.processedAddressData {
			transform := TransformedData{}
			if user.ID == address.ID {
				transform.UserID = user.ID
				transform.Name = user.Name
				transform.City = address.City
				transforms = append(transforms, transform)
			}
		}
	}

	// Step 3: Marshal transformed data and send to next stage
	if len(transforms) > 0 {
		dd, err := data.NewJSON(transforms)
		util.KillPipelineIfErr(err, killChan)
		outputChan <- dd
	}
}

func (t *myTransformer) validateAndParseUsers(d data.JSON, users *[]ReceivedData1) bool {
	err := data.ParseJSONSilent(d, users)
	if len(*users) > 0 {
		if err == nil {
			if (*users)[0].TypeHelper == Query1Helper {
				return true
			}
		} else {
			logger.Debug("validateAndParseUsers err is: ", err)
		}
	}
	return false
}

func (t *myTransformer) validateAndParseAddresses(d data.JSON, addresses *[]ReceivedData2) bool {
	err := data.ParseJSONSilent(d, addresses)
	if len(*addresses) > 0 {
		if err == nil {
			if (*addresses)[0].TypeHelper == Query2Helper {
				return true
			}
		} else {
			logger.Debug("validateAndParseAddresses err is: ", err)
		}
	}
	return false
}

func (t *myTransformer) String() string {
	return "myTransformer"
}
