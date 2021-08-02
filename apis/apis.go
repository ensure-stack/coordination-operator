package apis

import (
	"github.com/ensure-stack/coordination-operator/apis/coordination"
	"k8s.io/apimachinery/pkg/runtime"
)

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder = runtime.SchemeBuilder{
	coordination.AddToScheme,
}

// AddToScheme adds all addon Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
