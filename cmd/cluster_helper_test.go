package cmd




import (
	"fmt"
	"testing"
	"os"
	"strings"
	"path/filepath"
)

const (
	testImage string = "library/hello-world:latest"
	testPath string = "$HOME/sandbox/%s.tmp"
)

func TestPullKrakenContainerImage(t *testing.T) {
	// this may be flaky, should ignore if found to be the case too often.
	cli, ctx, err := pullKrakenContainerImage(testImage)

	if cli == nil {
		t.Errorf("Found client object to be nil")
	}

	if ctx == nil {
		t.Errorf("Found background context object to be nil")
	}

	if err != nil {
		t.Errorf("Expected to pull image", testImage, "without error: ", err)
	}

	// clean up and remove image
	removeImage(ctx, cli, testImage)
}

func TestPreRunGetClusterConfig(t *testing.T) {
	home := os.Getenv("HOME")
	ClusterConfigPath = fmt.Sprintf("%s/sandbox/%s.yaml", home, randStringBytesMaskImprSrc(randConfigNameLenght))

	// do not generate path first
	// because we are statically checking for the error string, it may change
	notExistErrorMsg := fmt.Sprintf("$HOME","file %s does not exist", ClusterConfigPath)
	err := preRunGetClusterConfig(clusterCmd, []string{})
	if strings.HasPrefix(err.Error(), notExistErrorMsg) {
		t.Errorf("Expected to get (", notExistErrorMsg, ") but got (", err.Error(), ")")
	}

	// try to generate path now
	_, err = generateFile(ClusterConfigPath)
	if err != nil {
		t.Errorf("Could not generate the test file.")
	}
	defer os.Remove(ClusterConfigPath)

	err = preRunGetClusterConfig(clusterCmd, []string{})
	if err != nil {
		t.Errorf("Unexpected error in initializing cluster config: ", err)
	}
}

func TestGetFirstClusterName(t *testing.T) {
	home := os.Getenv("HOME")
	ClusterConfigPath = fmt.Sprintf("%s/sandbox/%s.yaml", home, randStringBytesMaskImprSrc(randConfigNameLenght))

	_, err := generateFile(ClusterConfigPath)
	if err != nil {
		t.Errorf("Could not generate the test file.")
	}
	defer os.Remove(ClusterConfigPath)

	err = preRunGetClusterConfig(clusterCmd, []string{})
	if err != nil {
		t.Errorf("Unexpected error in initializing cluster config: ", err)
	}

	name := getFirstClusterName()

	if name != "cluster-name-missing" {
		t.Errorf("Expected to get cluster name (","cluster-name-missing", ") but found", name )
	}
}



func generateFile(path string) (*os.File, error) {
	// create directory structure if not exist
	err := os.MkdirAll(filepath.Dir(path), 0777)
	if err != nil {
		return nil, err
	}

	// create path
	return os.Create(path)
}

