package commonschema

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
)

func TagsDataSource() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeMap,
		Computed: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}

func TagsForceNew() *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeMap,
		Optional:     true,
		ForceNew:     true,
		ValidateFunc: ValidateTags,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}

func Tags() *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeMap,
		Optional:     true,
		ValidateFunc: ValidateTags,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}

func TagsWithLowerCaseKeys() *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeMap,
		Optional:     true,
		ValidateFunc: validateLowerCaseKeysForTags,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}

func FlattenTags(input *map[string]string) map[string]*string {
	output := make(map[string]*string)
	if input == nil {
		return output
	}

	for k, v := range *input {
		output[k] = pointer.ToStringPointer(v)
	}

	return output
}

func ExpandTags(input map[string]interface{}) *map[string]string {
	output := make(map[string]string)

	for k, v := range input {
		output[k] = v.(string)
	}

	return &output
}

func ValidateTags(v interface{}, _ string) (warnings []string, errors []error) {
	tagsMap := v.(map[string]interface{})

	if len(tagsMap) > 50 {
		errors = append(errors, fmt.Errorf("a maximum of 50 tags can be applied to each resource"))
	}

	for k, v := range tagsMap {
		if len(k) > 512 {
			errors = append(errors, fmt.Errorf("the maximum length for a tag key is 512 characters: %q is %d characters", k, len(k)))
		}

		value, err := tagValueToString(v)
		if err != nil {
			errors = append(errors, err)
		} else if len(value) > 256 {
			errors = append(errors, fmt.Errorf("the maximum length for a tag value is 256 characters: the value for %q is %d characters", k, len(value)))
		}
	}

	return warnings, errors
}

func tagValueToString(v interface{}) (string, error) {
	switch value := v.(type) {
	case string:
		return value, nil
	case int:
		return fmt.Sprintf("%d", value), nil
	default:
		return "", fmt.Errorf("unknown tag type %T in tag value", value)
	}
}

func validateLowerCaseKeysForTags(i interface{}, k string) (warnings []string, errors []error) {
	tagsMap, ok := i.(map[string]interface{})
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %s to be map", k))
		return warnings, errors
	}

	if len(tagsMap) > 50 {
		errors = append(errors, fmt.Errorf("a maximum of 50 tags can be applied to each ARM resource"))
	}

	for key, value := range tagsMap {
		if len(key) > 512 {
			errors = append(errors, fmt.Errorf("the maximum length for a tag key is 512 characters: %q has %d characters", key, len(key)))
			return warnings, errors
		}

		if strings.ToLower(key) != key {
			errors = append(errors, fmt.Errorf("a tag key %q expected to be all in lowercase", key))
			return warnings, errors
		}

		v, err := tagValueToString(value)
		if err != nil {
			errors = append(errors, err)
			return warnings, errors
		}
		if len(v) > 256 {
			errors = append(errors, fmt.Errorf("the maximum length for a tag value is 256 characters: the value for %q has %d characters", key, len(v)))
			return warnings, errors
		}
	}

	return warnings, errors
}
