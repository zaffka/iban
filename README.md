# IBAN Validator

[![Go Reference](https://pkg.go.dev/badge/github.com/zaffka/iban.svg)](https://pkg.go.dev/github.com/zaffka/iban)
[![Go Report Card](https://goreportcard.com/badge/github.com/zaffka/iban)](https://goreportcard.com/report/github.com/zaffka/iban)
[![Tests](https://github.com/zaffka/iban/actions/workflows/tests.yaml/badge.svg)](https://github.com/zaffka/iban/actions/workflows/tests.yaml)

A lightweight, zero-dependency Go package for validating International Bank Account Numbers (IBANs) according to ISO 13616.

## Installation

```bash
go get github.com/zaffka/iban
```

## Usage
```go
import "github.com/zaffka/iban"

func main() {
    iban.Valid("GB82WEST12345698765432")  // true
    iban.Valid("invalid-iban")           // false
}
```