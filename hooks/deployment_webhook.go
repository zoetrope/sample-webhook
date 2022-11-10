package hooks

import (
	"context"
	"net/http"

	autoscalingv1 "k8s.io/api/autoscaling/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	appsv1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var logger = logf.Log.WithName("deployment-webhook")

// +kubebuilder:webhook:path=/validate-apps-v1-deployment,mutating=false,failurePolicy=fail,sideEffects=None,groups=apps,resources=deployments;deployments/scale,verbs=create;update,versions=v1,name=vdeployment.kb.io,admissionReviewVersions=v1

type deploymentValidator struct {
	client  client.Client
	decoder *admission.Decoder
}

// NewDeploymentValidator creates a webhook handler for Deployment.
func NewDeploymentValidator(c client.Client, dec *admission.Decoder) http.Handler {
	v := &deploymentValidator{
		client:  c,
		decoder: dec,
	}
	return &webhook.Admission{Handler: v}
}

func (v *deploymentValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	logger.Info("DeploymentHandler", "req", req)

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
		return admission.Denied("invalid")
	}

	return admission.Allowed("ok")
}
