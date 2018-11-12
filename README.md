# goVGOexample
Simple VGO example in GO

This took a bit to get it up and running in Jetbrains GoLand.
The solution was the following: 

1. Make sure the checkbox 
beneath: Settings/Go/Go Modul (vgo) is enabled:
- [X] Enable Go Modules (vgo) integration 

 
2. Make sure you have NO entries 
beneath: Settings/GOPATH/Project GOPATH THIS IS CRUTIAL!!!! 
and make sure both checkboxes are checked: 

- [X] Use GOPATH thats defined in system environment 
- [X] Index entire GOPATH

This make it finally work fo me
