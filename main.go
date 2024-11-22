package main

import (
	"receipt-processor/handlers"
	"reflect"
	"regexp"

	"github.com/gin-gonic/gin"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func main() {
	// Initialize custom validator
	validate = validator.New()

	// Register custom 'regex' tag validator
	validate.RegisterValidation("regex", func(fl validator.FieldLevel) bool {
		// Use the regex pattern directly
		regex := regexp.MustCompile(`^\d+\.\d{2}$`) // Matches numbers with two decimal places (e.g., 12.34)
		return regex.MatchString(fl.Field().String())
	})

	// Set Gin's validator to use the custom validator
	binding.Validator = &customValidator{validate}

	r := gin.Default()

	r.POST("/receipts/process", handlers.ProcessReceipt)
	r.GET("/receipts/:id/points", handlers.GetPoints)

	r.Run(":8080")
}

type customValidator struct {
	validator *validator.Validate
}

func (v *customValidator) ValidateStruct(obj interface{}) error {
	if obj == nil {
		return nil
	}

	// Handle pointers
	if kind := reflect.TypeOf(obj).Kind(); kind == reflect.Ptr {
		obj = reflect.ValueOf(obj).Elem().Interface()
	}

	return v.validator.Struct(obj)
}

func (v *customValidator) Engine() interface{} {
	return v.validator
}
