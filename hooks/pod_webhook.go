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
var podlog = logf.Log.WithName("pod-resource")

func SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(&corev1.Pod{}).
		WithDefaulter(&PodDefaulter{}).
		WithValidator(&PodValidator{}).
		Complete()
}

//+kubebuilder:webhook:path=/mutate--v1-pod,mutating=true,failurePolicy=fail,sideEffects=None,groups=core,resources=pods,verbs=create;update,versions=v1,name=mpod.kb.io,admissionReviewVersions=v1

type PodDefaulter struct{}

var _ admission.CustomDefaulter = &PodDefaulter{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (*PodDefaulter) Default(ctx context.Context, obj runtime.Object) error {
	pod, ok := obj.(*corev1.Pod)
	if !ok {
		return fmt.Errorf("unknown newObj type %T", obj)
	}
	for i, c := range pod.Spec.Containers {
		if c.ImagePullPolicy != corev1.PullAlways {
			pod.Spec.Containers[i].ImagePullPolicy = corev1.PullAlways
		}
	}
	return nil
}

//+kubebuilder:webhook:path=/validate--v1-pod,mutating=false,failurePolicy=fail,sideEffects=None,groups=core,resources=pods,verbs=create;update,versions=v1,name=vpod.kb.io,admissionReviewVersions=v1

type PodValidator struct{}

var _ admission.CustomValidator = &PodValidator{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (*PodValidator) ValidateCreate(ctx context.Context, obj runtime.Object) error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (*PodValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) error {
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (*PodValidator) ValidateDelete(ctx context.Context, obj runtime.Object) error {
	return nil
}
