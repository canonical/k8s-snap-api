package apiv1

// RevokeKubernetesAuthTokenRPC is the path for the RevokeKubernetesAuthToken RPC.
const RevokeKubernetesAuthTokenRPC = "kubernetes/auth/tokens"

// RevokeKubernetesAuthTokenRequest is the request message for the RevokeKubernetesAuthToken RPC.
type RevokeKubernetesAuthTokenRequest struct {
	Token string `json:"token"`
}

// RevokeKubernetesAuthTokenResponse is the response message for the RevokeKubernetesAuthToken RPC.
type RevokeKubernetesAuthTokenResponse struct{}
