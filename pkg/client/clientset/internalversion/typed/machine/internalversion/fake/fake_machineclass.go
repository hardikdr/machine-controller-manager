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

// FakeMachineClasses implements MachineClassInterface
type FakeMachineClasses struct {
	Fake *FakeMachine
	ns   string
}

var machineclassesResource = schema.GroupVersionResource{Group: "machine.sapcloud.io", Version: "", Resource: "machineclasses"}

var machineclassesKind = schema.GroupVersionKind{Group: "machine.sapcloud.io", Version: "", Kind: "MachineClass"}

// Get takes name of the machineClass, and returns the corresponding machineClass object, and an error if there is any.
func (c *FakeMachineClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *machine.MachineClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(machineclassesResource, c.ns, name), &machine.MachineClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.MachineClass), err
}

// List takes label and field selectors, and returns the list of MachineClasses that match those selectors.
func (c *FakeMachineClasses) List(ctx context.Context, opts v1.ListOptions) (result *machine.MachineClassList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(machineclassesResource, machineclassesKind, c.ns, opts), &machine.MachineClassList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &machine.MachineClassList{ListMeta: obj.(*machine.MachineClassList).ListMeta}
	for _, item := range obj.(*machine.MachineClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested machineClasses.
func (c *FakeMachineClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(machineclassesResource, c.ns, opts))

}

// Create takes the representation of a machineClass and creates it.  Returns the server's representation of the machineClass, and an error, if there is any.
func (c *FakeMachineClasses) Create(ctx context.Context, machineClass *machine.MachineClass, opts v1.CreateOptions) (result *machine.MachineClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(machineclassesResource, c.ns, machineClass), &machine.MachineClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.MachineClass), err
}

// Update takes the representation of a machineClass and updates it. Returns the server's representation of the machineClass, and an error, if there is any.
func (c *FakeMachineClasses) Update(ctx context.Context, machineClass *machine.MachineClass, opts v1.UpdateOptions) (result *machine.MachineClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(machineclassesResource, c.ns, machineClass), &machine.MachineClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.MachineClass), err
}

// Delete takes name of the machineClass and deletes it. Returns an error if one occurs.
func (c *FakeMachineClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(machineclassesResource, c.ns, name), &machine.MachineClass{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMachineClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(machineclassesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &machine.MachineClassList{})
	return err
}

// Patch applies the patch and returns the patched machineClass.
func (c *FakeMachineClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *machine.MachineClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(machineclassesResource, c.ns, name, pt, data, subresources...), &machine.MachineClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.MachineClass), err
}
