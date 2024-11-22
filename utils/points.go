package utils

import (
	"math"
	"receipt-processor/models"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// 1. Points for alphanumeric characters in the retailer name
	points += len(regexp.MustCompile(`[a-zA-Z0-9]`).FindAllString(receipt.Retailer, -1))

	// 2. 50 points if the total is a round dollar amount with no cents
	if strings.HasSuffix(receipt.Total, ".00") {
		points += 50
	}

	// 3. 25 points if the total is a multiple of 0.25
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 4. 5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	// 5. Points for item descriptions whose trimmed length is a multiple of 3
	for _, item := range receipt.Items {
		descLen := len(strings.TrimSpace(item.ShortDescription))
		if descLen%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// 6. 6 points if the purchase date day is odd
	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 != 0 {
		points += 6
	}

	// 7. 10 points if the purchase time is between 2:00 PM and 4:00 PM
	t, _ := time.Parse("15:04", receipt.PurchaseTime)
	if t.Hour() == 14 {
		points += 10
	}

	return points
}
