package authproviders

import (
	"context"
	"net/http"
)

// BackendFactory is responsible for creating Backends.
type BackendFactory interface {
	// CreateBackend creates a new backend instance for the given auth provider, using the specified configuration.
	CreateBackend(ctx context.Context, id string, uiEndpoints []string, config map[string]string) (Backend, error)

	// ProcessHTTPRequest is the dispatcher for HTTP/1.1 requests to `<sso-prefix>/<provider-type>/...`. The envisioned
	// workflow consists of extracting the specific auth provider ID and clientState from the request, usually via a
	// `state` parameter, and returning this provider ID and clientState from the function (with the Registry taking
	// care of forwarding the request to that provider's HTTP handler). If there are any provider-independent HTTP
	// endpoints (such as the SP metadata for SAML), this can be handled in this function as well - an empty
	// provider ID along with a nil error needs to be returned in that case.
	ProcessHTTPRequest(w http.ResponseWriter, r *http.Request) (providerID string, clientState string, err error)

	// ResolveProviderAndClientState extracts the provider ID and client state (if any) from an (opaque) state string.
	ResolveProviderAndClientState(state string) (providerID string, clientState string, err error)

	// RedactConfig returns config or its copy with secrets (if any) replaced with dummy strings.
	// It is called on config loaded from storage to prepare it for display in UI.
	RedactConfig(config map[string]string) map[string]string
	// MergeConfig un-does the effects of RedactConfig. It is called on newCfg submitted by a user.
	// It should restore secrets missing from newCfg (if any) by copying them back from oldCfg (loaded from storage).
	MergeConfig(newCfg, oldCfg map[string]string) map[string]string
}

// BackendFactoryCreator is a function for creating a BackendFactory, given a base URL (excluding trailing slashes) for
// those factory's HTTP handlers.
type BackendFactoryCreator func(urlPathPrefix string) BackendFactory
