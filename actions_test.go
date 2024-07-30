package mvnc

import (
	"context"
	"log"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestWithReposilite(t *testing.T) {

	t.SkipNow()

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: "dzikoysk/reposilite:3.5.6",
			Cmd: []string{
				"--port", "8080",
				"--token", "test:foobar",
			},
			Env: map[string]string{
				"JAVA_OPTS": "-Xmx64M",
			},
			ExposedPorts: []string{"8080/tcp"},
			WaitingFor:   wait.ForLog("For help, type 'help' or '?'"),
		},
		Started: true,
	}

	ctx := context.Background()
	tco, err := testcontainers.GenericContainer(ctx, req)

	if err != nil {
		log.Fatalf("Unable to start reposilite: %s", err)
	}
	defer func() {
		if err := tco.Terminate(ctx); err != nil {
			log.Fatalf("Unable to stop reposilite: %s", err)
		}
	}()

	_TestList(t)
	_TestUpload(t)
	_TestDownload(t)
	_TestRemove(t)

}

func _TestList(t *testing.T) {
	err := List("x", "y", "z")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func _TestUpload(t *testing.T) {
	err := Upload("x", "y", "z", "LICENSE.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func _TestDownload(t *testing.T) {
	err := Download("x", "y", "z", "LICENSE.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func _TestRemove(t *testing.T) {
	err := Remove("x", "y", "z", "LICENSE.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
}
