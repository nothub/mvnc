package mvnc

const metaFile = "maven-metadata.xml"

// The Metadata model as defined by the Apache Maven project:
// https://maven.apache.org/ref/3.9.1/maven-repository-metadata/repository-metadata.html
type Metadata struct {
	// The version of the underlying metadata model.
	// TODO: implement (read XML Attribute of the root (`metadata`) node)
	ModelVersion string `xml:"modelVersion"`

	// The groupId when this directory represents "groupId/artifactId" or "groupId/artifactId/version".
	GroupId string `xml:"groupId"`

	// The artifactId when this directory represents "groupId/artifactId" or "groupId/artifactId/version".
	ArtifactId string `xml:"artifactId"`

	// The base version (ie. ending in -SNAPSHOT) when this directory represents a "groupId/artifactId/version" for a SNAPSHOT.
	Version string `xml:"version"`

	// Versioning information when this directory represents "groupId/artifactId" or "groupId/artifactId/version".
	Versioning Versioning `xml:"versioning"`

	// The set of plugins when this directory represents a "groupId" (deprecated)
	// TODO: implement
	Plugins []Plugin `xml:"plugins>plugin"`
}

// Versioning information for "groupId/artifactId" or "groupId/artifactId/version" SNAPSHOT
type Versioning struct {
	// What the last version added to the directory is, including both releases and snapshots ("groupId/artifactId" directory only)
	Latest string `xml:"latest"`

	// What the last version added to the directory is, for the releases only ("groupId/artifactId" directory only)
	Release string `xml:"release"`

	// Versions available of the artifact (both releases and snapshots) ("groupId/artifactId" directory only)
	Versions []string `xml:"versions>version"`

	// When the metadata was last updated (both "groupId/artifactId" and "groupId/artifactId/version" directories).
	// The timestamp is expressed using UTC in the format yyyyMMddHHmmss.
	// TODO: use Time type
	LastUpdated string `xml:"lastUpdated"`

	// The current snapshot data in use for this version ("groupId/artifactId/version" only)
	Snapshot Snapshot `xml:"snapshot"`

	// Information for each sub-artifact available in this artifact snapshot.
	// This is only the most recent SNAPSHOT for each unique extension/classifier combination.
	SnapshotVersions []SnapshotVersion `xml:"snapshotVersions"`
}

// Snapshot data for the last artifact corresponding to the SNAPSHOT base version
type Snapshot struct {
	// The timestamp when this version was deployed.
	// The timestamp is expressed using UTC in the format yyyyMMdd.HHmmss.
	// TODO: use Time type
	Timestamp string `xml:"timestamp"`

	// The incremental build number
	// (Default value is: 0)
	BuildNumber int `xml:"buildNumber"`

	// Whether to use a local copy instead (with filename that includes the base version)
	// (Default value is: false)
	LocalCopy bool `xml:"localCopy"`
}

// Versioning information for a sub-artifact of the current snapshot artifact.
type SnapshotVersion struct {
	// The classifier of the sub-artifact. Each classifier and extension pair may only appear once.
	// Default value is: ""
	Classifier string `xml:"classifier"`

	// The file extension of the sub-artifact. Each classifier and extension pair may only appear once.
	Extension string `xml:"extension"`

	// The resolved snapshot version of the sub-artifact.
	Value string `xml:"value"`

	// The timestamp when this version information was last updated. The timestamp is expressed using UTC in the format yyyyMMddHHmmss.
	Updated string `xml:"updated"`
}

// Mapping information for a single plugin within this group (deprecated).
type Plugin struct {
	// Display name for the plugin.
	Name string `xml:"name"`

	// The plugin invocation prefix (i.e. eclipse for eclipse:eclipse)
	Prefix string `xml:"prefix"`

	// The plugin artifactId
	ArtifactId string `xml:"artifactId"`
}
