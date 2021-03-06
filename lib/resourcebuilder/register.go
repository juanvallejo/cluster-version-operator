package resourcebuilder

import (
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	rbacv1beta1 "k8s.io/api/rbac/v1beta1"
	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

func init() {
	rm := NewResourceMapper()
	rm.RegisterGVK(apiextv1beta1.SchemeGroupVersion.WithKind("CustomResourceDefinition"), newCRDBuilder)
	rm.RegisterGVK(appsv1.SchemeGroupVersion.WithKind("Deployment"), newDeploymentBuilder)
	rm.RegisterGVK(batchv1.SchemeGroupVersion.WithKind("Job"), newJobBuilder)
	rm.RegisterGVK(corev1.SchemeGroupVersion.WithKind("ServiceAccount"), newServiceAccountBuilder)
	rm.RegisterGVK(corev1.SchemeGroupVersion.WithKind("ConfigMap"), newConfigMapBuilder)
	rm.RegisterGVK(corev1.SchemeGroupVersion.WithKind("Namespace"), newNamespaceBuilder)
	rm.RegisterGVK(rbacv1.SchemeGroupVersion.WithKind("ClusterRole"), newClusterRoleBuilder)
	rm.RegisterGVK(rbacv1.SchemeGroupVersion.WithKind("ClusterRoleBinding"), newClusterRoleBindingBuilder)
	rm.RegisterGVK(rbacv1.SchemeGroupVersion.WithKind("Role"), newRoleBuilder)
	rm.RegisterGVK(rbacv1.SchemeGroupVersion.WithKind("RoleBinding"), newRoleBindingBuilder)
	rm.RegisterGVK(rbacv1beta1.SchemeGroupVersion.WithKind("ClusterRole"), newClusterRoleBuilder)
	rm.RegisterGVK(rbacv1beta1.SchemeGroupVersion.WithKind("ClusterRoleBinding"), newClusterRoleBindingBuilder)
	rm.RegisterGVK(rbacv1beta1.SchemeGroupVersion.WithKind("Role"), newRoleBuilder)
	rm.RegisterGVK(rbacv1beta1.SchemeGroupVersion.WithKind("RoleBinding"), newRoleBindingBuilder)

	rm.AddToMap(Mapper)
}
