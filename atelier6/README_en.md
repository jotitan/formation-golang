# Workshop 6
### Concurrency management

**Branch Workshop 6**

#### Objectives

* _Distribute_ our tasks across multiple instances and call them intelligently.
* _Secure_ usage by managing a timeout

#### Steps

Currently, the coordinator is not very intelligent: when he receives a task, he immediately dispatches it to a worker.
This mechanism, the round robin, has several disadvantages:
* It does not take into account the processing capacity of the worker
* It is not possible to limit the number of requests on a worker

We are going to change the mecanism by using channels

These are the steps :
* Generate a unique id (uuid) during the worker creation
* When registering the worker with the coordinator, transmit the uuid
* We can also transmit the processing capacity of a worker (number of tasks in parallel)
* In the status of a task, allow the storage of the identifier of the node which processes it
* When a worker indicates that a task is completed (or in error), pass the uuid
* The running status is not currently used (because the task is sent immediately to the worker): see how to integrate this status
* We will have to ensure that the taskManager and the PoolWorker communicate with each other. Add tests.

#### Help

* To generate a uuid, the [google library](github.com/google/uuid) is easy to use
* Reminders about channels :
    * The limit the number of parallel calls, we can use a channel with a size corresponding to the number of parallel tasks
    * To exchange data in a threadsafe manner, the channel is an ideal solution
