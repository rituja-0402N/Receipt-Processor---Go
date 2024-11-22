Receipt Processor

Description

Receipt Processor is a web service designed to process and evaluate receipts according to predefined rules. The service includes endpoints to submit receipts, calculate points based on receipt data, and retrieve the points awarded for a specific receipt.

Features

Submit Receipt: Accepts receipt data and returns a unique ID for the receipt.
Get Points: Retrieves the total points awarded for a submitted receipt based on predefined rules.
In-memory storage: Data persists only during the application's runtime, ensuring simplicity without requiring a database.
Endpoints

1. Submit Receipt
   Path: /receipts/process
   Method: POST
   Request Body:
   {
   "retailer": "Target",
   "purchaseDate": "2022-01-01",
   "purchaseTime": "13:01",
   "items": [
   { "shortDescription": "Mountain Dew 12PK", "price": "6.49" },
   { "shortDescription": "Emils Cheese Pizza", "price": "12.25" }
   ],
   "total": "35.35"
   }
   Response:
   {
   "id": "7fb1377b-b223-49d9-a31a-5a02701dd310"
   }
   Error Response:
   {
   "error": "Key: 'Receipt.Total' Error:Field validation for 'Total' failed on the 'regex' tag"
   }
2. Get Points
   Path: /receipts/{id}/points
   Method: GET
   Response:
   {
   "points": 28
   }
   Error Response:
   {
   "error": "Receipt not found"
   }
   Points Calculation Rules

1 point for every alphanumeric character in the retailer name.
50 points if the total is a round dollar amount with no cents.
25 points if the total is a multiple of 0.25.
5 points for every two items on the receipt.
Description-based points:
If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer.
6 points if the day in the purchase date is odd.
10 points if the time of purchase is between 2:00 PM and 4:00 PM.
Getting Started

1. Prerequisites
   Go (v1.19 or later)
   cURL or Postman for API testing
   Optionally, Docker (for containerized deployment)
2. Clone the Repository
   git clone https://github.com/your-username/receipt-processor.git
   cd receipt-processor
3. Run Locally
   Build the project:
   go build
   Start the server:
   go run main.go
   The server runs on http://localhost:8080.
4. Test the API
   Using cURL

Submit a Receipt:
curl -X POST http://localhost:8080/receipts/process \
-H "Content-Type: application/json" \
-d '{
"retailer": "Target",
"purchaseDate": "2022-01-01",
"purchaseTime": "13:01",
"items": [
{ "shortDescription": "Mountain Dew 12PK", "price": "6.49" },
{ "shortDescription": "Emils Cheese Pizza", "price": "12.25" }
],
"total": "35.35"
}'
Get Points: Replace {id} with the ID from the previous step.
curl -X GET http://localhost:8080/receipts/{id}/points 5. Run with Docker
Build the Docker image:
docker build -t receipt-processor .
Run the Docker container:
docker run -p 8080:8080 receipt-processor
Project Structure

receipt-processor/
├── handlers/ # API endpoint handlers
│ └── receipt.go
├── models/ # Data models for receipts and items
│ └── receipt.go
├── storage/ # In-memory storage implementation
│ └── memory.go
├── utils/ # Utility functions for point calculations
│ └── points.go
├── main.go # Application entry point
├── Dockerfile # Docker configuration
├── go.mod # Go module file
└── README.md # Project documentation
Testing

Unit Tests
Run unit tests to verify the functionality of point calculation and handlers:

go test ./...
End-to-End Tests
Use Postman or cURL to simulate requests to the API endpoints.

Load Testing
Use k6 or similar tools to perform load testing:

k6 run load-test.js
Future Enhancements

Add persistent storage (e.g., PostgreSQL or MongoDB).
Enhance error handling for edge cases.
Add authentication and rate limiting for API security.
Provide Swagger/OpenAPI documentation for better API usability.
License
