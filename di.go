package erru

import "github.com/clavoie/di/v2"

// NewDiDefs returns the dependency definitions for this package.
//
// - Multiplexer is set to resolve per dependency
func NewDiDefs() []*di.Def {
	return []*di.Def{
		{NewMultiplexer, di.PerDependency},
		{NewImpl, di.Singleton},
	}
}
