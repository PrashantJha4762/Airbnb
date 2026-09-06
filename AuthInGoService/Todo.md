create a go.main file and write a simple hello world program in Go.

make two folders app and config. app is used for creating the app object which we used to do in server.ts file but since there will be more line of code we have create an app folder.

Now, make 2 structs Application and Config. in Config add a filed named addr of string type

The Application struct will have a field named config of type Config.

write a method fn named Run which has an o/p type of error.It will have a pointer reciever of type Application. Inside it, create a server using &http.Server and set Addr ,Handler(rn nil) ,readtimeout, and writetimeout. Finally, call ListenAndServe on the server and return the error.

Now in main file create the instance of COnfig and Application structs named as cf and App using app.Config and app.Application. in the main fn write app.Run()

write two constructor named NewApplication and Newconfig which will return the instance of Application and Config structs respectively.

Config k lie value hoga aur Application k lie pointer.

o/p type of NewConfig will be Config and NewApplication will be *Application.NewConfig k return type m return Config aur phir addr and NewApplication k return type m &Application aur fir config

ab main.go file m change kr do app.Newconfig aur app.NewApplication call kro

make a folder inside config named env and in that folder write the load fn which loads the env variables
after that write fn called getstring,getint and getboolean

Make a router folder and then create a router object using .NewRouter fn of chi

Make a controller folder and then write the response and register this response in the router 