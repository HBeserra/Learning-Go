

Creating a package 

`go mod init <package_name>`

Use the github link of the repository as the package name like `go mod github.com/HBeserra/Learning-Go`, in this case I will use `workspace` as the package name. 


**Don't forget**
  Only methods starting with capital letters can be imported! So `workspace.Hello()` work but `workspace.hello()` will not.