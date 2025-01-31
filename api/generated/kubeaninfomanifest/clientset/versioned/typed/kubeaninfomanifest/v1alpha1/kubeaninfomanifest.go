// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubean.io/api/apis/kubeaninfomanifest/v1alpha1"
	scheme "kubean.io/api/generated/kubeaninfomanifest/clientset/versioned/scheme"
)

// KubeanInfoManifestsGetter has a method to return a KubeanInfoManifestInterface.
// A group's client should implement this interface.
type KubeanInfoManifestsGetter interface {
	KubeanInfoManifests() KubeanInfoManifestInterface
}

// KubeanInfoManifestInterface has methods to work with KubeanInfoManifest resources.
type KubeanInfoManifestInterface interface {
	Create(ctx context.Context, kubeanInfoManifest *v1alpha1.KubeanInfoManifest, opts v1.CreateOptions) (*v1alpha1.KubeanInfoManifest, error)
	Update(ctx context.Context, kubeanInfoManifest *v1alpha1.KubeanInfoManifest, opts v1.UpdateOptions) (*v1alpha1.KubeanInfoManifest, error)
	UpdateStatus(ctx context.Context, kubeanInfoManifest *v1alpha1.KubeanInfoManifest, opts v1.UpdateOptions) (*v1alpha1.KubeanInfoManifest, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.KubeanInfoManifest, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.KubeanInfoManifestList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KubeanInfoManifest, err error)
	KubeanInfoManifestExpansion
}

// kubeanInfoManifests implements KubeanInfoManifestInterface
type kubeanInfoManifests struct {
	client rest.Interface
}

// newKubeanInfoManifests returns a KubeanInfoManifests
func newKubeanInfoManifests(c *KubeanV1alpha1Client) *kubeanInfoManifests {
	return &kubeanInfoManifests{
		client: c.RESTClient(),
	}
}

// Get takes name of the kubeanInfoManifest, and returns the corresponding kubeanInfoManifest object, and an error if there is any.
func (c *kubeanInfoManifests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KubeanInfoManifest, err error) {
	result = &v1alpha1.KubeanInfoManifest{}
	err = c.client.Get().
		Resource("kubeaninfomanifests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KubeanInfoManifests that match those selectors.
func (c *kubeanInfoManifests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KubeanInfoManifestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KubeanInfoManifestList{}
	err = c.client.Get().
		Resource("kubeaninfomanifests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kubeanInfoManifests.
func (c *kubeanInfoManifests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("kubeaninfomanifests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a kubeanInfoManifest and creates it.  Returns the server's representation of the kubeanInfoManifest, and an error, if there is any.
func (c *kubeanInfoManifests) Create(ctx context.Context, kubeanInfoManifest *v1alpha1.KubeanInfoManifest, opts v1.CreateOptions) (result *v1alpha1.KubeanInfoManifest, err error) {
	result = &v1alpha1.KubeanInfoManifest{}
	err = c.client.Post().
		Resource("kubeaninfomanifests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kubeanInfoManifest).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a kubeanInfoManifest and updates it. Returns the server's representation of the kubeanInfoManifest, and an error, if there is any.
func (c *kubeanInfoManifests) Update(ctx context.Context, kubeanInfoManifest *v1alpha1.KubeanInfoManifest, opts v1.UpdateOptions) (result *v1alpha1.KubeanInfoManifest, err error) {
	result = &v1alpha1.KubeanInfoManifest{}
	err = c.client.Put().
		Resource("kubeaninfomanifests").
		Name(kubeanInfoManifest.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kubeanInfoManifest).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *kubeanInfoManifests) UpdateStatus(ctx context.Context, kubeanInfoManifest *v1alpha1.KubeanInfoManifest, opts v1.UpdateOptions) (result *v1alpha1.KubeanInfoManifest, err error) {
	result = &v1alpha1.KubeanInfoManifest{}
	err = c.client.Put().
		Resource("kubeaninfomanifests").
		Name(kubeanInfoManifest.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kubeanInfoManifest).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the kubeanInfoManifest and deletes it. Returns an error if one occurs.
func (c *kubeanInfoManifests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("kubeaninfomanifests").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kubeanInfoManifests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("kubeaninfomanifests").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched kubeanInfoManifest.
func (c *kubeanInfoManifests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KubeanInfoManifest, err error) {
	result = &v1alpha1.KubeanInfoManifest{}
	err = c.client.Patch(pt).
		Resource("kubeaninfomanifests").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
