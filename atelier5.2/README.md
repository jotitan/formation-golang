# Workshop 5.2
### Create a REST API with the GIN library

#### Objectives

* Know how to use the GIN Gonic library to simplify Rest API development

#### Steps

Remove all dependencies on the standard library to use the GIN library

#### Help

* [GIN Library](github.com/gin-gonic/gin) to add
* You can keep the wrapServer and only change the handler to that of Gin
* To create an endpoint with a parameter in the url, the syntax is: "/tasks/:id"
* To return an error with a particular http code, the **AbortWithError** function will help you
* To properly handle 405 (method not allowed) errors, configure the **HandleMethodNotAllowed** property of the engine to true