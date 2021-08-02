package controllers

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/selection"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func relabel(
	ctx context.Context, c client.Client,
	adoption client.Object, specLabels map[string]string,
	objListType *unstructured.UnstructuredList,
	gvk schema.GroupVersionKind,
) error {
	// Build selector
	var requirements []labels.Requirement
	for k := range specLabels {
		requirement, err := labels.NewRequirement(
			k, selection.DoesNotExist, nil)
		if err != nil {
			return fmt.Errorf("building selector: %w", err)
		}
		requirements = append(requirements, *requirement)
	}
	selector := labels.NewSelector().Add(requirements...)

	// List all the things!
	if err := c.List(
		ctx, objListType,
		// can also set this for ClusterAdoption without issue,
		// because the namespace will be ignored for Cluster scoped objects.
		client.InNamespace(adoption.GetNamespace()),
		client.MatchingLabelsSelector{
			Selector: selector,
		},
	); err != nil {
		return fmt.Errorf("listing %s: %w", gvk, err)
	}

	for _, obj := range objListType.Items {
		labels := obj.GetLabels()
		if labels == nil {
			labels = map[string]string{}
		}
		for k, v := range specLabels {
			labels[k] = v
		}
		obj.SetLabels(labels)

		if err := c.Update(ctx, &obj); err != nil {
			return fmt.Errorf("setting labels: %w", err)
		}
	}

	return nil
}
