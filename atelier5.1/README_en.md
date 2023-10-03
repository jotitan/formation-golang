# Workshop 5.1
### Connect the coordinator to the worker(s)

#### Objectives

* Know how to manage a server pool

#### Steps

The coordinator will receive requests and delegate them to a pool of workers.
Here are the different steps

- When a worker connects, it knows the coordinator's URL and connects to it
- When a task is received, the coordinator will delegate it to one of the connected workers
- When the worker has finished the task, it will notify the coordinator

The code to add :
- On the coordinator, add a _POST_ **/register** route where the worker indicates its url
- Implement a simple version of the pool worker to execute a task (not always the same worker)
    - The **TaskSenderToWorker** interface allows you to decorrelate pool management from the sending action (notably for mocking)
    - When starting a worker, register with the coordinator
    - Modify the TestCompleteChain test in order to test the chain: launching the coordinator, the worker and adding a print task until the notification

#### Help

* To sort a list : sort.Slice