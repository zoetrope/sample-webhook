package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var sampleresourcelog = logf.Log.WithName("sampleresource-resource")

func (r *SampleResource) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-sample-zoetrope-github-io-v1-sampleresource,mutating=true,failurePolicy=fail,sideEffects=None,groups=sample.zoetrope.github.io,resources=sampleresources,verbs=create;update,versions=v1,name=msampleresource.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &SampleResource{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *SampleResource) Default() {
	sampleresourcelog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-sample-zoetrope-github-io-v1-sampleresource,mutating=false,failurePolicy=fail,sideEffects=None,groups=sample.zoetrope.github.io,resources=sampleresources,verbs=create;update,versions=v1,name=vsampleresource.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &SampleResource{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *SampleResource) ValidateCreate() error {
	sampleresourcelog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *SampleResource) ValidateUpdate(old runtime.Object) error {
	sampleresourcelog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *SampleResource) ValidateDelete() error {
	sampleresourcelog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
