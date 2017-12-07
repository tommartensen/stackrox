package benchmarks

import (
	"fmt"
	"strings"
	"time"

	"bitbucket.org/stack-rox/apollo/docker-bench/utils"
	"bitbucket.org/stack-rox/apollo/pkg/api/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/docker"
	"bitbucket.org/stack-rox/apollo/pkg/env"
	"bitbucket.org/stack-rox/apollo/pkg/logging"
	"bitbucket.org/stack-rox/apollo/pkg/uuid"
	"github.com/golang/protobuf/ptypes"
	"golang.org/x/net/context"
)

var (
	log = logging.New("benchmark")
)

// RunBenchmark runs a benchmark based on environment variables
func RunBenchmark() *v1.BenchmarkResult {
	hostname, err := getHostname()
	if err != nil {
		log.Fatalf("Could not find this node's hostname: %+v", err)
	}
	protoStartTime := ptypes.TimestampNow()
	checkResults := runBenchmark()
	protoEndTime := ptypes.TimestampNow()
	result := &v1.BenchmarkResult{
		Id:        uuid.NewV4().String(),
		Results:   checkResults,
		StartTime: protoStartTime,
		EndTime:   protoEndTime,
		Host:      hostname,
		ScanId:    env.ScanID.Setting(),
	}
	return result
}

func runBenchmark() []*v1.CheckResult {
	checks := renderChecks()

	results := make([]*v1.CheckResult, 0, len(checks))
	for _, check := range checks {
		definition := check.Definition().CheckDefinition
		for _, dep := range check.Definition().Dependencies {
			if err := dep(); err != nil {
				msg := fmt.Sprintf("Skipping Test %v due to err in dependency: %+v\n", check.Definition().Name, err)
				log.Error(msg)
				result := &v1.CheckResult{
					Definition: &definition,
					Result:     v1.CheckStatus_NOTE,
					Notes:      []string{msg},
				}
				results = append(results, result)
				continue
			}
		}
		result := check.Run()
		result.Definition = &definition
		results = append(results, &result)
	}
	return results
}

func getHostname() (string, error) {
	cli, err := docker.NewClient()
	if err != nil {
		return "", fmt.Errorf("docker client setup: %s", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	info, err := cli.Info(ctx)
	if err != nil {
		return "", fmt.Errorf("docker info: %s", err)
	}
	return info.Name, nil
}

func renderChecks() []utils.Check {
	checkStrs := strings.Split(env.Checks.Setting(), ",")
	var checks []utils.Check
	for _, checkStr := range checkStrs {
		check, ok := Registry[checkStr]
		if !ok {
			log.Errorf("Check %v is not currently supported", checkStr)
			continue
		}
		checks = append(checks, check)
	}
	return checks
}
