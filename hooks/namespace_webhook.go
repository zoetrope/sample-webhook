package hooks

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var nsLogger = logf.Log.WithName("namespace-validator")

func SetupNamespaceWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(&corev1.Namespace{}).
		WithValidator(&NamespaceValidator{}).
		Complete()
}

//+kubebuilder:webhook:path=/validate--v1-namespace,mutating=false,failurePolicy=fail,sideEffects=None,groups=core,resources=namespaces,verbs=delete,versions=v1,name=vnamespace.kb.io,admissionReviewVersions=v1

type NamespaceValidator struct{}

var _ admission.CustomValidator = &NamespaceValidator{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (v *NamespaceValidator) ValidateCreate(ctx context.Context, obj runtime.Object) error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (v *NamespaceValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) error {
	return nil
}

const annotationForDelete = "i-am-sure-to-delete"

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (*NamespaceValidator) ValidateDelete(ctx context.Context, obj runtime.Object) error {
	ns, ok := obj.(*corev1.Namespace)
	if !ok {
		return fmt.Errorf("unknown newObj type %T", obj)
	}

	ann := ns.GetAnnotations()
	name := ns.GetName()

	if val, ok := ann[annotationForDelete]; ok && val == name {
		return nil
	}

	return fmt.Errorf(`add "i-am-sure-to-delete: %s" annotation to delete this`, name)
}
