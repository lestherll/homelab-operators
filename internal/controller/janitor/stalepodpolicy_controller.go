package janitor

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	janitorv1alpha1 "github.com/lestherll/homelab-operators/api/janitor/v1alpha1"
)

// StalePodPolicyReconciler reconciles a StalePodPolicy object.
type StalePodPolicyReconciler struct {
	client.Client
}

// +kubebuilder:rbac:groups=janitor.homelab,resources=stalepodpolicies,verbs=get;list;watch
// +kubebuilder:rbac:groups=janitor.homelab,resources=stalepodpolicies/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="",resources=pods,verbs=get;list;watch;delete

func (r *StalePodPolicyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

// SetupWithManager registers the controller with the manager.
func (r *StalePodPolicyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.Client = mgr.GetClient()
	return ctrl.NewControllerManagedBy(mgr).
		For(&janitorv1alpha1.StalePodPolicy{}).
		Complete(r)
}
