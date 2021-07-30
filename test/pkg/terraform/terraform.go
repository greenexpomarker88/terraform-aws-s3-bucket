package terraform

import (
	"encoding/json"
	"fmt"
	tfjson "github.com/hashicorp/terraform-json"
)

func NewRepository(planJson string) *Repository {
	var plan *tfjson.Plan
	err := json.Unmarshal([]byte(planJson), &plan)
	if err != nil {
		return nil
	}

	return &Repository{plan: plan}
}

type Repository struct {
	plan *tfjson.Plan
}

type ModuleRepository struct {
	module *tfjson.StateModule
}

func (r *Repository) GetModule(address string) (*ModuleRepository, error) {
	for _, m := range r.plan.PlannedValues.RootModule.ChildModules {
		if m.Address == address {
			return &ModuleRepository{module: m}, nil
		}
	}
	return nil, fmt.Errorf("could not find module with address %v", address)
}

func (mr *ModuleRepository) GetTaggableResources() []*tfjson.StateResource {
	results := findResources(mr.module, func(rsc *tfjson.StateResource) bool {
		_, ok := rsc.AttributeValues["tags"]
		return ok && rsc.Mode != "data"
	})

	return results
}

func (mr *ModuleRepository) GetResource(address string) (*tfjson.StateResource, error) {
	results := findResources(mr.module, func(rsc *tfjson.StateResource) bool {
		return rsc.Address == address && rsc.Mode != "data"
	})

	if len(results) > 1 {
		return nil, fmt.Errorf("found more than one resource with the address %v", address)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("could not find a resource with the address %v", address)
	}

	return results[0], nil
}

func findResources(obj interface{}, f func(r *tfjson.StateResource) bool) []*tfjson.StateResource {
	var resources []*tfjson.StateResource

	switch v := obj.(type) {
	case *tfjson.StateResource:
		if f(v) {
			resources = append(resources, v)
		}
	case *tfjson.StateModule:
		for _, o := range v.Resources {
			if r := findResources(o, f); r != nil {
				resources = append(resources, r...)
			}
		}
		for _, o := range v.ChildModules {
			if r := findResources(o, f); r != nil {
				resources = append(resources, r...)
			}
		}
	case *tfjson.Plan:
		for _, o := range v.PlannedValues.RootModule.Resources {
			if r := findResources(o, f); r != nil {
				resources = append(resources, r...)
			}
		}
		for _, o := range v.PlannedValues.RootModule.ChildModules {
			if r := findResources(o, f); r != nil {
				resources = append(resources, r...)
			}
		}
	}
	return resources
}
