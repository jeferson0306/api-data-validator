package utils

import (
	"regexp"
	"strings"
	"unicode"
)

func ValidateEmail(email string) (bool, string, string) {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	if re.MatchString(email) {
		return true, email, "Formato de email válido"
	}
	return false, email, "Formato de email inválido"
}

func ValidateCPF(cpf string) (bool, string, string) {
	re := regexp.MustCompile(`\D`)
	sanitizedCPF := re.ReplaceAllString(cpf, "")
	if len(sanitizedCPF) != 11 || !isValidCPF(sanitizedCPF) {
		return false, sanitizedCPF, "CPF inválido"
	}
	return true, sanitizedCPF, "Formato de CPF válido"
}

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

func ValidateNome(nome string) (bool, string, string) {
	rawNome := nome
	sanitizedNome := removeAccents(nome)

	re := regexp.MustCompile(`[^a-zA-Z\s]`)
	sanitizedNome = re.ReplaceAllString(sanitizedNome, "")

	sanitizedNome = strings.ToUpper(sanitizedNome)

	if sanitizedNome != "" {
		return true, sanitizedNome, "Formato de nome válido"
	}
	return false, rawNome, "Formato de nome inválido"
}

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

func ValidateTelefone(telefone string) (bool, string, string) {
	re := regexp.MustCompile(`\D`)
	sanitizedTelefone := re.ReplaceAllString(telefone, "")
	if len(sanitizedTelefone) == 11 || len(sanitizedTelefone) == 10 {
		return true, sanitizedTelefone, "Formato de telefone válido"
	}
	return false, telefone, "Formato de telefone inválido"
}
