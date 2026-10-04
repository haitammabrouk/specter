package detector

import (
	"fmt"
	"specter/internal/rules"
	"strings"
)

func ApplyRules(line string) {

	rs := rules.CollectRules()
	for _, r := range rs {
		matches := r.Pattern.FindAllString(line, -1)
		fmt.Println(strings.Join(matches, ", "))
	}
}
