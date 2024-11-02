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
// @Summary Valida o formato dos dados fornecidos
// @Description Valida diferentes tipos de dados, como email, CPF, nome e telefone
// @Tags Validation
// @Param email query string false "Email a ser validado"
// @Param cpf query string false "CPF a ser validado"
// @Param nome query string false "Nome a ser validado"
// @Param telefone query string false "Telefone a ser validado"
// @Success 200 {object} models.ValidationResponse
// @Failure 400 {object} models.ValidationResponse
// @Router /validate [get]
func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var response models.ValidationResponse

	if email := r.URL.Query().Get("email"); email != "" {
		isValid, sanitizedValue, message := utils.ValidateEmail(email)
		response = createResponse("email", email, sanitizedValue, isValid, message, start)
	} else if cpf := r.URL.Query().Get("cpf"); cpf != "" {
		isValid, sanitizedValue, message := utils.ValidateCPF(cpf)
		response = createResponse("cpf", cpf, sanitizedValue, isValid, message, start)
	} else if nome := r.URL.Query().Get("nome"); nome != "" {
		isValid, sanitizedValue, message := utils.ValidateNome(nome)
		response = createResponse("nome", nome, sanitizedValue, isValid, message, start)
	} else if telefone := r.URL.Query().Get("telefone"); telefone != "" {
		isValid, sanitizedValue, message := utils.ValidateTelefone(telefone)
		response = createResponse("telefone", telefone, sanitizedValue, isValid, message, start)
	} else {
		response = models.ValidationResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Nenhum parâmetro de validação fornecido",
			IsValid:    false,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func createResponse(key, rawValue, sanitizedValue string, isValid bool, message string, start time.Time) models.ValidationResponse {
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
	}
}
