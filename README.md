# goVGOexample
Simple VGO example in GO

This took a bit to get it up and running in Jetbrains GoLand.
The solution was the following: 
1. make sure the checkbox 
[X] Enable Go Modules (vgo) integration 
beneath: Settings/Go/Go Modul (vgo) is enabled. 
2. make sure you have NO entries 
beneath: Settings/GOPATH/Project GOPATH THIS IS CRUTIAL!!!! 
and make sure both checkboxes are checked: 
[X] Use GOPATH thats defined in system environment 
[X] Index entire GOPATH

This make it finally work fo me
