package rules

import (
	"regexp"
	"specter/internal/rule"
)

func GitHubPersonalAccessToken() *rule.Rule {

	return &rule.Rule{
		RuleID:      "github-personal-access-token",
		Description: "GitHub personal access token (classic)",
		Keywords:    []string{"ghp_"},
		Pattern:     regexp.MustCompile(`ghp_[A-Za-z0-9]{36}\b`),
		MinEntropy:  0,
	}
}

func GitHubOAuthAccessToken() *rule.Rule {

	return &rule.Rule{
		RuleID:      "github-oauth-access-token",
		Description: "GitHub OAuth access token",
		Keywords:    []string{"gho_"},
		Pattern:     regexp.MustCompile(`gho_[A-Za-z0-9]{36}\b`),
		MinEntropy:  0,
	}
}

func GitHubUserToServerToken() *rule.Rule {

	return &rule.Rule{
		RuleID:      "github-user-to-server-token",
		Description: "GitHub user-to-server token",
		Keywords:    []string{"ghu_"},
		Pattern:     regexp.MustCompile(`ghu_[A-Za-z0-9]{36}\b`),
		MinEntropy:  0,
	}
}

func GitHubServerToServerToken() *rule.Rule {

	return &rule.Rule{
		RuleID:      "github-server-to-server-token",
		Description: "GitHub server-to-server token",
		Keywords:    []string{"ghs_"},
		Pattern:     regexp.MustCompile(`ghs_[A-Za-z0-9]{36}\b`),
		MinEntropy:  0,
	}
}

func GitHubRefreshToken() *rule.Rule {

	return &rule.Rule{
		RuleID:      "github-refresh-token",
		Description: "GitHub refresh token",
		Keywords:    []string{"ghr_"},
		Pattern:     regexp.MustCompile(`ghr_[A-Za-z0-9]{36}\b`),
		MinEntropy:  0,
	}
}

func GitHubFineGrainedPAT() *rule.Rule {

	return &rule.Rule{
		RuleID:      "github-fine-grained-pat",
		Description: "GitHub fine-grained personal access token",
		Keywords:    []string{"github_pat_"},
		Pattern:     regexp.MustCompile(`github_pat_[A-Za-z0-9]{22}_[A-Za-z0-9]{59}\b`),
		MinEntropy:  0,
	}
}
