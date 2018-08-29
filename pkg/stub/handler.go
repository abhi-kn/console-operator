package stub

import (
	"context"

	"github.com/openshift/console-operator/pkg/apis/console/v1alpha1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
	// appsv1 "k8s.io/api/apps/v1"
	// corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/apimachinery/pkg/runtime/schema"
	"github.com/openshift/console-operator/pkg/console"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.Console:
		// Vault version has some vault.Reconcile function:
		// https://github.com/operator-framework/operator-sdk-samples/blob/master/vault-operator/pkg/stub/handler.go#L22
		// this is probably a good idea!
		// err := sdk.Create(newbusyBoxPod(o))
		err := console.Reconcile(o)
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("failed to create busybox pod : %v", err)
			return err
		}
	}
	return nil
}

// newbusyBoxPod demonstrates how to create a busybox pod
//func newbusyBoxPod(cr *v1alpha1.Console) *corev1.Pod {
//	labels := map[string]string{
//		"app": "busy-box",
//	}
//	return &corev1.Pod{
//		TypeMeta: metav1.TypeMeta{
//			Kind:       "Pod",
//			APIVersion: "v1",
//		},
//		ObjectMeta: metav1.ObjectMeta{
//			Name:      "busy-box",
//			Namespace: cr.Namespace,
//			OwnerReferences: []metav1.OwnerReference{
//				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
//					Group:   v1alpha1.SchemeGroupVersion.Group,
//					Version: v1alpha1.SchemeGroupVersion.Version,
//					Kind:    "Console",
//				}),
//			},
//			Labels: labels,
//		},
//		Spec: corev1.PodSpec{
//			Containers: []corev1.Container{
//				{
//					Name:    "busybox",
//					Image:   "busybox",
//					Command: []string{"sleep", "3600"},
//				},
//			},
//		},
//	}
//}