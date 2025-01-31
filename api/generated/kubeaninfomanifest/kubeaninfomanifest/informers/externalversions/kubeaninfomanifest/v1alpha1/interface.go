// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "kubean.io/api/generated/kubeaninfomanifest/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// KubeanInfoManifests returns a KubeanInfoManifestInformer.
	KubeanInfoManifests() KubeanInfoManifestInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// KubeanInfoManifests returns a KubeanInfoManifestInformer.
func (v *version) KubeanInfoManifests() KubeanInfoManifestInformer {
	return &kubeanInfoManifestInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
