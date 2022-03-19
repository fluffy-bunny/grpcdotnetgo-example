package entrypoint_authorization

import (
	proto_helloworld "github.com/fluffy-bunny/grpcdotnetgo-example/internal/grpcContracts/helloworld"
	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
	core_services_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/claimsprincipal"
)

// BuildGrpcEntrypointPermissionsClaimsMap ...
func BuildGrpcEntrypointPermissionsClaimsMap() map[string]*middleware_oidc.EntryPointConfig {
	entryPointClaimsBuilder := core_services_claimsprincipal.NewEntryPointClaimsBuilder()

	// HEALTH SERVICE START
	//---------------------------------------------------------------------------------------------------
	// health check is open to anyone
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen("/grpc.health.v1.Health/Check")

	// GREETER SERVICE START
	//---------------------------------------------------------------------------------------------------
	// FMN_Greeter_SayHello requires permission "read" or "read.write"
	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimFactsMapOR(proto_helloworld.FMN_Greeter_SayHello,
		core_services_claimsprincipal.NewClaimFactTypeAndValue("permission", "read"),
		core_services_claimsprincipal.NewClaimFactTypeAndValue("permission", "read.write"),
	)

	entryPointClaimsBuilder.WithGrpcEntrypointPermissionsClaimsMapOpen(proto_helloworld.FMN_Greeter2_SayHello)

	return entryPointClaimsBuilder.GrpcEntrypointClaimsMap
}
