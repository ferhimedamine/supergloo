// Code generated by protoc-gen-solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionRoutingRuleFunc func(original, desired *RoutingRule) (bool, error)

type RoutingRuleReconciler interface {
	Reconcile(namespace string, desiredResources RoutingRuleList, transition TransitionRoutingRuleFunc, opts clients.ListOpts) error
}

func routingRulesToResources(list RoutingRuleList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, routingRule := range list {
		resourceList = append(resourceList, routingRule)
	}
	return resourceList
}

func NewRoutingRuleReconciler(client RoutingRuleClient) RoutingRuleReconciler {
	return &routingRuleReconciler{
		base: reconcile.NewReconciler(client.BaseClient()),
	}
}

type routingRuleReconciler struct {
	base reconcile.Reconciler
}

func (r *routingRuleReconciler) Reconcile(namespace string, desiredResources RoutingRuleList, transition TransitionRoutingRuleFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "routingRule_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*RoutingRule), desired.(*RoutingRule))
		}
	}
	return r.base.Reconcile(namespace, routingRulesToResources(desiredResources), transitionResources, opts)
}
