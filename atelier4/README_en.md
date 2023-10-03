# Workshop 4
### Test management

#### Objectives

* Simplify unit test writing with testify
* Write quality tests

#### Steps

The tests that have been used currently are very complicated to read, maintain and debug.

A good practice is to force yourself to use the diagram [Given-When-Then](https://martinfowler.com/bliki/GivenWhenThen.html) and add functions to the test.

[Testify](https://github.com/stretchr/testify) is an assertions/mock framework to simplify the writing of tests

* Add the dependency _github.com/stretchr/testify_
* Modify the test function of TestManager to use the library testify
* Refactor the test to respect tje Given-When-Then principle
* Write a unit test to automatically test that the downsizing of the photo **/resources/photo_test.jpg** works

#### Help

* To structure a code, we divide it into 3 parts :
    * GIVEN : initial conditions
    * WHEN : the action we perform
    * THEN : expected results
* [Create a temporary directory](https://pkg.go.dev/io/ioutil#TempDir)
* [Read the contents of a file](https://pkg.go.dev/io/ioutil#ReadFile)
* To check that a photo has been reduced, compare the weight of the image before / after