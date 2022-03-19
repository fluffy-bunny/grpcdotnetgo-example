package server

import (
	"context"
	"log"
	"testing"

	development "github.com/fluffy-bunny/grpcdotnetgo-example/internal/development"
	contracts_core_backgroundtasks "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/backgroundtasks"
	grpcdotnetgo_core "github.com/fluffy-bunny/grpcdotnetgo/pkg/core"
	mocks_core_backgroundtasks "github.com/fluffy-bunny/grpcdotnetgo/pkg/mocks/backgroundtasks"
	services_auth_inmemory "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/auth/inmemory"
	grpcdotnetgo_testing "github.com/fluffy-bunny/grpcdotnetgo/pkg/testing"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	health "google.golang.org/grpc/health/grpc_health_v1"
	bufconn "google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

func TestHealthCheck(t *testing.T) {
	grpcdotnetgo_testing.RunTest(t, func(ctrl *gomock.Controller) {
		testStartup := &TestStartup{
			ConfigureServicesHook: func(builder *di.Builder) {
				authClaimsMap := development.BuildValidClaimsPrincipalMap()
				// we don't want any real auth going to OIDC endpoints, we are going to produce our own claims principals
				services_auth_inmemory.AddSingletonIModularAuthMiddleware(builder, authClaimsMap)
				mockIJobsProvider := mocks_core_backgroundtasks.NewMockIJobsProvider(ctrl)
				contracts_core_backgroundtasks.AddSingletonIBackgroundTasksByObj(builder, mockIJobsProvider)
			},
		}

		plugins := NewPluginsWithTestStartup(testStartup)
		lis := bufconn.Listen(bufSize)
		myRuntime := grpcdotnetgo_core.NewRuntime()
		future := grpcdotnetgo_testing.ExecuteWithPromiseAsync(myRuntime, lis, plugins)

		ctx := context.Background()
		conn, err := grpcdotnetgo_testing.CreateConnection(ctx, lis)
		if err != nil {
			t.Fatalf("Failed to dial bufnet: %v", err)
		}
		defer conn.Close()

		healthClient := health.NewHealthClient(conn)

		resp, err := healthClient.Check(ctx, &health.HealthCheckRequest{})
		if err != nil {
			t.Fatalf("healthClient.Check failed: %v", err)
		}
		log.Printf("Response: %+v", resp)
		assert.Equal(t, health.HealthCheckResponse_SERVING, resp.Status)
		// Test for output here.
		myRuntime.Stop()
		future.Get()
	})
}
