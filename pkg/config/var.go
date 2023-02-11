package config

// will be set during compilation time via -X
var (
	BuiltAt   string
	GoVersion string
	GitAuthor string
	GitCommit string
	Version   string = "dev"
)
