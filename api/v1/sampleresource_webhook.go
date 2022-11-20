package v1

import (
	"strings"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/utils/pointer"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

var logger = logf.Log.WithName("sampleresource")

func (r *SampleResource) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-sample-zoetrope-github-io-v1-sampleresource,mutating=true,failurePolicy=fail,sideEffects=None,groups=sample.zoetrope.github.io,resources=sampleresources,verbs=create;update,versions=v1,name=msampleresource.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &SampleResource{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *SampleResource) Default() {
	logger.Info("default", "name", r.Name)

	if r.Spec.Replicas == nil {
		r.Spec.Replicas = pointer.Int32(1)
	}
}

//+kubebuilder:webhook:path=/validate-sample-zoetrope-github-io-v1-sampleresource,mutating=false,failurePolicy=fail,sideEffects=None,groups=sample.zoetrope.github.io,resources=sampleresources,verbs=create;update,versions=v1,name=vsampleresource.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &SampleResource{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *SampleResource) ValidateCreate() error {
	logger.Info("validate create", "name", r.Name)

	return r.validate()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *SampleResource) ValidateUpdate(old runtime.Object) error {
	logger.Info("validate update", "name", r.Name)

	return r.validate()
}

func (r *SampleResource) validate() error {
	var errs field.ErrorList

	if r.Spec.Replicas == nil {
		errs = append(errs, field.Invalid(field.NewPath("spec", "replicas"), r.Spec.Replicas, "replicas cannot be empty"))
	} else if *r.Spec.Replicas < 1 {
		errs = append(errs, field.Invalid(field.NewPath("spec", "replicas"), r.Spec.Replicas, "replicas should be grater than 0"))
	}

	if len(r.Spec.Image) == 0 {
		errs = append(errs, field.Invalid(field.NewPath("spec", "image"), r.Spec.Image, "image cannot be empty"))
	} else if !strings.Contains(r.Spec.Image, ":") {
		errs = append(errs, field.Invalid(field.NewPath("spec", "image"), r.Spec.Image, "image should have a tag"))
	} else {
		images := strings.Split(r.Spec.Image, ":")
		if len(images) != 2 {
			errs = append(errs, field.Invalid(field.NewPath("spec", "image"), r.Spec.Image, "image is not valid format"))
		} else if images[1] == "latest" {
			errs = append(errs, field.Invalid(field.NewPath("spec", "image"), r.Spec.Image, "image cannot have latest tag"))
		}
	}

	if len(errs) > 0 {
		err := apierrors.NewInvalid(schema.GroupKind{Group: GroupVersion.Group, Kind: "Sample"}, r.Name, errs)
		return err
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *SampleResource) ValidateDelete() error {
	logger.Info("validate delete", "name", r.Name)

	return nil
}
