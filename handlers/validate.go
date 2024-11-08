package handlers

import (
	"DataValidatorAPI/models"
	"DataValidatorAPI/utils"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// ValidateHandler godoc
// @Summary Validates provided data format
// @Description Validates different types of data, including email, CPF, name, phone, RG, CEP, and credit card number
// @Tags Validation
// @Param email query string false "Email to be validated"
// @Param cpf query string false "CPF to be validated"
// @Param name query string false "Name to be validated"
// @Param telephone query string false "Phone number to be validated"
// @Param plastic query string false "Credit card number to be validated"
// @Param rg query string false "RG to be validated"
// @Param cep query string false "CEP (postal code) to be validated"
// @Success 200 {object} models.ValidationResponse
// @Failure 400 {object} models.ValidationResponse
// @Router /validate [get]
func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var response models.ValidationResponse

	if email := r.URL.Query().Get("email"); email != "" {
		isValid, sanitizedValue, message := utils.ValidateEmail(email)
		response = createResponse("email", email, sanitizedValue, isValid, message, start, false)
	} else if cpf := r.URL.Query().Get("cpf"); cpf != "" {
		isValid, sanitizedValue, message, fromCache := utils.ValidateCPFWithCache(cpf)
		response = createResponse("cpf", cpf, sanitizedValue, isValid, message, start, fromCache)
	} else if name := r.URL.Query().Get("name"); name != "" {
		isValid, sanitizedValue, message := utils.ValidateName(name)
		response = createResponse("name", name, sanitizedValue, isValid, message, start, false)
	} else if telephone := r.URL.Query().Get("telephone"); telephone != "" {
		isValid, sanitizedValue, message := utils.ValidatePhone(telephone)
		response = createResponse("telephone", telephone, sanitizedValue, isValid, message, start, false)
	} else if plastic := r.URL.Query().Get("plastic"); plastic != "" {
		isValid, sanitizedValue, message := utils.ValidatePlastic(plastic)
		response = createResponse("plastic", plastic, sanitizedValue, isValid, message, start, false)
	} else if rg := r.URL.Query().Get("rg"); rg != "" {
		isValid, sanitizedValue, message := utils.ValidateRG(rg)
		response = createResponse("rg", rg, sanitizedValue, isValid, message, start, false)
	} else if cep := r.URL.Query().Get("cep"); cep != "" {
		isValid, sanitizedValue, message := utils.ValidateCEP(cep)
		response = createResponse("cep", cep, sanitizedValue, isValid, message, start, false)
	} else {
		response = models.ValidationResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "No validation parameter provided",
			IsValid:    false,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

// createResponse creates a response with the provided data.
func createResponse(key, rawValue, sanitizedValue string, isValid bool, message string, start time.Time, fromCache bool) models.ValidationResponse {
	return models.ValidationResponse{
		StatusCode:        http.StatusOK,
		ParameterKey:      key,
		RawParameterValue: rawValue,
		ParameterValue:    sanitizedValue,
		IsValid:           isValid,
		Message:           message,
		RequestID:         uuid.New().String(),
		Timestamp:         time.Now(),
		ExecutionTimeMs:   int(time.Since(start).Milliseconds()),
		FromCache:         fromCache,
	}
}
