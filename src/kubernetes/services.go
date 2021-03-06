package kubernetes

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetService will get the service object by the service name
func (kc *KubeCtl) GetService(name string, namespace string) (*corev1.Service, error) {
	return kc.Clientset.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})
}

// GetServices will get all services from the k8s cluster
func (kc *KubeCtl) GetServices(namespace string) ([]*corev1.Service, error) {
	services := []*corev1.Service{}
	servicesList, err := kc.Clientset.CoreV1().Services(namespace).List(metav1.ListOptions{})
	if err != nil {
		return services, err
	}
	for _, s := range servicesList.Items {
		services = append(services, &s)
	}
	return services, nil
}

// CreateService will create the service by the service object
func (kc *KubeCtl) CreateService(service *corev1.Service, namespace string) (*corev1.Service, error) {
	return kc.Clientset.CoreV1().Services(namespace).Create(service)
}

// DeleteService will delete the service by the service name
func (kc *KubeCtl) DeleteService(name string, namespace string) error {
	return kc.Clientset.CoreV1().Services(namespace).Delete(name, &metav1.DeleteOptions{})
}
