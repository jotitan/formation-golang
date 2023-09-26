# Workshop 2.0
### Data structure and interface

#### Objectives

* Initialize a project with modules
* Model data structures to enable image resizing
* Use interfaces to represent a task that can be executed

This workshop is done in two steps : workshop2.0 et workshop2.1

#### Step 1 => Directory workshop2.0

##### Task management

* Initialize the modules with project name "formation-go"
* Modify the example of workshop 1 to use the **logger** package and launch the test newgopher_test.go
* In the model package which contains the useful structures
  * The Task interface represents a task to execute : add to it the method **Do()**
  * Add to the Print structure a Message field
  * Implement the necessary methods for **Print** to be a **Task**
  * Add to the Resizze structure the necessary fields :
    * width
    * length
    * path to the source image
    * path to the target destination
  * Implement in Resize the method **Do()** in order to respect the Task interface.
  The implementation will be basic and will only display a message with the logger : "Run resize from, to, height, width".
  
  _Exemple "Run resize img.jpg img2.jpg 400px 200px"_
  * Launch the tests tasks_test.go

##### Table manipulation
* Implement the methods of the structure Taskmanager
* launch the test taskmanager_test.go

#### Help
* Initializing the modules in a project (in the src directory)
```bash
go mod init project_name 
```
* Launch a test : change path to be in the directory and launch : 
 ```bash
 go test
 ```

* [reflect.TypeOf()](https://pkg.go.dev/reflect#TypeOf) allows you to know the type of a structure. (method Name for the name)
* To avoid exposing the detail of a task, we can create a function which creates the task, NewPrint for example, and we change the visibility of Print
* As a reminder, when the name of a structure begins with a capital letter, it is public, otherwise private