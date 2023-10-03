# Workshop 2.1
### Data structure and interface

#### Objectives

* Refactoring code

The code we wrote is good but it exposes too much of the implementation of our tasks when we would like to hide that.
We will also add the notion of task deletion to the taskmanager.

* Explain the benefit of the manager's NextId method and implement it.
* Implement the TaskManager's Remove method and run the taskmanager_test.go tests
* Modify Print and Resize tasks to not expose the implementation


#### Help
* Initializing the modules in a project
```bash
go mod init project_name 
```
* To avoid exposing the detail of a task, we can create a function which creates the task, NewPrint for example, and we change the visibility of Print
* As a reminder, when the name of a structure begins with a capital letter, it is public, otherwise private