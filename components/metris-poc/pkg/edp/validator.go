package edp

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateConsumptionMetrics(data string) (bool, error) {
	var consumptionMetrics *ConsumptionMetrics

	err := json.Unmarshal([]byte(data), &consumptionMetrics)
	if err != nil {
		fmt.Printf("invalid JSON: %v", err)
		return false, err
	}

	v := validator.New()
	err = v.Struct(consumptionMetrics)
	if err != nil {
		//TODO Remove me!
		fmt.Println(err)
		return false, err
	}
	if &consumptionMetrics != nil {
		//TODO Remove me!
		fmt.Println(consumptionMetrics.ResourceGroups)
		fmt.Println(consumptionMetrics.Compute)
		fmt.Println(consumptionMetrics.Networking)
		fmt.Println(consumptionMetrics.EventHub)
	}
	return true, nil
}
