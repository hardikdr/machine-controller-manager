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

package fake

import (
	"context"

	machine "github.com/gardener/machine-controller-manager/pkg/apis/machine"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAzureMachineClasses implements AzureMachineClassInterface
type FakeAzureMachineClasses struct {
	Fake *FakeMachine
	ns   string
}

var azuremachineclassesResource = schema.GroupVersionResource{Group: "machine.sapcloud.io", Version: "", Resource: "azuremachineclasses"}

var azuremachineclassesKind = schema.GroupVersionKind{Group: "machine.sapcloud.io", Version: "", Kind: "AzureMachineClass"}

// Get takes name of the azureMachineClass, and returns the corresponding azureMachineClass object, and an error if there is any.
func (c *FakeAzureMachineClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *machine.AzureMachineClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(azuremachineclassesResource, c.ns, name), &machine.AzureMachineClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.AzureMachineClass), err
}

// List takes label and field selectors, and returns the list of AzureMachineClasses that match those selectors.
func (c *FakeAzureMachineClasses) List(ctx context.Context, opts v1.ListOptions) (result *machine.AzureMachineClassList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(azuremachineclassesResource, azuremachineclassesKind, c.ns, opts), &machine.AzureMachineClassList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &machine.AzureMachineClassList{ListMeta: obj.(*machine.AzureMachineClassList).ListMeta}
	for _, item := range obj.(*machine.AzureMachineClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azureMachineClasses.
func (c *FakeAzureMachineClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(azuremachineclassesResource, c.ns, opts))

}

// Create takes the representation of a azureMachineClass and creates it.  Returns the server's representation of the azureMachineClass, and an error, if there is any.
func (c *FakeAzureMachineClasses) Create(ctx context.Context, azureMachineClass *machine.AzureMachineClass, opts v1.CreateOptions) (result *machine.AzureMachineClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(azuremachineclassesResource, c.ns, azureMachineClass), &machine.AzureMachineClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.AzureMachineClass), err
}

// Update takes the representation of a azureMachineClass and updates it. Returns the server's representation of the azureMachineClass, and an error, if there is any.
func (c *FakeAzureMachineClasses) Update(ctx context.Context, azureMachineClass *machine.AzureMachineClass, opts v1.UpdateOptions) (result *machine.AzureMachineClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(azuremachineclassesResource, c.ns, azureMachineClass), &machine.AzureMachineClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.AzureMachineClass), err
}

// Delete takes name of the azureMachineClass and deletes it. Returns an error if one occurs.
func (c *FakeAzureMachineClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(azuremachineclassesResource, c.ns, name), &machine.AzureMachineClass{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzureMachineClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(azuremachineclassesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &machine.AzureMachineClassList{})
	return err
}

// Patch applies the patch and returns the patched azureMachineClass.
func (c *FakeAzureMachineClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *machine.AzureMachineClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(azuremachineclassesResource, c.ns, name, pt, data, subresources...), &machine.AzureMachineClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.AzureMachineClass), err
}
