package erru

import "github.com/clavoie/di"

// NewDiDefs returns the dependency definitions for this package.
//
// - Multiplexer is set to resolve per dependency
func NewDiDefs() []*di.Def {
	return []*di.Def{
		&di.Def{NewMultiplexer, di.PerDependency},
		&di.Def{NewImpl, di.Singleton},
	}
}
