Data Validator API

The Data Validator API is a service designed to validate common data inputs, such as email, CPF, name, and phone number. It supports flexible input formats and sanitizes data to ensure accuracy and consistency. The API automatically detects the type of parameter provided, sanitizes it, and returns a validation result in a structured JSON format.

Features

	•	Email Validation: Verifies if the email format is correct.
	•	CPF Validation: Checks if the Brazilian CPF number is valid, with flexible input formats (e.g., 04217318189, 042.173.181-89, etc.).
	•	Name Validation: Ensures that names contain only alphabetic characters, removes accents and special characters, and converts the name to uppercase.
	•	Phone Number Validation: Validates Brazilian phone numbers, allowing various formats.

API Endpoint

GET /validate

This endpoint receives various query parameters and automatically identifies which type of data to validate. Only one parameter should be provided at a time.

Query Parameters

	•	email - (optional) Email address to validate.
	•	cpf - (optional) Brazilian CPF number to validate, accepts different formats.
	•	name - (optional) Name to validate, sanitizes input by removing special characters and converting to uppercase.
	•	phone - (optional) Phone number to validate, accepts Brazilian phone number formats.

Response Format

The response is a JSON object that includes the validation results for the given parameter.

	•	status_code: HTTP status code indicating success or failure.
	•	parameter_key: The key of the parameter provided (e.g., “email”, “cpf”).
	•	raw_parameter_value: The original input provided by the user.
	•	parameter_value: The sanitized and validated value, if applicable.
	•	is_valid: Boolean indicating whether the parameter is valid.
	•	message: A message describing the validation result.
	•	request_id: Unique identifier for the request.
	•	timestamp: Timestamp of the request.
	•	execution_time_ms: Time taken to process the request in milliseconds.