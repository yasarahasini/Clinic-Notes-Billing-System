package services

import "strings"

type ParsedData struct {
	Drugs []map[string]string
	Labs  []string
	Notes string
}

func ParseText(input string) ParsedData {
	data := ParsedData{}

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if strings.Contains(line, "mg") {
			data.Drugs = append(data.Drugs, map[string]string{
				"name":   line,
				"dosage": "1x daily",
			})
		} else if strings.Contains(line, "test") {
			data.Labs = append(data.Labs, line)
		} else {
			data.Notes += line + " "
		}
	}

	return data
}