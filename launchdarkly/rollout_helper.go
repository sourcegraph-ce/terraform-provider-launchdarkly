package launchdarkly

import (
	log "github.com/sourcegraph-ce/logrus"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	ldapi "github.com/launchdarkly/api-client-go"
)

func rolloutSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeInt,
			ValidateFunc: validation.IntBetween(0, 100000),
		},
	}
}

func rolloutFromResourceData(val interface{}) *ldapi.Rollout {
	rolloutList := val.([]interface{})
	variations := []ldapi.WeightedVariation{}
	for idx, k := range rolloutList {
		weight := k.(int)
		variations = append(variations,
			ldapi.WeightedVariation{
				Variation: int32(idx),
				Weight:    int32(weight),
			})
	}

	r := ldapi.Rollout{
		Variations: variations,
	}
	log.Printf("[DEBUG] %+v\n", r)

	return &r
}

func rolloutsToResourceData(rollouts *ldapi.Rollout) interface{} {
	transformed := make([]interface{}, len(rollouts.Variations))

	for _, r := range rollouts.Variations {
		transformed[r.Variation] = r.Weight
	}
	return transformed
}
