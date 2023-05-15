package mavengo

import (
	"testing"
)

func TestFetchMetadata(t *testing.T) {
	baseUrl := "https://maven.quiltmc.org/repository/release"
	groupId := "org.quiltmc"
	artifactId := "quilt-installer"

	m, err := FetchMetadata(baseUrl, groupId, artifactId)
	if err != nil {
		t.Fatal(err.Error())
	}

	if m.ArtifactId != "quilt-installer" {
		t.Fatal("wrong artifact id")
	}

	t.Logf("%++v", m)
}
