package iban

import (
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

const (
	// MinIBANLength - minimum allowed IBAN length (country code + check digits + BBAN).
	MinIBANLength = 5

	// MaxIBANLength - maximum allowed IBAN length (SWIFT standard).
	MaxIBANLength = 34
)

var (
	// IbanRegex validates the basic IBAN structure.
	// - Starts with 2 letter country code
	// - Followed by 2 check digits
	// - Ends with alphanumeric BBAN (country-specific format).
	IbanRegex = regexp.MustCompile(`^[A-Z]{2}\d{2}[A-Z0-9]{1,30}$`)

	// ModValue - modulus value in big.Int format for IBAN validation (ISO 13616).
	ModValue = big.NewInt(97)
)

// Valid checks whether the provided string is a valid IBAN according to ISO 13616.
//
// Validation includes:
//  1. Basic format check (length and regex pattern)
//  2. Rearranging the string (moving first 4 chars to end)
//  3. Converting letters to numbers (A=10, B=11, ..., Z=35)
//  4. Verifying the checksum using mod-97 operation
//
// Returns:
//   - true if the IBAN passes all validation steps
//   - false if any validation step fails.
func Valid(iban string) bool {
	// Step 0: Sanitize input and basic format check.
	cleanIBAN := strings.ToUpper(strings.ReplaceAll(iban, " ", ""))

	if len(cleanIBAN) < MinIBANLength || len(cleanIBAN) > MaxIBANLength {
		return false
	}

	if !IbanRegex.MatchString(cleanIBAN) {
		return false
	}

	// Step 1: Rearrange - move first 4 characters to the end.
	rearranged := cleanIBAN[4:] + cleanIBAN[:4]

	// Step 2: Convert letters to numbers.
	builder := strings.Builder{}
	builder.Grow(len(rearranged) * 2)

	for _, r := range rearranged {
		if r >= 'A' && r <= 'Z' {
			builder.WriteString(strconv.Itoa(int(r - 'A' + 10)))
		}

		if r >= '0' && r <= '9' {
			builder.WriteByte(byte(r))
		}
	}

	// Step 3: Big integer "mod - 97" calculation.
	number, ok := new(big.Int).SetString(builder.String(), 10)
	if !ok {
		return false
	}

	return new(big.Int).Mod(number, ModValue).Int64() == 1
}
