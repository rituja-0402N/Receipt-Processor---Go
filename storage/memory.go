package storage

import (
	"errors"
	"receipt-processor/models"
	"sync"
)

var (
	receipts = make(map[string]models.Receipt)
	points   = make(map[string]int)
	mu       sync.Mutex
)

func SaveReceipt(id string, receipt models.Receipt, receiptPoints int) {
	mu.Lock()
	defer mu.Unlock()
	receipts[id] = receipt
	points[id] = receiptPoints
}

func GetPoints(id string) (int, error) {
	mu.Lock()
	defer mu.Unlock()
	if p, exists := points[id]; exists {
		return p, nil
	}
	return 0, errors.New("receipt not found")
}
