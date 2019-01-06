package mappers

import (
	"fmt"
	"github.com/ASV44/DeliveryManagement-DS/warehouse/models"
)

func OrderErrorsToLog(insertErrors []models.OrderError, logMessage string) string {
	var log string
	for _, element := range insertErrors {
		log += fmt.Sprintf(logMessage, element.OrderID, element.Error)
	}

	return log
}
