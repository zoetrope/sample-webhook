package hooks

import (
	"context"
	"net/http"

	appsv1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var depLogger = logf.Log.WithName("deployment-webhook")

// +kubebuilder:webhook:path=/validate-apps-v1-deployment,mutating=false,failurePolicy=fail,sideEffects=None,groups=apps,resources=deployments;deployments/scale,verbs=create;update,versions=v1,name=vdeployment.kb.io,admissionReviewVersions=v1

type deploymentValidator struct {
	decoder *admission.Decoder
}

// NewDeploymentValidator creates a webhook handler for Deployment.
func NewDeploymentValidator() http.Handler {
	return &webhook.Admission{Handler: &deploymentValidator{}}
}

var _ admission.DecoderInjector = &deploymentValidator{}

func (v *deploymentValidator) InjectDecoder(d *admission.Decoder) error {
	v.decoder = d
	return nil
}

func (v *deploymentValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	depLogger.Info("DeploymentHandler", "req", req)

	var replicas int32
	gk := schema.GroupKind{Group: req.Kind.Group, Kind: req.Kind.Kind}
	switch gk {
	case schema.GroupKind{Group: "apps", Kind: "Deployment"}:
		dep := &appsv1.Deployment{}
		err := v.decoder.Decode(req, dep)
		if err != nil {
			return admission.Errored(http.StatusBadRequest, err)
		}
		if dep.Spec.Replicas != nil {
			replicas = *dep.Spec.Replicas
		}
	case schema.GroupKind{Group: "autoscaling", Kind: "Scale"}:
		scale := &autoscalingv1.Scale{}
		err := v.decoder.Decode(req, scale)
		if err != nil {
			return admission.Errored(http.StatusBadRequest, err)
		}
		replicas = scale.Spec.Replicas
	}

	if replicas < 0 || replicas > 10 {
		return admission.Denied(field.Forbidden(field.NewPath("spec", "replicas"), "replicas should be between 0 and 10").Error())
	}

	return admission.Allowed("ok")
}
