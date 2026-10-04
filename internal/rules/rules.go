package rules

import (
	"specter/internal/rule"
)

func CollectRules() []rule.Rule {
	return []rule.Rule{
		*GitHubPersonalAccessToken(),
		*GitHubOAuthAccessToken(),
		*GitHubUserToServerToken(),
		*GitHubServerToServerToken(),
		*GitHubRefreshToken(),
		*GitHubFineGrainedPAT(),
	}
}
