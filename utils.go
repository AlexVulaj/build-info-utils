package build_info_utils

import (
	"errors"
	"fmt"
	"runtime/debug"
)

var readBuildInfoFunc = debug.ReadBuildInfo

func GetDependencyVersion(dependencyPath string) (string, error) {
	buildInfo, ok := readBuildInfoFunc()
	if !ok {
		return "", errors.New("failed to parse build info")
	}
	for _, dep := range buildInfo.Deps {
		if dep.Path == dependencyPath {
			return dep.Version, nil
		}
	}

	return "", fmt.Errorf("unable to find version for %v", dependencyPath)
}
