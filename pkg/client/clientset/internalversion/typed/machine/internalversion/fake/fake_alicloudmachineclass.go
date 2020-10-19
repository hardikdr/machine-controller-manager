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

// FakeAlicloudMachineClasses implements AlicloudMachineClassInterface
type FakeAlicloudMachineClasses struct {
	Fake *FakeMachine
	ns   string
}

var alicloudmachineclassesResource = schema.GroupVersionResource{Group: "machine.sapcloud.io", Version: "", Resource: "alicloudmachineclasses"}

var alicloudmachineclassesKind = schema.GroupVersionKind{Group: "machine.sapcloud.io", Version: "", Kind: "AlicloudMachineClass"}

// Get takes name of the alicloudMachineClass, and returns the corresponding alicloudMachineClass object, and an error if there is any.
func (c *FakeAlicloudMachineClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *machine.AlicloudMachineClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(alicloudmachineclassesResource, c.ns, name), &machine.AlicloudMachineClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.AlicloudMachineClass), err
}

// List takes label and field selectors, and returns the list of AlicloudMachineClasses that match those selectors.
func (c *FakeAlicloudMachineClasses) List(ctx context.Context, opts v1.ListOptions) (result *machine.AlicloudMachineClassList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(alicloudmachineclassesResource, alicloudmachineclassesKind, c.ns, opts), &machine.AlicloudMachineClassList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &machine.AlicloudMachineClassList{ListMeta: obj.(*machine.AlicloudMachineClassList).ListMeta}
	for _, item := range obj.(*machine.AlicloudMachineClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested alicloudMachineClasses.
func (c *FakeAlicloudMachineClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(alicloudmachineclassesResource, c.ns, opts))

}

// Create takes the representation of a alicloudMachineClass and creates it.  Returns the server's representation of the alicloudMachineClass, and an error, if there is any.
func (c *FakeAlicloudMachineClasses) Create(ctx context.Context, alicloudMachineClass *machine.AlicloudMachineClass, opts v1.CreateOptions) (result *machine.AlicloudMachineClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(alicloudmachineclassesResource, c.ns, alicloudMachineClass), &machine.AlicloudMachineClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.AlicloudMachineClass), err
}

// Update takes the representation of a alicloudMachineClass and updates it. Returns the server's representation of the alicloudMachineClass, and an error, if there is any.
func (c *FakeAlicloudMachineClasses) Update(ctx context.Context, alicloudMachineClass *machine.AlicloudMachineClass, opts v1.UpdateOptions) (result *machine.AlicloudMachineClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(alicloudmachineclassesResource, c.ns, alicloudMachineClass), &machine.AlicloudMachineClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.AlicloudMachineClass), err
}

// Delete takes name of the alicloudMachineClass and deletes it. Returns an error if one occurs.
func (c *FakeAlicloudMachineClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(alicloudmachineclassesResource, c.ns, name), &machine.AlicloudMachineClass{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAlicloudMachineClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(alicloudmachineclassesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &machine.AlicloudMachineClassList{})
	return err
}

// Patch applies the patch and returns the patched alicloudMachineClass.
func (c *FakeAlicloudMachineClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *machine.AlicloudMachineClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(alicloudmachineclassesResource, c.ns, name, pt, data, subresources...), &machine.AlicloudMachineClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.AlicloudMachineClass), err
}
