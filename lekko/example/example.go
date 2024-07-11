package lekkoexample

// Example feature flag, enabled in development environments
func getExampleFlag(env string) bool {
	if env == "development" {
		return true
	}
	return false
}

// Example lekko that controls which LLM users interact with
func getExampleModel(isAdmin bool, plan string) string {
	if isAdmin {
		return "anthropic-claude3-opus"
	} else if plan == "team" || plan == "enterprise" {
		return "openai-chatgpt4-o"
	}
	return "openai-chatgpt3.5-turbo"
}

// Example lekko that controls sampling rate based on contextual factors
func getExampleSampleRate(env string, load float64, msgType string, throttle bool) float64 {
	if env == "staging" {
		return 1
	} else if msgType == "critical" {
		return 1
	} else if load > 0.75 || throttle {
		return 0.3
	}
	return 0.75
}

func getReturnText(env string) string {
	if env == "production" {
		return "foobar"
	} else if env == "development" {
		return "foo"
	}
	return "bar"
}
