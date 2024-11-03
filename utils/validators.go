package utils

import (
	"regexp"
	"strings"
	"unicode"
)

// ValidateEmail validates if the email format is correct
func ValidateEmail(email string) (bool, string, string) {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	if re.MatchString(email) {
		return true, email, "Valid email format"
	}
	return false, email, "Invalid email format"
}

// ValidateCPF validates if the CPF number format is correct
func ValidateCPF(cpf string) (bool, string, string) {
	re := regexp.MustCompile(`\D`)
	sanitizedCPF := re.ReplaceAllString(cpf, "")
	if len(sanitizedCPF) != 11 || !isValidCPF(sanitizedCPF) {
		return false, sanitizedCPF, "Invalid CPF format"
	}
	return true, sanitizedCPF, "Valid CPF format"
}

// isValidCPF checks if a CPF number is valid according to CPF rules
func isValidCPF(cpf string) bool {
	if cpf == "00000000000" || cpf == "11111111111" ||
		cpf == "22222222222" || cpf == "33333333333" ||
		cpf == "44444444444" || cpf == "55555555555" ||
		cpf == "66666666666" || cpf == "77777777777" ||
		cpf == "88888888888" || cpf == "99999999999" {
		return false
	}

	for i := 9; i < 11; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			num := int(cpf[j] - '0')
			sum += num * (i + 1 - j)
		}
		result := (sum * 10) % 11
		if result == 10 {
			result = 0
		}
		if result != int(cpf[i]-'0') {
			return false
		}
	}
	return true
}

// ValidateRG validates the format of a Brazilian RG (Registro Geral).
func ValidateRG(rg string) (bool, string, string) {
	re := regexp.MustCompile(`\D`)
	sanitizedRG := re.ReplaceAllString(rg, "")

	if len(sanitizedRG) < 7 || len(sanitizedRG) > 9 {
		return false, sanitizedRG, "Invalid RG format (incorrect length)"
	}

	return true, sanitizedRG, "Valid RG format"
}

// ValidateCEP validates a Brazilian postal code (CEP).
func ValidateCEP(cep string) (bool, string, string) {
	re := regexp.MustCompile(`\D`)
	sanitizedCEP := re.ReplaceAllString(cep, "")

	if len(sanitizedCEP) != 8 {
		return false, sanitizedCEP, "Invalid CEP format (incorrect length)"
	}

	return true, sanitizedCEP, "Valid CEP format"
}

// ValidateName sanitizes and validates the name format
func ValidateName(name string) (bool, string, string) {
	rawName := name
	sanitizedName := removeAccents(name)

	re := regexp.MustCompile(`[^a-zA-Z\s-']`)
	sanitizedName = re.ReplaceAllString(sanitizedName, "")

	sanitizedName = strings.TrimSpace(sanitizedName)
	sanitizedName = strings.ToUpper(sanitizedName)

	if len(sanitizedName) < 3 {
		return false, rawName, "Invalid name format (too short)"
	}
	if len(sanitizedName) > 60 {
		return false, rawName, "Invalid name format (too long)"
	}

	if strings.Contains(sanitizedName, "  ") {
		return false, rawName, "Invalid name format (contains multiple spaces)"
	}

	return true, sanitizedName, "Valid name format"
}

// removeAccents removes accents from characters
func removeAccents(input string) string {
	var output strings.Builder
	for _, char := range input {
		if unicode.IsLetter(char) {
			decomposed := unicode.ToLower(char)
			output.WriteRune(decomposed)
		} else if unicode.IsSpace(char) {
			output.WriteRune(char)
		}
	}
	return output.String()
}

// ValidatePhone validates and sanitizes phone numbers
func ValidatePhone(phone string) (bool, string, string) {
	re := regexp.MustCompile(`\D`)
	sanitizedPhone := re.ReplaceAllString(phone, "")
	if len(sanitizedPhone) == 11 || len(sanitizedPhone) == 10 {
		return true, sanitizedPhone, "Valid phone format"
	}
	return false, phone, "Invalid phone format"
}

// ValidatePlastic validates a credit card number using the Luhn Algorithm and determines its brand
func ValidatePlastic(cardNumber string) (bool, string, string) {
	re := regexp.MustCompile(`\D`)
	sanitizedCardNumber := re.ReplaceAllString(cardNumber, "")

	if len(sanitizedCardNumber) < 13 || len(sanitizedCardNumber) > 19 {
		return false, sanitizedCardNumber, "Invalid credit card number (incorrect length)"
	}

	// Check validity using the Luhn Algorithm
	if !isValidLuhn(sanitizedCardNumber) {
		return false, sanitizedCardNumber, "Invalid credit card number"
	}

	// Get the card brand
	brand := GetCardBrand(sanitizedCardNumber)
	return true, sanitizedCardNumber, "Valid credit card number (" + brand + ")"
}

// GetCardBrand identifies the card brand based on the BIN
func GetCardBrand(cardNumber string) string {
	cardNumber = strings.ReplaceAll(cardNumber, " ", "") // Remove spaces

	if len(cardNumber) >= 2 {
		switch {
		case strings.HasPrefix(cardNumber, "34") || strings.HasPrefix(cardNumber, "37"):
			return "American Express"
		case strings.HasPrefix(cardNumber, "36"):
			return "Diners Club"
		case strings.HasPrefix(cardNumber, "54") || strings.HasPrefix(cardNumber, "55") ||
			(cardNumber[:2] >= "51" && cardNumber[:2] <= "55"):
			return "MasterCard"
		case strings.HasPrefix(cardNumber, "4"):
			return "Visa"
		case len(cardNumber) >= 4 && (strings.HasPrefix(cardNumber, "6011") ||
			(cardNumber[:3] >= "644" && cardNumber[:3] <= "649") || strings.HasPrefix(cardNumber, "65")):
			return "Discover"
		case len(cardNumber) >= 4 && (strings.HasPrefix(cardNumber, "5067") || strings.HasPrefix(cardNumber, "4576") ||
			strings.HasPrefix(cardNumber, "4011")):
			return "Elo"
		case len(cardNumber) >= 6 && (strings.HasPrefix(cardNumber, "384100") || strings.HasPrefix(cardNumber, "384140") ||
			strings.HasPrefix(cardNumber, "384160") || strings.HasPrefix(cardNumber, "606282") || strings.HasPrefix(cardNumber, "637095")):
			return "Hipercard"
		}
	}

	return "Unknown"
}

// isValidLuhn implements the Luhn Algorithm to verify credit card numbers
func isValidLuhn(cardNumber string) bool {
	var sum int
	alt := false

	// Iterate over the card digits from back to front
	for i := len(cardNumber) - 1; i >= 0; i-- {
		n := int(cardNumber[i] - '0')
		if alt {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		alt = !alt
	}

	return sum%10 == 0
}
