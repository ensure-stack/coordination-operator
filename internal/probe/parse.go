package probe

import (
	coordinationv1alpha1 "github.com/ensure-stack/coordination-operator/apis/coordination/v1alpha1"
)

func Parse(probeSpecs []coordinationv1alpha1.Probe) Interface {
	var probeList ProbeList
	for _, probeSpec := range probeSpecs {
		// main probe type
		var probe Interface
		switch probeSpec.Type {
		case coordinationv1alpha1.ProbeCondition:
			if probeSpec.Condition == nil {
				continue
			}

			probe = &ConditionProbe{
				Type:   probeSpec.Condition.Type,
				Status: probeSpec.Condition.Status,
			}

		case coordinationv1alpha1.ProbeFieldsEqual:
			if probeSpec.FieldsEqual == nil {
				continue
			}

			probe = &FieldsEqualProbe{
				FieldA: probeSpec.FieldsEqual.FieldA,
				FieldB: probeSpec.FieldsEqual.FieldB,
			}

		default:
			// Unknown probe type
			panic("unknown probe type")
			// continue
		}

		probeList = append(probeList, probe)
	}

	return probeList
}
