package rule

import (
	"regexp"
)

type Rule struct {
	RuleId string
	Description string
	Keywords []string
	Pattern *regexp.Regexp
	MinEntropy float64
}