package mappers

import (
	"fmt"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/models"
)

func InsertErrorsToLog(insertErrors []models.InsertError, logMessage string) string {
	var log string
	for _, element := range insertErrors {
		log += fmt.Sprintf(logMessage, element.OrderID, element.OrderAwbNumber, element.Error)
	}

	return log
}
