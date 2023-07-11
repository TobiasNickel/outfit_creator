# outfit creator
A task to verify the software development skills on the example of a webapp for creating a random outfit based on the new yorker Produkt-REST-API.

# Task
A random outfit generator is to be created as a web application and a corresponding backend. An outfit consists of three products: upper garment, under garment and accessory. Randomly one product shall be searched for and displayed respectively.

In the frontend a random outfit is loaded via a REST API call to the to implement backend endpoint. It must be ensured that men's and women's articles are not mixed. In addition to the product images, relevant master data on the selected products shall be displayed in the frontend.

 - The choice of technologies and programming languages is left to your discretion
 - The source code must be made available to NEW YORKER digital in any form
 - NEW YORKER must provide a working version (hosted or locally executable)
 - The product REST API (basis for the NEW YORKER product area Web/App) from NEW YORKER is to be used
 - The solution should be documented and subsequently presented.

# Architecture
The frontend is developed using `React` js and is build using `esbuild`. It is in the `frontend` directory. Once build, it can be found in the public directory. 

The backend is implemented using golang. As per the requiremens, I do not see a need for persistent datastore. The backend is golang project is in the root directory of this repository. This is because the golang `embed` feature is used. With this the build fontend application is embedded into the executable program, So that the application can be deployed with just this one executable.

The package.json in the backend is used for auto restart the golang server in development.

## development
In development we install two npm projects, the root and the frontend. The npm root project is used to start frontend and backend both in watch mode using a single console window.
```sh
npm install
npm install --prefix frontend

npm run watch
```

## build
For creating production builds you can use the following commands. Skip the installation if you already did for development. You find the two commands for building the go program choos one of the commands with your target operating system. (this program was developed on a windows machine, cross compiled for linux and deployed online on a virtual host). 
```sh
npm install
npm install --prefix frontend
npm run build --prefix frontend
GOOS=darwin go build -o outfitgenerator ./programs/server/server.go
GOOS=linux go build -o outfitgenerator ./programs/server/server.go
GOOS=win go build -o outfitgenerator.exe ./programs/server/server.go
```

# Frontend
As said, React and build using esbuild. It contains a `frontend.go` file for embedding the public directory. The App is rather small, so the state is held in the `App` compnent. For each Item of an outfit, a dedicated `Item` component is implemented with its own styles. 

To call the backend, a dedicated SDK module is provided. It also contains auto automatically generated type definitions. 

# backend
The backend is a golang application made using the gin framework. The Entry point to the backend server is  `programs/server/server.go`. If we wanted an extra CLI we could add this as an other program.

All other Modules are in the `internal` directory, that is a golang convention to tell that these modules only belong to this application and are not meaned to be used in other Goprograms. Within the internal directoy you find the `handler` `service` and `type` directories. These represent the structure for the MVC Framework.

Utils contains a helper to do http requests with included json parsing. And a gin middleware for serving static files. I needed to implement this, because standard solutions did not work well for Single Page Applications (did not serve the index.html, even with configuration).

# tests
The Backend is simple and consists mainly out of integrations. Thus I deem integration tests as sufficient.
To do the integration tests, the creation of the gin server is put into its own server package.
you can run the tests using:

```sh
go test ./integration
``` 

```go
go test -tags=integration ./integration
```

# API
There is only one endpoint:

```
GET /api/outfit/random
```
It can receive two paramete:
 - **gender** can be male or female (default male)
 - **country** the two letter short name for a country (default de)




