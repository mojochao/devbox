// Package config contains global config state used by other packages.
package config

// DryRun indicates if commands are to be shown only and not executed.
var DryRun bool

// StateFile is the location of the boxes file to use.
var StateFile string

// Verbose indicates verbose output should be shown.
var Verbose bool
