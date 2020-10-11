# Frequently Asked Questions

The answers in this FAQ apply to the newest (HEAD) version of Machine Controller Manager. If
you're using an older version of MCM please refer to corresponding version of
this document. Few of answers assumes that the MCM is being used in conjuction with [cluster-autoscaler](https://github.com/gardener/autoscaler):

# Table of Contents:
<!--- TOC BEGIN -->
* [Basics](#basics)
  * [What is Machine Controller Manager?](#what-is-machine-controller-manager)
  * [Why is my machine deleted ?](#Why-is-my-machine-deleted)
  * [What are different sub-controllers in MCM ?](#What-are-different-sub-controllers-in-MCM)
  * [What is safety-controller in MCM ?](#What-is-safety-controller-in-MCM)

* [How to?](#how-to)
  * [How to install the MCM in a kubernetes cluster?](#How-to-install-the-MCM-in-a-kubernetes-cluster)
  * [How to better control the roll-out process of the worker-nodes?](#How-to-better-control-the-roll-out-process-of-the-worker-nodes)
  * [How to scale-down the machine-deployment by selectively deleting the machines?](#How-to-scale-down-the-machine-deployment-by-selectively-deleting-the-machines)
  * [How to force-delete the machine?](#How-to-force-delete-the-machine)

* [Internals](#internals)
  * [What is the high-level design of the MCM?](#What-is-the-high-level-design-of-the-MCM)
  * [What are different configuration option in MCM?](#What-are-different-configuration-option-in-MCM)
  * [What are the different timeouts/configurations in the machine's lifecycle?](#What-are-the-different-timeouts/configurations-in-the-machine's-lifecycle)
  * [How is the drain of the machine implemented?](#How-is-the-drain-of-the-machine-implemented)
  * [How are the stateful applications drained during machine-deletion??](#How-are-the-stateful-applications-drained-during-machine-deletion?)
  * [How does maxEvictRetries configuration work with drainTimeout configuration?](#How-does-maxEvictRetries-configuration-work-with-drainTimeout-configuration)
  * [What are different phases of the machine?](#What-are-different-phases-of-the-machine)

* [Troubleshooting](#troubleshooting)
  * [My machine is stuck in deletion for 1 hr, why?](#My-machine-is-stuck-in-deletion-for-1-hr-why)
  * [My machine is not joining the cluster, why?](#My-machine-is-not-joining-the-cluster-why)
* [Developer](#developer)
  * [How should I test my code before submitting a PR?](#How-should-I-test-my-code-before-submitting-a-PR)
  * [I need to change the APIs, what are the recommended steps?](#I-need-to-change-the-APIs-what-are-the-recommended-steps)
  * [How can I update the dependencies of MCM?](#How-can-I-update-the-dependencies-of-MCM)
* [In the context of Gardener](#in-the-context-of-gardener)
  * [How can I configure MCM using Shoot-resource?](#How-can-I-configure-MCM-using-Shoot-resource)
  * [How is my worker-pool spread across zones?](#How-is-my-worker-pool-spread-across-zones)

<!--- TOC END -->

# Basics

### What is Machine Controller Manager ?

Machine controller manager aka MCM, is a bunch of controllers which is used to manage the lifecycle of the worker-machines. It reconciles a set of CRDs such as `Machine`, `MachineSet`, `MachineDeployment` which depicts the functionality of `Pod`, `Replicaset`, `Deployment` of the core kubernetes respectively. Read more about it at [README](https://github.com/gardener/machine-controller-manager/tree/master/docs).

* It is also used at gardener to manage the kubernetes-nodes of the shoot-cluster. By design it can be used independent of the gardener itself.

### Why is my machine deleted ?

A machine could be deleted by MCM, generally for 2 reasons.

1. Machine has been unhealthy for at least `MachineHealthTimeout` period, default is 10 minutes.

   * Machine is considered unhealthy if any of the following node-conditions `DiskPressure,KernelDeadlock,FileSystemReadonly` are set to true, or `KubeletReady` is set to false.
2. Machine has been scale-down by the MachineDeployment resource.

   * This is very usual if cluster-autoscaler(aka CA) is being used with MCM. CA deletes the under-utilized machines by scaling down the MachineDeployment. Read more about cluster-autoscaler's scale-down behavior [here](https://github.com/gardener/autoscaler/blob/machine-controller-manager-provider/cluster-autoscaler/FAQ.md#how-does-scale-down-work).

### What are different sub-controllers in MCM ?

MCM mainly contains following sub-controllers:

* Machine-deployment controller, responsible for reconciling the `MachineDeployment` object. It manages the lifecycle of the machine-set objects.
* Machine-set controller, responsible for reconciling the `MachineSet` object. It manages the lifecycle of the machine objects.
* Machine-controller, responsible for reconciling the `Machine` object. This controller manages the lifecycle of the actual VMs/machines created in cloud/on-prem. This controller has been moved out of tree, please refer an AWS machine-controller for more info- [link](https://github.com/gardener/machine-controller-manager-provider-gcp).
* Safety-controller, responsible for handling the unidentified/unknown behaviors from the cloud-providers. Please read more about it functionality [below](#what-is-safety-controller).

### What is safety-controller in MCM ?

Safety-controller contains following functions:

* Orphan VM handler:
  * It lists all the VMs in the cloud matching the `tag` of given cluster-name. Then it maps the VMs with the machine-objects using the `ProviderID` field. VMs which don't have any backing machine-objects are logged and deleted after confirmation.
  * This handler is run every 30 minutes, and configurable via flag [machine-safety-orphan-vms-period](https://github.com/gardener/machine-controller-manager/blob/master/cmd/machine-controller-manager/app/options/options.go#L112).
* Freeze mechanism: 
  * Safety-controller freezes the machine-deployment and machine-set controller if the number of machine-objects goes beyond a certain threshold on top of `Spec.Replicas`. It can be configured by the flag [--safety-up or --safety-down](https://github.com/gardener/machine-controller-manager/blob/master/cmd/machine-controller-manager/app/options/options.go#L102-L103) and also [machine-safety-overshooting-period](https://github.com/gardener/machine-controller-manager/blob/master/cmd/machine-controller-manager/app/options/options.go#L113).
  * Safety-controller freezes the functionality of the MCM if either of the target-apiserver or the control-apiserver is not reachable.
  * Safety-controller unfreezes the MCM automatically once situation is resolved to normal. A `freeze` label has been put on MachineDeployment/MachineSet to enforce the freeze condition.

# How to?

### How to install the MCM in a kubernetes cluster?

MCM can be installed in a cluster with following steps:

* Apply all the CRDs from [here](https://github.com/gardener/machine-controller-manager/tree/master/kubernetes/deployment/in-tree)
* Apply all the deployment, role-related objects from [here](https://github.com/gardener/machine-controller-manager/tree/master/kubernetes/deployment/in-tree).

  * Control-cluster is one where the machine-* objects are stored. Target cluster is where all the node-objects are registered.

### How to better control the roll-out process of the worker-nodes?

MCM allows configuring the roll-out of the worker-machines using `maxSurge` and `maxUnavailable` fields. These fields are applicable only during the roll-out process, and means nothing in general scale-up/down scenarios.
The overall process is very similar to how the `deployment-controller` manages pods during `RollingUpdate`.

* `maxSurge` refers to the number of more machines which can be added on top of the `Spec.Replicas` of MachineDeployment _during rollout process_.
* `maxUnavailable` refers to the number of machines which could be deleted from `Spec.Replicas` field of the MachineDeployment _during rollout process_.


### How to scale-down the machine-deployment by selectively deleting the machines?

During scale-down triggered via machine-deployment/set, MCM prefers to delete the machine/s which have the least priority set.
Each machine-object has an annotation `machinepriority.machine.sapcloud.io` set to `3` by default. Admin can reduce the priority of the given machines by changing the annotation value to `1`. The next scale-down from MachineDeployment shall delete the machines with the least priority first.

### How to force-delete the machine?

A machine can be force deleted by adding the label `force-deletion: "True"` on the machine-object if it's already being deleted. During force-deletion, MCM skips the drain function and simply triggers the deletion of the machine.
This label should be used with caution, as it can violate the PDBs for pods running on machine.


# Internals
### What is the high-level design of the MCM?

Please refer the following [document](https://github.com/gardener/machine-controller-manager/tree/master/docs/design).

### What are different configuration option in MCM?

MCM allows configuring many knobs to fine-tune its behavior according to the user's need. 
Please refer to the [link](https://github.com/gardener/machine-controller-manager/blob/master/cmd/machine-controller-manager/app/options/options.go) to check the exact configuration options.

### What are the different timeouts/configurations in the machine's lifecycle?

A machine's lifecycle is governed by mainly following timeouts, which can be configured [here](https://github.com/gardener/machine-controller-manager/blob/master/kubernetes/machine_objects/machine-deployment.yaml#L30-L34).

* MachineDrainTimeout: Amount of time after which drain times out, and machine is force-deleted. Default ~2 hours.
* MachineHealthTimeout: Amount of time after which an unhealthy machine is declared `Failed` and machine is replaced by machine-set controller.
* MachineCreationTimeout: Amount of time which machine creation is declared `Failed`, and machine is replaced by the machine-set controller.
* NodeConditions: List of node-conditions which are if set to true for `MachineHealthTimeout` period, the machine is declared failed, and replaced by machine-set controller.
* MaxEvictRetries: An integer number, depicting the number of times a failed _eviction_ should be retried on a pod during drain process. A pod is _deleted_ after max-retries.

### How is the drain of the machine implemented?

MCM imports the functionality from the upstream kubernetes-drain library. Although, few parts have been modified to make it work best in the context of MCM. Drain is executed before machine-deletion, for gracefully migrate the applications. 
Drain internally uses the `EvictionAPI` to evict the pods, and later triggers the `Deletion` of pods after `MachineDrainTimeout`. Please note:

* Stateless pods are evicted are in parallel.
* Stateful applications (with PVCs) are serially evicted. Please find more info in this [answer below](How-are-the-stateful-applications-drained-during-machine-deletion?).


### How are the stateful applications drained during machine-deletion?

Drain function serially evicts the stateful-pods. We identified that serially evicting stateful pods yields better overall availability of pods, as underlying cloud in most cases detaches and reattaches disks serially anyways.
It is implemented in the following manner::

* Drain lists all the pods with attached volumes. It evicts very first stateful-pod, and waits for its related entry in Node-object's `.status.volumesAttached`, to be removed by KCM. It does the same process for all the stateful-pods.
* It waits for `PvDetachTimeout`(default 2 minutes) for a given pod's PVC to be removed, else moves forward.

### How does maxEvictRetries configuration work with drainTimeout configuration?

We recommend you to set only `MachineDrainTimeout`, it should satisfy the related requirements. `MaxEvictRetries` is auto-calculated based on `MachineDrainTimeout`, if maxEvictRetries is not provided.
Though following will be overall behavior of both configurations together:

* If you dont set the maxEvictRetries, and set only the maxDrainTimeout.
  * MCM auto calculates the maxEvictRetries based on the drainTimeout.
* If you dont set the drainTimeout, and only set the maxEvictRetries:
  * Default drainTimeout is considered , and the maxEvictRetries you provided for each pod.
* If you set both maxEvictRetries and drainTimoeut:
  * Then both will be respected.
* If you set none:
  * Defaults are respected.

### What are different phases of the machine?

A phase of the machine can be identified with `Machine.Status.CurrentStatus.Phase`. Following are the possible phases of the machine-object:

* `Pending`: Machine-creation call has succeed, and MCM is waiting for machine to join the cluster.
* `CrashLoopBackOff`: Machine-creation call has failed, and MCM will retry the operation after a minor delay.
* `Running`: Machine-creation call has succeed, and Machine has joined the cluster successfully.
* `Unknown`: Machine health-checks are failing, eg kubelet has stopped posting the status.
* `Failed`: Machine health-checks have failed for a prolonged time, and hence it is declared failed. Machine-set controller will replace such machines immediately.
* `Terminating`: Machine is being terminaed. Terminating state is set immediately the deletion has been triggered for machine-object, also includes time when it's being drained. 

# Troubleshooting
### My machine is stuck in deletion for 1 hr, why?

In most cases, the `Machine.Status.LastOperation` should provide information around why machine can't be deleted.
Though following could be the reasons but not limited to:

* Pod/s with mis-configured PDBs block the drain operation. PDBs with `maxUnavailable` set to 0, doesn't allow eviction of the pods, hence drain/eviction is retried till `MachineDrainTimeout`. Default `MachineDrainTimeout` could be as large as ~2hours, hence blocking the machine-deletion. 
  * Short term: you can manually delete the pod in question, _with caution_. 
  * Long term: please set more appropriate PDBs, which allows disruption of at least one pod.
* Expired cloud-credentials can block the deletion of the machine from infrastructure.
* Cloud provider cant delete the machine due to internal errors, such situations are best debugged by using cloud-provider specific CLI or cloud-console.


### My machine is not joining the cluster, why?

In most cases, the `Machine.Status.LastOperation` should provide information around why machine can't be created.
It could possibly be debugged with following steps:

* Please verify if machine is actually created in the cloud. You can use the `Machine.Spec.ProviderId` to query the machine in cloud.
* A kubernetes node is generally bootstrapped with the cloud-config. Please verify, if MachineDeployment is pointing the correct MachineClass, and MachineClass is pointing to the correct `Secret`. The secret object contains the actual cloud-config in base64 format which will be used to boot the machine.
* You must also check the logs of the MCM pod to understand if any logical flow of reconciliation is broken.


# Developer

### How should I test my code before submitting a PR?

You can locally setup the MCM using following [guide](https://github.com/gardener/machine-controller-manager/blob/master/docs/development/local_setup.md). You must also enhance the unit-tests related to your changes. You can locally run the unit-test by executing:
```
ginkgo --cover pkg/controller/.
```

### I need to change the APIs, what are the recommended steps?

You should add/update the API fields at both of the following places:

* https://github.com/gardener/machine-controller-manager/blob/master/pkg/apis/machine/types.go
* https://github.com/gardener/machine-controller-manager/tree/master/pkg/apis/machine/v1alpha1

Once API changes are done, you should auto-generate the code using following command:
```
./hack/generate-code
```
Please ignore the API-violation errors for now.

### How can I update the dependencies of MCM?

MCM uses `gomod` for depedency management.
You should add/udpate depedency in the go.mod file. Please run following command to automatically revendor the dependencies.
```
make revendor
```


# In the context of Gardener

### How can I configure MCM using Shoot-resource?

All of the knobs of MCM can be configured by the `workers` [section](https://github.com/gardener/gardener/blob/master/example/90-shoot.yaml#L29-L126) of the shoot-resource.

* Gardener creates a machine-deployment per each zone for each worker-pool under `workers` section. 
* `workers.dataVolumes` allows to attach multiple disks to a machine during creation. Refer the [link](https://github.com/gardener/gardener/blob/master/example/90-shoot.yaml#L29-L126).
* `workers.machineControllerManager` allows to configure multiple knobs of the MachineDeployment from the shoot-resource.

### How is my worker-pool spread across zones?

Shoot resource allows to spread the worker-pool across multiple zones using the field `workers.zones`, refer [link](https://github.com/gardener/gardener/blob/master/example/90-shoot.yaml#L115).

* Gardener creates one machine-deployment per each zones. Each machine deployment is initiated with the following replica:
```
MachineDeployment.Spec.Replicas = (Workers.Minimum)/(Number of availibility zones)
```