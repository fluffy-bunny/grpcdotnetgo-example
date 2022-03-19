package server

import (
	grpcServerStartup "github.com/fluffy-bunny/grpcdotnetgo-example/internal/startup"
	coreContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
	coreUtilsTests "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils/tests"
	coreUtilsTestsAuth "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils/tests/auth"
	di "github.com/fluffy-bunny/sarulabsdi"
)

type (
	ITestStartup interface {
		GetConfigureServicesHook() func(builder *di.Builder)
	}
	TestStartup struct {
		ConfigureServicesHook func(builder *di.Builder)
	}
)

func (s *TestStartup) GetConfigureServicesHook() func(builder *di.Builder) {
	return s.ConfigureServicesHook
}
func NewTestStartup(testStartup ITestStartup) coreContracts.IStartup {
	config := &coreUtilsTests.TestStartupWrapperConfig{
		InnerStartup:          grpcServerStartup.NewStartup(),
		ConfigureServicesHook: testStartup.GetConfigureServicesHook(),
	}
	wrapper := coreUtilsTests.NewTestStartupWrapper(config)

	return wrapper
}

func BuildEmptyClaimsPrincipalMap() coreUtilsTestsAuth.EntryPointToClaimsMap {
	claimsMap := make(coreUtilsTestsAuth.EntryPointToClaimsMap)
	return claimsMap
}
