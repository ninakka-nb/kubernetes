package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	gatewayapiv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

type Container struct {
	Name               string                      `json:"name"`
	Image              string                      `json:"image"`
	Args               []string                    `json:"args,omitempty" protobuf:"bytes,4,rep,name=args"`
	ArgsType           string                      `json:"argsType,omitempty"`
	EnvFrom            []corev1.EnvFromSource      `json:"envFrom,omitempty" protobuf:"bytes,19,rep,name=envFrom"`
	Env                []corev1.EnvVar             `json:"env,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,7,rep,name=env"`
	EnvType            string                      `json:"envType,omitempty"`
	Resources          corev1.ResourceRequirements `json:"resources,omitempty"`
	LivenessProbeType  string                      `json:"livenessProbeType,omitempty"`
	LivenessProbe      *corev1.Probe               `json:"livenessProbe,omitempty"`
	ReadinessProbeType string                      `json:"readinessProbeType,omitempty"`
	ReadinessProbe     *corev1.Probe               `json:"readinessProbe,omitempty"`
}

type Port struct {
	Type       string `json:"type"`
	PortNumber uint   `json:"port"`
}

type ModuleRef struct {
	Kind       string      `json:"kind"`
	Name       string      `json:"name"`
	Containers []Container `json:"containers,omitempty"`
	Ports      []Port      `json:"ports,omitempty"`
}

type BreakoutHTTPRoutingRuleMatchesSpec struct {
	Matches []gatewayapiv1beta1.HTTPRouteMatch `json:"matches"`
}

type BreakoutHTTPRoutingRuleSpec struct {
	Rules []BreakoutHTTPRoutingRuleMatchesSpec `json:"rules"`
}

type RoutingRuleSpec struct {
	HTTPRoute BreakoutHTTPRoutingRuleSpec `json:"httpRoute"`
}

type BreakoutSpec struct {
	Environment      string          `json:"environment"`
	Feature          string          `json:"feature"`
	BaseEnvironment  string          `json:"baseEnvironment,omitempty"`
	BaseFeature      string          `json:"baseFeature,omitempty"`
	BaseBreakout     string          `json:"baseBreakout,omitempty"`
	BaseNamespace    string          `json:"baseNamespace,omitempty"`
	ReferredBreakout string          `json:"referredBreakout,omitempty"` // Used only for current breakout
	ModuleRefs       []ModuleRef     `json:"moduleRefs,omitempty"`
	RoutingRules     RoutingRuleSpec `json:"routingRules,omitempty"`
}

type BreakoutResourceInstance struct {
	// API version of the resource.
	APIVersion string `json:"apiVersion" protobuf:"bytes,5,opt,name=apiVersion"`

	// Kind of the resource.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind" protobuf:"bytes,1,opt,name=kind"`

	// Name of the resource.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names
	Name string `json:"name" protobuf:"bytes,3,opt,name=name"`

	// GenerateName, an optional prefix, of the resource used by the server, to generate a unique
	// name ONLY IF the Name field has not been provided.
	// If this field is used, the name returned to the client will be different
	// than the name passed. This value will also be combined with a unique suffix.
	// The provided value has the same validation rules as the Name field,
	// and may be truncated by the length of the suffix required to make the value
	// unique on the server.
	//
	// +optional
	GenerateName string `json:"generateName,omitempty" protobuf:"bytes,2,opt,name=generateName"`

	// UID of the resource. TODO: Remove omitempty once we fill in from basenamespace
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids
	UID types.UID `json:"uid,omitempty" protobuf:"bytes,4,opt,name=uid,casttype=k8s.io/apimachinery/pkg/types.UID"`

	// Namespace defines the space within which each name must be unique.
	// Must be a DNS_LABEL.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces
	Namespace string `json:"namespace" protobuf:"bytes,3,opt,name=namespace"`

	// Fully qualified Name along with hash
	NameWithHash string `json:"nameWithHash,omitempty"`

	// List of objects depended by this object.
	OwnerReferences []metav1.OwnerReference `json:"ownerReferences,omitempty" patchStrategy:"merge" patchMergeKey:"uid" protobuf:"bytes,13,rep,name=ownerReferences"`
}

type BreakoutResourceMap struct {
	// Map of resource instances indexed by name
	ResourceMap map[string][]BreakoutResourceInstance `json:"resourceMap"`
}

type BreakoutStatus struct {
	// Map of resource types indexed by group+version, kind
	ResourceTypeMap map[string]map[string]BreakoutResourceMap `json:"resourceTypeMap,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName={bo,bos,brkout,brkouts}
type Breakout struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BreakoutSpec   `json:"spec,omitempty"`
	Status BreakoutStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type BreakoutList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Breakout `json:"items"`
}
