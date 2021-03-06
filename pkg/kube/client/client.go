package client

import (
	"context"

	"github.com/mongodb/mongodb-kubernetes-operator/pkg/kube/configmap"
	"github.com/mongodb/mongodb-kubernetes-operator/pkg/kube/secret"
	"github.com/mongodb/mongodb-kubernetes-operator/pkg/kube/service"
	"github.com/mongodb/mongodb-kubernetes-operator/pkg/kube/statefulset"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	k8sClient "sigs.k8s.io/controller-runtime/pkg/client"
)

func NewClient(c k8sClient.Client) Client {
	return client{
		Client: c,
	}
}

type Client interface {
	k8sClient.Client
	// TODO: remove this function, add mongodb package which has GetAndUpdate function
	GetAndUpdate(nsName types.NamespacedName, obj runtime.Object, updateFunc func()) error
	configmap.GetUpdateCreator
	service.GetUpdateCreator
	secret.GetUpdateCreateDeleter
	statefulset.GetUpdateCreateDeleter
}

type client struct {
	k8sClient.Client
}

// GetAndUpdate fetches the most recent version of the runtime.Object with the provided
// nsName and applies the update function. The update function should update "obj" from
// an outer scope
func (c client) GetAndUpdate(nsName types.NamespacedName, obj runtime.Object, updateFunc func()) error {
	err := c.Get(context.TODO(), nsName, obj)
	if err != nil {
		return err
	}
	// apply the function on the most recent version of the resource
	updateFunc()
	return c.Update(context.TODO(), obj)
}

// GetConfigMap provides a thin wrapper and client.client to access corev1.ConfigMap types
func (c client) GetConfigMap(objectKey k8sClient.ObjectKey) (corev1.ConfigMap, error) {
	cm := corev1.ConfigMap{}
	if err := c.Get(context.TODO(), objectKey, &cm); err != nil {
		return corev1.ConfigMap{}, err
	}
	return cm, nil
}

// UpdateConfigMap provides a thin wrapper and client.Client to update corev1.ConfigMap types
func (c client) UpdateConfigMap(cm corev1.ConfigMap) error {
	return c.Update(context.TODO(), &cm)
}

// CreateConfigMap provides a thin wrapper and client.Client to create corev1.ConfigMap types
func (c client) CreateConfigMap(cm corev1.ConfigMap) error {
	return c.Create(context.TODO(), &cm)
}

// GetSecret provides a thin wrapper and client.Client to access corev1.Secret types
func (c client) GetSecret(objectKey k8sClient.ObjectKey) (corev1.Secret, error) {
	s := corev1.Secret{}
	if err := c.Get(context.TODO(), objectKey, &s); err != nil {
		return corev1.Secret{}, err
	}
	return s, nil
}

// UpdateSecret provides a thin wrapper and client.Client to update corev1.Secret types
func (c client) UpdateSecret(secret corev1.Secret) error {
	return c.Update(context.TODO(), &secret)
}

// CreateSecret provides a thin wrapper and client.Client to create corev1.Secret types
func (c client) CreateSecret(secret corev1.Secret) error {
	return c.Create(context.TODO(), &secret)
}

// DeleteSecret provides a thin wrapper and client.Client to delete corev1.Secret types
func (c client) DeleteSecret(key k8sClient.ObjectKey) error {
	s := corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      key.Name,
			Namespace: key.Namespace,
		},
	}
	return c.Delete(context.TODO(), &s)
}

// GetService provides a thin wrapper and client.Client to access corev1.Service types
func (c client) GetService(objectKey k8sClient.ObjectKey) (corev1.Service, error) {
	s := corev1.Service{}
	if err := c.Get(context.TODO(), objectKey, &s); err != nil {
		return corev1.Service{}, err
	}
	return s, nil
}

// UpdateService provides a thin wrapper and client.Client to update corev1.Service types
func (c client) UpdateService(service corev1.Service) error {
	return c.Update(context.TODO(), &service)
}

// CreateService provides a thin wrapper and client.Client to create corev1.Service types
func (c client) CreateService(service corev1.Service) error {
	return c.Create(context.TODO(), &service)
}

// GetStatefulSet provides a thin wrapper and client.Client to access appsv1.StatefulSet types
func (c client) GetStatefulSet(objectKey k8sClient.ObjectKey) (appsv1.StatefulSet, error) {
	sts := appsv1.StatefulSet{}
	if err := c.Get(context.TODO(), objectKey, &sts); err != nil {
		return appsv1.StatefulSet{}, err
	}
	return sts, nil
}

// UpdateStatefulSet provides a thin wrapper and client.Client to update appsv1.StatefulSet types
func (c client) UpdateStatefulSet(sts appsv1.StatefulSet) error {
	return c.Update(context.TODO(), &sts)
}

// CreateStatefulSet provides a thin wrapper and client.Client to create appsv1.StatefulSet types
func (c client) CreateStatefulSet(sts appsv1.StatefulSet) error {
	return c.Create(context.TODO(), &sts)
}

// DeleteStatefulSet provides a thin wrapper and client.Client to delete appsv1.StatefulSet types
func (c client) DeleteStatefulSet(objectKey k8sClient.ObjectKey) error {
	sts := appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      objectKey.Name,
			Namespace: objectKey.Namespace,
		},
	}
	return c.Delete(context.TODO(), &sts)
}
