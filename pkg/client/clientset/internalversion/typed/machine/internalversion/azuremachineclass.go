/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"context"
	"time"

	machine "github.com/gardener/machine-controller-manager/pkg/apis/machine"
	scheme "github.com/gardener/machine-controller-manager/pkg/client/clientset/internalversion/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AzureMachineClassesGetter has a method to return a AzureMachineClassInterface.
// A group's client should implement this interface.
type AzureMachineClassesGetter interface {
	AzureMachineClasses(namespace string) AzureMachineClassInterface
}

// AzureMachineClassInterface has methods to work with AzureMachineClass resources.
type AzureMachineClassInterface interface {
	Create(ctx context.Context, azureMachineClass *machine.AzureMachineClass, opts v1.CreateOptions) (*machine.AzureMachineClass, error)
	Update(ctx context.Context, azureMachineClass *machine.AzureMachineClass, opts v1.UpdateOptions) (*machine.AzureMachineClass, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*machine.AzureMachineClass, error)
	List(ctx context.Context, opts v1.ListOptions) (*machine.AzureMachineClassList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *machine.AzureMachineClass, err error)
	AzureMachineClassExpansion
}

// azureMachineClasses implements AzureMachineClassInterface
type azureMachineClasses struct {
	client rest.Interface
	ns     string
}

// newAzureMachineClasses returns a AzureMachineClasses
func newAzureMachineClasses(c *MachineClient, namespace string) *azureMachineClasses {
	return &azureMachineClasses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the azureMachineClass, and returns the corresponding azureMachineClass object, and an error if there is any.
func (c *azureMachineClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *machine.AzureMachineClass, err error) {
	result = &machine.AzureMachineClass{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azuremachineclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzureMachineClasses that match those selectors.
func (c *azureMachineClasses) List(ctx context.Context, opts v1.ListOptions) (result *machine.AzureMachineClassList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &machine.AzureMachineClassList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azuremachineclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azureMachineClasses.
func (c *azureMachineClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("azuremachineclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a azureMachineClass and creates it.  Returns the server's representation of the azureMachineClass, and an error, if there is any.
func (c *azureMachineClasses) Create(ctx context.Context, azureMachineClass *machine.AzureMachineClass, opts v1.CreateOptions) (result *machine.AzureMachineClass, err error) {
	result = &machine.AzureMachineClass{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("azuremachineclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(azureMachineClass).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a azureMachineClass and updates it. Returns the server's representation of the azureMachineClass, and an error, if there is any.
func (c *azureMachineClasses) Update(ctx context.Context, azureMachineClass *machine.AzureMachineClass, opts v1.UpdateOptions) (result *machine.AzureMachineClass, err error) {
	result = &machine.AzureMachineClass{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("azuremachineclasses").
		Name(azureMachineClass.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(azureMachineClass).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the azureMachineClass and deletes it. Returns an error if one occurs.
func (c *azureMachineClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azuremachineclasses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azureMachineClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azuremachineclasses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched azureMachineClass.
func (c *azureMachineClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *machine.AzureMachineClass, err error) {
	result = &machine.AzureMachineClass{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("azuremachineclasses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
