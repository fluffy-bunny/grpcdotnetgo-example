package development

import (
	proto_helloworld "github.com/fluffy-bunny/grpcdotnetgo-example/internal/grpcContracts/helloworld"
	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	coreUtilsTestsAuth "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils/tests/auth"
)

// BuildValidClaimsPrincipalMap ...
// this is for testing.  Instead or our middleware cracking an access_token, we simply create a custom
// claims principal for each endpoint.  The auth middleware does the same thing, but it cracks an access_token
// and turns that into a claims principal.
func BuildValidClaimsPrincipalMap() coreUtilsTestsAuth.EntryPointToClaimsMap {
	claimsMap := make(coreUtilsTestsAuth.EntryPointToClaimsMap)
	claimsMap[proto_helloworld.FMN_Greeter_SayHello] = []claimsprincipalContracts.Claim{
		{
			Type:  "permission",
			Value: "read",
		}, {
			Type:  "permission",
			Value: "read.write",
		},
	}
	return claimsMap
}
