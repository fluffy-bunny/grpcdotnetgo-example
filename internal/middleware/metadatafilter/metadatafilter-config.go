package metadatafilter

import (
	grpcContracts_helloworld "github.com/fluffy-bunny/grpcdotnetgo-example/internal/grpcContracts/helloworld"
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
	coreServicesMetadatafilter "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/metadatafilter"
)

func BuildMetadataFilterConfig() map[string]*hashset.StringSet {
	mdFilterBuilder := coreServicesMetadatafilter.NewEntryPointAllowedMetadataMapBuilder()
	mdFilterBuilder.WithAllowedMetadataHeader(grpcContracts_helloworld.FMN_Greeter_SayHello, "header1")
	return mdFilterBuilder.EntryPointAllowedMetadataMap
}
