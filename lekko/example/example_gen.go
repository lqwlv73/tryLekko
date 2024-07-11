// Generated by Lekko. DO NOT EDIT.
package lekkoexample

import (
	"context"
	"errors"

	"github.com/lekkodev/go-sdk/client"
	"github.com/lekkodev/go-sdk/pkg/debug"
)

type LekkoClient struct {
	client.Client
}

// Example feature flag, enabled in development environments
func (c *LekkoClient) GetExampleFlag(env string) bool {
	ctx := context.Background()
	ctx = client.Add(ctx, "env", env)
	result, err := c.GetBool(ctx, "example", "example-flag")
	if err == nil {
		return result
	}
	result = getExampleFlag(env)
	if !errors.Is(err, client.ErrNoOpProvider) {
		debug.LogError("Lekko evaluation error", "err", err)
	}
	debug.LogDebug("Lekko fallback", "result", result)
	return result
}

// Example lekko that controls which LLM users interact with
func (c *LekkoClient) GetExampleModel(isAdmin bool, plan string) string {
	ctx := context.Background()
	ctx = client.Add(ctx, "is_admin", isAdmin)
	ctx = client.Add(ctx, "plan", plan)
	result, err := c.GetString(ctx, "example", "example-model")
	if err == nil {
		return result
	}
	result = getExampleModel(isAdmin, plan)
	if !errors.Is(err, client.ErrNoOpProvider) {
		debug.LogError("Lekko evaluation error", "err", err)
	}
	debug.LogDebug("Lekko fallback", "result", result)
	return result
}

// Example lekko that controls sampling rate based on contextual factors
func (c *LekkoClient) GetExampleSampleRate(env string, load float64, msgType string, throttle bool) float64 {
	ctx := context.Background()
	ctx = client.Add(ctx, "env", env)
	ctx = client.Add(ctx, "load", load)
	ctx = client.Add(ctx, "msg_type", msgType)
	ctx = client.Add(ctx, "throttle", throttle)
	result, err := c.GetFloat(ctx, "example", "example-sample-rate")
	if err == nil {
		return result
	}
	result = getExampleSampleRate(env, load, msgType, throttle)
	if !errors.Is(err, client.ErrNoOpProvider) {
		debug.LogError("Lekko evaluation error", "err", err)
	}
	debug.LogDebug("Lekko fallback", "result", result)
	return result
}

func (c *LekkoClient) GetReturnText(env string) string {
	ctx := context.Background()
	ctx = client.Add(ctx, "env", env)
	result, err := c.GetString(ctx, "example", "return-text")
	if err == nil {
		return result
	}
	result = getReturnText(env)
	if !errors.Is(err, client.ErrNoOpProvider) {
		debug.LogError("Lekko evaluation error", "err", err)
	}
	debug.LogDebug("Lekko fallback", "result", result)
	return result
}
