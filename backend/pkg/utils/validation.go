package utils

/*
Key Features:
Enhanced Validation:

✅ Structured Error Messages - User-friendly error responses
✅ JSON Field Names - Proper field names in errors
✅ Custom Validators - 10+ custom validation rules
✅ Helper Functions - Quick validation for common cases

Custom Validators:

🔐 Strong Password - Requires uppercase, lowercase, number, special char
👤 Username - 3-20 chars, letters/numbers/underscores only
🔗 Slug - URL-friendly slugs
🎨 Hex Color - Valid hex color codes (#RRGGBB)
💳 Credit Card - Luhn algorithm validation
📱 Phone - International phone format
🌍 Timezone - Timezone string validation
📁 File Extension - File extension validation
📄 MIME Type - MIME type format validation

Ready-to-Use Structs:

📄 PaginationParams - Page, limit, sort validation
🔍 SearchParams - Search query validation
📎 FileUploadParams - File upload validation
📧 ContactInfo - Email, phone, social media
🏠 Address - Complete address validation
💰 PaymentInfo - Payment card validation
*/

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
    validate = validator.New()
    
    // Register custom tag name function for better JSON field names in errors
    validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
        name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
        if name == "-" {
            return ""
        }
        return name
    })
    
    // Register custom validators
    registerCustomValidators()
}

// ValidateStruct validates a struct and returns formatted errors
func ValidateStruct(s interface{}) error {
    err := validate.Struct(s)
    if err != nil {
        return FormatValidationErrors(err)
    }
    return nil
}

// ValidateVar validates a single variable
func ValidateVar(field interface{}, tag string) error {
    err := validate.Var(field, tag)
    if err != nil {
        return FormatValidationErrors(err)
    }
    return nil
}

// ValidationErrors represents structured validation errors
type ValidationErrors struct {
    Errors []FieldError `json:"errors"`
}

func (v ValidationErrors) Error() string {
    var messages []string
    for _, err := range v.Errors {
        messages = append(messages, err.Message)
    }
    return strings.Join(messages, "; ")
}

// FieldError represents a single field validation error
type FieldError struct {
    Field   string `json:"field"`
    Value   string `json:"value"`
    Tag     string `json:"tag"`
    Message string `json:"message"`
}

// FormatValidationErrors converts validator errors to structured format
func FormatValidationErrors(err error) error {
    var validationErrors []FieldError
    
    if validatorErrors, ok := err.(validator.ValidationErrors); ok {
        for _, e := range validatorErrors {
            validationErrors = append(validationErrors, FieldError{
                Field:   getJSONFieldName(e),
                Value:   fmt.Sprintf("%v", e.Value()),
                Tag:     e.Tag(),
                Message: getErrorMessage(e),
            })
        }
    }
    
    return ValidationErrors{Errors: validationErrors}
}

// getJSONFieldName extracts JSON field name from validator error
func getJSONFieldName(e validator.FieldError) string {
    field := e.Field()
    if field == "" {
        return e.StructField()
    }
    return field
}

// getErrorMessage returns user-friendly error messages
func getErrorMessage(e validator.FieldError) string {
    field := getJSONFieldName(e)
    
    switch e.Tag() {
    case "required":
        return fmt.Sprintf("%s is required", field)
    case "email":
        return fmt.Sprintf("%s must be a valid email address", field)
    case "min":
        return fmt.Sprintf("%s must be at least %s characters long", field, e.Param())
    case "max":
        return fmt.Sprintf("%s must be at most %s characters long", field, e.Param())
    case "len":
        return fmt.Sprintf("%s must be exactly %s characters long", field, e.Param())
    case "gte":
        return fmt.Sprintf("%s must be greater than or equal to %s", field, e.Param())
    case "lte":
        return fmt.Sprintf("%s must be less than or equal to %s", field, e.Param())
    case "gt":
        return fmt.Sprintf("%s must be greater than %s", field, e.Param())
    case "lt":
        return fmt.Sprintf("%s must be less than %s", field, e.Param())
    case "alphanum":
        return fmt.Sprintf("%s must contain only alphanumeric characters", field)
    case "alpha":
        return fmt.Sprintf("%s must contain only alphabetic characters", field)
    case "numeric":
        return fmt.Sprintf("%s must contain only numeric characters", field)
    case "url":
        return fmt.Sprintf("%s must be a valid URL", field)
    case "uri":
        return fmt.Sprintf("%s must be a valid URI", field)
    case "uuid":
        return fmt.Sprintf("%s must be a valid UUID", field)
    case "oneof":
        return fmt.Sprintf("%s must be one of: %s", field, e.Param())
    case "eqfield":
        return fmt.Sprintf("%s must equal %s", field, e.Param())
    case "nefield":
        return fmt.Sprintf("%s must not equal %s", field, e.Param())
    case "unique":
        return fmt.Sprintf("%s values must be unique", field)
    case "base64":
        return fmt.Sprintf("%s must be valid base64", field)
    case "json":
        return fmt.Sprintf("%s must be valid JSON", field)
    case "jwt":
        return fmt.Sprintf("%s must be a valid JWT token", field)
    case "phone":
        return fmt.Sprintf("%s must be a valid phone number", field)
    case "strong_password":
        return fmt.Sprintf("%s must contain at least 8 characters with uppercase, lowercase, number and special character", field)
    case "username":
        return fmt.Sprintf("%s must be 3-20 characters long and contain only letters, numbers, and underscores", field)
    case "slug":
        return fmt.Sprintf("%s must be a valid slug (lowercase letters, numbers, and hyphens)", field)
    case "hex_color":
        return fmt.Sprintf("%s must be a valid hex color code", field)
    case "credit_card":
        return fmt.Sprintf("%s must be a valid credit card number", field)
    case "isbn":
        return fmt.Sprintf("%s must be a valid ISBN", field)
    case "mac":
        return fmt.Sprintf("%s must be a valid MAC address", field)
    case "latitude":
        return fmt.Sprintf("%s must be a valid latitude", field)
    case "longitude":
        return fmt.Sprintf("%s must be a valid longitude", field)
    case "timezone":
        return fmt.Sprintf("%s must be a valid timezone", field)
    case "file_extension":
        return fmt.Sprintf("%s must have a valid file extension", field)
    case "mime_type":
        return fmt.Sprintf("%s must be a valid MIME type", field)
    default:
        return fmt.Sprintf("%s is invalid", field)
    }
}

// registerCustomValidators registers all custom validation rules
func registerCustomValidators() {
    // Strong password validator
    validate.RegisterValidation("strong_password", validateStrongPassword)
    
    // Username validator
    validate.RegisterValidation("username", validateUsername)
    
    // Slug validator
    validate.RegisterValidation("slug", validateSlug)
    
    // Hex color validator
    validate.RegisterValidation("hex_color", validateHexColor)
    
    // Credit card validator
    validate.RegisterValidation("credit_card", validateCreditCard)
    
    // Phone number validator
    validate.RegisterValidation("phone", validatePhone)
    
    // Timezone validator
    validate.RegisterValidation("timezone", validateTimezone)
    
    // File extension validator
    validate.RegisterValidation("file_extension", validateFileExtension)
    
    // MIME type validator
    validate.RegisterValidation("mime_type", validateMimeType)
}

// Custom validation functions

// validateStrongPassword validates password strength
func validateStrongPassword(fl validator.FieldLevel) bool {
    password := fl.Field().String()
    
    if len(password) < 8 {
        return false
    }
    
    var hasUpper, hasLower, hasNumber, hasSpecial bool
    
    for _, char := range password {
        switch {
        case unicode.IsUpper(char):
            hasUpper = true
        case unicode.IsLower(char):
            hasLower = true
        case unicode.IsDigit(char):
            hasNumber = true
        case unicode.IsPunct(char) || unicode.IsSymbol(char):
            hasSpecial = true
        }
    }
    
    return hasUpper && hasLower && hasNumber && hasSpecial
}

// validateUsername validates username format
func validateUsername(fl validator.FieldLevel) bool {
    username := fl.Field().String()
    
    if len(username) < 3 || len(username) > 20 {
        return false
    }
    
    // Only letters, numbers, and underscores
    matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, username)
    return matched
}

// validateSlug validates URL slug format
func validateSlug(fl validator.FieldLevel) bool {
    slug := fl.Field().String()
    
    // Lowercase letters, numbers, and hyphens only
    matched, _ := regexp.MatchString(`^[a-z0-9-]+$`, slug)
    return matched
}

// validateHexColor validates hex color codes
func validateHexColor(fl validator.FieldLevel) bool {
    color := fl.Field().String()
    
    // Match #RRGGBB or #RGB format
    matched, _ := regexp.MatchString(`^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`, color)
    return matched
}

// validateCreditCard validates credit card numbers using Luhn algorithm
func validateCreditCard(fl validator.FieldLevel) bool {
    cardNumber := strings.ReplaceAll(fl.Field().String(), " ", "")
    
    if len(cardNumber) < 13 || len(cardNumber) > 19 {
        return false
    }
    
    // Check if all characters are digits
    for _, char := range cardNumber {
        if !unicode.IsDigit(char) {
            return false
        }
    }
    
    // Luhn algorithm
    sum := 0
    alternate := false
    
    for i := len(cardNumber) - 1; i >= 0; i-- {
        digit := int(cardNumber[i] - '0')
        
        if alternate {
            digit *= 2
            if digit > 9 {
                digit = (digit % 10) + 1
            }
        }
        
        sum += digit
        alternate = !alternate
    }
    
    return sum%10 == 0
}

// validatePhone validates phone numbers (basic international format)
func validatePhone(fl validator.FieldLevel) bool {
    phone := fl.Field().String()
    
    // Basic international phone number format
    matched, _ := regexp.MatchString(`^\+?[1-9]\d{1,14}$`, phone)
    return matched
}

// validateTimezone validates timezone strings
func validateTimezone(fl validator.FieldLevel) bool {
    timezone := fl.Field().String()
    
    // Basic timezone validation (can be enhanced with actual timezone list)
    matched, _ := regexp.MatchString(`^[A-Za-z_]+/[A-Za-z_]+$`, timezone)
    return matched
}

// validateFileExtension validates file extensions
func validateFileExtension(fl validator.FieldLevel) bool {
    filename := fl.Field().String()
    
    // Check if filename has an extension
    matched, _ := regexp.MatchString(`\.[a-zA-Z0-9]+$`, filename)
    return matched
}

// validateMimeType validates MIME types
func validateMimeType(fl validator.FieldLevel) bool {
    mimeType := fl.Field().String()
    
    // Basic MIME type format validation
    matched, _ := regexp.MatchString(`^[a-zA-Z0-9][a-zA-Z0-9!#$&\-\^]*\/[a-zA-Z0-9][a-zA-Z0-9!#$&\-\^]*$`, mimeType)
    return matched
}

// Validation helper functions

// ValidateEmail validates email format
func ValidateEmail(email string) error {
    return ValidateVar(email, "required,email")
}

// ValidatePassword validates password strength
func ValidatePassword(password string) error {
    return ValidateVar(password, "required,strong_password")
}

// ValidateUsername validates username format
func ValidateUsername(username string) error {
    return ValidateVar(username, "required,username")
}

// ValidateURL validates URL format
func ValidateURL(url string) error {
    return ValidateVar(url, "required,url")
}

// ValidateUUID validates UUID format
func ValidateUUID(uuid string) error {
    return ValidateVar(uuid, "required,uuid")
}

// ValidateJSON validates JSON string
func ValidateJSON(jsonStr string) error {
    return ValidateVar(jsonStr, "required,json")
}

// ValidatePhone validates phone number
func ValidatePhone(phone string) error {
    return ValidateVar(phone, "required,phone")
}

// ValidateCreditCard validates credit card number
func ValidateCreditCard(cardNumber string) error {
    return ValidateVar(cardNumber, "required,credit_card")
}

// Validation for common struct patterns

// PaginationParams represents pagination parameters
type PaginationParams struct {
    Page     int `json:"page" validate:"min=1"`
    Limit    int `json:"limit" validate:"min=1,max=100"`
    SortBy   string `json:"sort_by" validate:"omitempty,oneof=id name email created_at updated_at"`
    SortDir  string `json:"sort_dir" validate:"omitempty,oneof=asc desc"`
}

// SearchParams represents search parameters
type SearchParams struct {
    Query    string `json:"query" validate:"omitempty,min=1,max=100"`
    Category string `json:"category" validate:"omitempty,alpha"`
    Tags     []string `json:"tags" validate:"omitempty,dive,min=1,max=50"`
}

// FileUploadParams represents file upload validation
type FileUploadParams struct {
    FileName    string `json:"file_name" validate:"required,min=1,max=255,file_extension"`
    FileSize    int64  `json:"file_size" validate:"required,min=1,max=10485760"` // 10MB max
    MimeType    string `json:"mime_type" validate:"required,mime_type"`
    ContentType string `json:"content_type" validate:"required,oneof=image/jpeg image/png image/gif application/pdf text/plain"`
}

// ContactInfo represents contact information validation
type ContactInfo struct {
    Email       string `json:"email" validate:"required,email"`
    Phone       string `json:"phone" validate:"omitempty,phone"`
    Website     string `json:"website" validate:"omitempty,url"`
    SocialMedia map[string]string `json:"social_media" validate:"omitempty,dive,keys,oneof=twitter facebook instagram linkedin,endkeys,url"`
}

// Address represents address validation
type Address struct {
    Street     string `json:"street" validate:"required,min=5,max=100"`
    City       string `json:"city" validate:"required,min=2,max=50,alpha"`
    State      string `json:"state" validate:"required,min=2,max=50"`
    ZipCode    string `json:"zip_code" validate:"required,min=5,max=10,alphanum"`
    Country    string `json:"country" validate:"required,len=2,alpha"` // ISO country code
    Latitude   float64 `json:"latitude" validate:"omitempty,latitude"`
    Longitude  float64 `json:"longitude" validate:"omitempty,longitude"`
}

// PaymentInfo represents payment information validation
type PaymentInfo struct {
    CardNumber    string `json:"card_number" validate:"required,credit_card"`
    ExpiryMonth   int    `json:"expiry_month" validate:"required,min=1,max=12"`
    ExpiryYear    int    `json:"expiry_year" validate:"required,min=2024,max=2050"`
    CVV           string `json:"cvv" validate:"required,len=3,numeric"`
    HolderName    string `json:"holder_name" validate:"required,min=2,max=100"`
    BillingAddress Address `json:"billing_address" validate:"required"`
}

// ValidateEmailList validates a list of email addresses
func ValidateEmailList(emails []string) error {
    for i, email := range emails {
        if err := ValidateEmail(email); err != nil {
            return fmt.Errorf("invalid email at index %d: %w", i, err)
        }
    }
    return nil
}

// ValidatePasswordConfirmation validates password and confirmation match
func ValidatePasswordConfirmation(password, confirmPassword string) error {
    if err := ValidatePassword(password); err != nil {
        return err
    }
    
    if password != confirmPassword {
        return errors.New("password confirmation does not match")
    }
    
    return nil
}

// ValidateRequiredFields validates that required fields are not empty
func ValidateRequiredFields(fields map[string]interface{}) error {
    var missingFields []string
    
    for fieldName, value := range fields {
        if value == nil {
            missingFields = append(missingFields, fieldName)
            continue
        }
        
        switch v := value.(type) {
        case string:
            if strings.TrimSpace(v) == "" {
                missingFields = append(missingFields, fieldName)
            }
        case []string:
            if len(v) == 0 {
                missingFields = append(missingFields, fieldName)
            }
        }
    }
    
    if len(missingFields) > 0 {
        return fmt.Errorf("required fields missing: %s", strings.Join(missingFields, ", "))
    }
    
    return nil
}