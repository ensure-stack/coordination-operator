package controllers

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	coordinationv1alpha1 "github.com/ensure-stack/coordination-operator/apis/coordination/v1alpha1"
	internalprobe "github.com/ensure-stack/coordination-operator/internal/probe"
)

// Process objects chosen to be handed over.
func handleProcessing(
	ctx context.Context,
	c client.Client,
	log logr.Logger,
	handoverStrategy coordinationv1alpha1.HandoverStrategy,
	objType *unstructured.Unstructured,
	probe internalprobe.Interface,
	processing []coordinationv1alpha1.HandoverRef,
) (stillProcessing []coordinationv1alpha1.HandoverRef, err error) {
	for _, processing := range processing {
		processingObj := objType.DeepCopy()
		err := c.Get(ctx, types.NamespacedName{
			Name:      processing.Name,
			Namespace: processing.Namespace,
		}, processingObj)
		if errors.IsNotFound(err) {
			// Object gone, remove it from processing queue.
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("getting object in process queue: %w", err)
		}

		labels := processingObj.GetLabels()
		if labels == nil ||
			labels[handoverStrategy.Relabel.LabelKey] != handoverStrategy.Relabel.ToValue {
			labels[handoverStrategy.Relabel.LabelKey] = handoverStrategy.Relabel.ToValue
			processingObj.SetLabels(labels)
			if err := c.Update(ctx, processingObj); err != nil {
				return nil, fmt.Errorf("updating object in process queue: %w", err)
			}
		}

		statusValue, ok, err := getValueFromJsonPath(processingObj, handoverStrategy.Relabel.StatusPath)
		if err != nil {
			return nil, fmt.Errorf("getting status value: %w", err)
		}
		if !ok || statusValue != handoverStrategy.Relabel.ToValue {
			log.Info("waiting for status field to update", "objName", processing.Name)
			stillProcessing = append(stillProcessing, processing)
			break
		}

		if success, message := probe.Probe(processingObj); !success {
			log.Info("waiting to be ready", "objName", processing.Name, "failure", message)
			stillProcessing = append(stillProcessing, processing)
		}
	}
	return stillProcessing, nil
}
