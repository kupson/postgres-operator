package controller

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"

	"github.com/zalando/postgres-operator/pkg/util"
)

func getPodEnvironmentConfigMapListWatchOptions(options metav1.ListOptions, configMapName string) metav1.ListOptions {
	return metav1.ListOptions{
		Watch:           options.Watch,
		ResourceVersion: options.ResourceVersion,
		TimeoutSeconds:  options.TimeoutSeconds,
		FieldSelector: fields.OneTermEqualSelector("metadata.name", configMapName).String(),
	}
}

func (c *Controller) podEnvironmentConfigMapListFunc(options metav1.ListOptions) (runtime.Object, error) {
	return c.KubeClient.ConfigMaps(c.opConfig.WatchedNamespace).List(getPodEnvironmentConfigMapListWatchOptions(options, c.opConfig.PodEnvironmentConfigMap))
}

func (c *Controller) podEnvironmentConfigMapWatchFunc(options metav1.ListOptions) (watch.Interface, error) {
	return c.KubeClient.ConfigMaps(c.opConfig.WatchedNamespace).Watch(getPodEnvironmentConfigMapListWatchOptions(options, c.opConfig.PodEnvironmentConfigMap))
}

func (c *Controller) podEnvironmentConfigMapAdd(obj interface{}) {
	configMap, ok := obj.(*v1.ConfigMap)
	if !ok {
		return
	}

	c.logger.Debugf("the pod environment configmap has been added: %q", util.NameFromMeta(configMap.ObjectMeta))
}

func (c *Controller) podEnvironmentConfigMapUpdate(_, obj interface{}) {
	configMap, ok := obj.(*v1.ConfigMap)
	if !ok {
		return
	}

	c.logger.Debugf("the pod environment configmap has been updated: %q", util.NameFromMeta(configMap.ObjectMeta))
}

func (c *Controller) podEnvironmentConfigMapDelete(obj interface{}) {
	configMap, ok := obj.(*v1.ConfigMap)
	if !ok {
		return
	}

	c.logger.Debugf("the pod environment configmap has been deleted: %q", util.NameFromMeta(configMap.ObjectMeta))
}
