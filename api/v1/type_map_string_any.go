package apiv1

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// MapStringAny is a map[string]any that can be loaded from YAML.
type MapStringAny map[string]any

func (dst *MapStringAny) UnmarshalYAML(f func(any) error) error {
	var raw map[any]any
	if err := f(&raw); err != nil {
		return err
	}

	*dst = recursiveToMapStringAny(raw)
	return nil
}

func recursiveToMapStringAny(in map[any]any) MapStringAny {
	if in == nil {
		return nil
	}

	out := make(MapStringAny, len(in))
	for _kin, vin := range in {
		kin := fmt.Sprintf("%v", _kin)
		if mvin, ok := vin.(map[any]any); ok {
			out[kin] = recursiveToMapStringAny(mvin)
		} else {
			out[kin] = vin
		}
	}

	return out
}

var _ yaml.Unmarshaler = &MapStringAny{}
