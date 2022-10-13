package memberlist

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/os-observability/tempo-operator/apis/tempo/v1alpha1"
	"github.com/os-observability/tempo-operator/internal/manifests/manifestutils"
)

const (
	PortMemberlist = 7946
	componentName  = "gossip-ring"
)

var (
	GossipSelector = map[string]string{"tempo-gossip-member": "true"}
)

// BuildGossip creates Kubernetes objects that are needed for memberlist.
func BuildGossip(tempo v1alpha1.Microservices) *corev1.Service {
	labels := manifestutils.ComponentLabels(componentName, tempo.Name)
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      manifestutils.Name(componentName, tempo.Name),
			Namespace: tempo.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Selector: GossipSelector,
			Ports: []corev1.ServicePort{
				{
					Name:       componentName,
					Protocol:   corev1.ProtocolTCP,
					Port:       PortMemberlist,
					TargetPort: intstr.FromString("http-memberlist"),
				},
			},
		},
	}
}