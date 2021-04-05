// Package build provides build identity set at compile time by govvv.
// See https://github.com/ahmetb/govvv for details.
package build

var (
	// Date is the  build date.
	Date string

	// GitBranch is the build git branch.
	GitBranch string

	// GitCommit is the build git commit short hash.
	GitCommit string

	// GitState is the build git state.
	GitState string

	// GitSummary is the build git summary.
	GitSummary string

	// Version is the build version.
	Version string
)
