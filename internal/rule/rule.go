package rule

import (
	"regexp"
)

type Rule struct {
	RuleID      string
	Description string
	Keywords    []string
	Pattern     *regexp.Regexp
	MinEntropy  float64
}
