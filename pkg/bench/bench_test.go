package bench

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

// We haven't used standard Go benchmark because we need more
//	detailed reporting such as error rate.

func TestBenchAPI(t *testing.T) {
	ctx := context.Background()
	testName := "TestRESTClient"
	host := os.Getenv("BENCH_HOST")
	workerCount, _ := strconv.Atoi(os.Getenv("WORKER_COUNT")) // nolint: errcheck
	loopCount, _ := strconv.Atoi(os.Getenv("LOOP_COUNT"))     // nolint: errcheck
	if host == "" {
		t.Skip("BENCH_HOST isn't set. skipping")
		return
	}
	restClient := client.NewHTTP(
		host,
		host+"/keystone/v3",
		testName,
		testName,
		true,
		keystone.NewScope("", "default", "", testName),
	)
	restClient.InSecure = true
	restClient.Init()
	var err error

	logrus.Info("Benchmark create:")
	Benchmark(workerCount, loopCount, func(w, l int) error {
		// Contact the server and print out its response.
		project := models.MakeProject()
		project.FQName = []string{"default-domain", "project", fmt.Sprintf("%d_%d", w, l)}
		project.ParentType = "domain"
		project.ParentUUID = "beefbeef-beef-beef-beef-beefbeef0002"
		project.ConfigurationVersion = 1
		_, err = restClient.CreateProject(ctx, &services.CreateProjectRequest{
			Project: project,
		})
		return err
	})

	logrus.Info("Benchmark list:")
	Benchmark(workerCount, loopCount, func(w, l int) error {
		_, err := restClient.ListProject(ctx, &services.ListProjectRequest{
			Spec: &baseservices.ListSpec{},
		})
		return err
	})

	//cleanup
	if os.Getenv("CLEAN_UP") != "true" {
		return
	}
	cleanup(ctx, t, restClient)
}

func cleanup(ctx context.Context, t *testing.T, restClient *client.HTTP) {
	//cleanup
	for i := 0; i < 10000; i++ {
		projects, err := restClient.ListProject(ctx, &services.ListProjectRequest{
			Spec: &baseservices.ListSpec{},
		})
		assert.NoError(t, err)
		if len(projects.GetProjects()) <= 1 {
			break
		}
		wg := &sync.WaitGroup{}
		for _, project := range projects.GetProjects() {
			wg.Add(1)
			go func(uuid string) {
				restClient.DeleteProject(ctx, &services.DeleteProjectRequest{ // nolint: errcheck
					ID: uuid,
				})
				wg.Done()
			}(project.UUID)
		}
		wg.Wait()
	}
}
