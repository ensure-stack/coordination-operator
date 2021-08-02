package v1alpha1

// TargetAPI specifis an API to use for operations.
type TargetAPI struct {
	Group   string `json:"group"`
	Version string `json:"version"`
	Kind    string `json:"kind"`
}

// Defines probe parameters to check parts of a package.
type Probe struct {
	// Type of the probe.
	// +kubebuilder:validation:Enum=Condition;FieldsEqual
	Type ProbeType `json:"type"`
	// Condition specific configuration parameters. Only present if Type = Condition.
	Condition   *ProbeConditionSpec   `json:"condition,omitempty"`
	FieldsEqual *ProbeFieldsEqualSpec `json:"fieldsEqual,omitempty"`
}

type ProbeType string

const (
	ProbeCondition   ProbeType = "Condition"
	ProbeFieldsEqual ProbeType = "FieldsEqual"
)

// Condition Probe parameters.
type ProbeConditionSpec struct {
	// Condition Type to probe for.
	Type string `json:"type"`
	// Condition status to probe for.
	// +kubebuilder:default="True"
	Status string `json:"status"`
}

// Compares two fields specified by JSON Paths.
type ProbeFieldsEqualSpec struct {
	FieldA string `json:"fieldA"`
	FieldB string `json:"fieldB"`
}
