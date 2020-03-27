// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/openshift/aws-ebs-csi-driver-operator/pkg/apis/operator/v1alpha1"
	scheme "github.com/openshift/aws-ebs-csi-driver-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EBSCSIDriversGetter has a method to return a EBSCSIDriverInterface.
// A group's client should implement this interface.
type EBSCSIDriversGetter interface {
	EBSCSIDrivers() EBSCSIDriverInterface
}

// EBSCSIDriverInterface has methods to work with EBSCSIDriver resources.
type EBSCSIDriverInterface interface {
	Create(ctx context.Context, eBSCSIDriver *v1alpha1.EBSCSIDriver, opts v1.CreateOptions) (*v1alpha1.EBSCSIDriver, error)
	Update(ctx context.Context, eBSCSIDriver *v1alpha1.EBSCSIDriver, opts v1.UpdateOptions) (*v1alpha1.EBSCSIDriver, error)
	UpdateStatus(ctx context.Context, eBSCSIDriver *v1alpha1.EBSCSIDriver, opts v1.UpdateOptions) (*v1alpha1.EBSCSIDriver, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.EBSCSIDriver, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.EBSCSIDriverList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EBSCSIDriver, err error)
	EBSCSIDriverExpansion
}

// eBSCSIDrivers implements EBSCSIDriverInterface
type eBSCSIDrivers struct {
	client rest.Interface
}

// newEBSCSIDrivers returns a EBSCSIDrivers
func newEBSCSIDrivers(c *CsiV1alpha1Client) *eBSCSIDrivers {
	return &eBSCSIDrivers{
		client: c.RESTClient(),
	}
}

// Get takes name of the eBSCSIDriver, and returns the corresponding eBSCSIDriver object, and an error if there is any.
func (c *eBSCSIDrivers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EBSCSIDriver, err error) {
	result = &v1alpha1.EBSCSIDriver{}
	err = c.client.Get().
		Resource("ebscsidrivers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EBSCSIDrivers that match those selectors.
func (c *eBSCSIDrivers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EBSCSIDriverList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EBSCSIDriverList{}
	err = c.client.Get().
		Resource("ebscsidrivers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested eBSCSIDrivers.
func (c *eBSCSIDrivers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ebscsidrivers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a eBSCSIDriver and creates it.  Returns the server's representation of the eBSCSIDriver, and an error, if there is any.
func (c *eBSCSIDrivers) Create(ctx context.Context, eBSCSIDriver *v1alpha1.EBSCSIDriver, opts v1.CreateOptions) (result *v1alpha1.EBSCSIDriver, err error) {
	result = &v1alpha1.EBSCSIDriver{}
	err = c.client.Post().
		Resource("ebscsidrivers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eBSCSIDriver).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a eBSCSIDriver and updates it. Returns the server's representation of the eBSCSIDriver, and an error, if there is any.
func (c *eBSCSIDrivers) Update(ctx context.Context, eBSCSIDriver *v1alpha1.EBSCSIDriver, opts v1.UpdateOptions) (result *v1alpha1.EBSCSIDriver, err error) {
	result = &v1alpha1.EBSCSIDriver{}
	err = c.client.Put().
		Resource("ebscsidrivers").
		Name(eBSCSIDriver.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eBSCSIDriver).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *eBSCSIDrivers) UpdateStatus(ctx context.Context, eBSCSIDriver *v1alpha1.EBSCSIDriver, opts v1.UpdateOptions) (result *v1alpha1.EBSCSIDriver, err error) {
	result = &v1alpha1.EBSCSIDriver{}
	err = c.client.Put().
		Resource("ebscsidrivers").
		Name(eBSCSIDriver.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eBSCSIDriver).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the eBSCSIDriver and deletes it. Returns an error if one occurs.
func (c *eBSCSIDrivers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ebscsidrivers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *eBSCSIDrivers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ebscsidrivers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched eBSCSIDriver.
func (c *eBSCSIDrivers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EBSCSIDriver, err error) {
	result = &v1alpha1.EBSCSIDriver{}
	err = c.client.Patch(pt).
		Resource("ebscsidrivers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
