package utils

import (
	"encoding/json"
	"fmt"

	"github.com/isurusiri/services/shared/models"
)

// PareseAPIResponse , performs json parsing
// of the APIResponse object.
func PareseAPIResponse(respBoday []byte) (*models.APIResponse, error) {
	var obj = new(models.APIResponse)
	err := json.Unmarshal(respBoday, &obj)

	if err != nil {
		fmt.Println(err)
	}

	return obj, err
}
