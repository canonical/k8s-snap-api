package apiv1

import "time"

// GetJoinTokenRPC is the path for the GetJoinToken RPC.
const GetJoinTokenRPC = "k8sd/cluster/tokens"

// GetJoinTokenRequest is the request message for the GetJoinToken RPC.
type GetJoinTokenRequest struct {
	// Name is the name of the token to generate.
	Name string `json:"name"`
	// Worker should be set to true to generate a token for joining a worker node.
	Worker bool `json:"worker"`
	// TTL is the duration until the token expires (time-to-live).
	TTL time.Duration `json:"ttl,omitempty"`
}

// GetJoinTokenResponse is the response message for the GetJoinToken RPC.
type GetJoinTokenResponse struct {
	// EncodedToken is the generated join token.
	EncodedToken string `json:"token"`
}
