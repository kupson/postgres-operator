package controller

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"

	"github.com/zalando/postgres-operator/pkg/util"
)

func getPodEnvironmentSecretListWatchOptions(options metav1.ListOptions, secretName string) metav1.ListOptions {
	return metav1.ListOptions{
		Watch:           options.Watch,
		ResourceVersion: options.ResourceVersion,
		TimeoutSeconds:  options.TimeoutSeconds,
		FieldSelector: fields.OneTermEqualSelector("metadata.name", secretName).String(),
	}
}

func (c *Controller) podEnvironmentSecretListFunc(options metav1.ListOptions) (runtime.Object, error) {
	return c.KubeClient.Secrets(c.opConfig.WatchedNamespace).List(getPodEnvironmentSecretListWatchOptions(options, c.opConfig.PodEnvironmentSecretName))
}

func (c *Controller) podEnvironmentSecretWatchFunc(options metav1.ListOptions) (watch.Interface, error) {
	return c.KubeClient.Secrets(c.opConfig.WatchedNamespace).Watch(getPodEnvironmentSecretListWatchOptions(options, c.opConfig.PodEnvironmentSecretName))
}

func (c *Controller) podEnvironmentSecretAdd(obj interface{}) {
	secret, ok := obj.(*v1.Secret)
	if !ok {
		return
	}

	c.logger.Debugf("the pod environment secret has been added: %q", util.NameFromMeta(secret.ObjectMeta))
}

func (c *Controller) podEnvironmentSecretUpdate(_, obj interface{}) {
	secret, ok := obj.(*v1.Secret)
	if !ok {
		return
	}

	c.logger.Debugf("the pod environment secret has been updated: %q", util.NameFromMeta(secret.ObjectMeta))
}

func (c *Controller) podEnvironmentSecretDelete(obj interface{}) {
	secret, ok := obj.(*v1.Secret)
	if !ok {
		return
	}

	c.logger.Debugf("the pod environment secret has been deleted: %q", util.NameFromMeta(secret.ObjectMeta))
}
