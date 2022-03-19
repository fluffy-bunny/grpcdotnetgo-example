package server

import (
	coreContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/core"
	pluginContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/plugin"
)

type pluginServiceWithTestStartup struct {
	Startup     coreContracts.IStartup
	testStartup ITestStartup
}

func NewPluginsWithTestStartup(testStartup ITestStartup) []pluginContracts.IGRPCDotNetGoPlugin {
	return []pluginContracts.IGRPCDotNetGoPlugin{
		NewPluginWithTestStartup(testStartup),
	}
}

// NewPluginWithClaimsMap ...
func NewPluginWithTestStartup(testStartup ITestStartup) pluginContracts.IGRPCDotNetGoPlugin {
	return &pluginServiceWithTestStartup{
		testStartup: testStartup,
	}
}

// GetName of the plugin
func (p *pluginServiceWithTestStartup) GetName() string {
	return "grpcdotnetgo-example-micro"
}

// GetStartup gets the IStartup object
func (p *pluginServiceWithTestStartup) GetStartup() coreContracts.IStartup {
	if p.Startup == nil {
		p.Startup = NewTestStartup(p.testStartup)
	}
	return p.Startup
}
