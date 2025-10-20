1. Design and create data types and functions that will store movie characters in memory and provide CRUD operations on such memory database.
 
Entities: Movie (title, year), Character (movie, name).
Functional requirments: Provide CRUD operations on entities, list all movies, list all given movie characters.

2. Create simple web application using Uber FX (following Uber FX documentation) that will expose endpoints allowing to execute functional requirements from previous project (no need to validate data or return responses - results can be printed to the console; no need to take care about concurrency issues - race conditions).

3. Design and create OpenApi spec for the endpoints from the previous project. Generate implementation from the spec. Add request validation and return results instead of printing them to the console.

For the "Star Wars" movie, when creating a character, do an http call first to check if the provided character exists in the following public api database: [https://swapi.dev/api/people/?search=]()

4. Secure previous project in order to avoid concurrency issues. Load test some endpoints with k6 tool.

5. Dockerize previous project. Create Makefile in order to facilitate developer work.

For each entity start creating its own certyficate. For each character start creating their own certyficate. Character certyficates should be signed by a related movie certyficate. Movie certyficates should be signed by an CA certyficate provided to the application on startup via env variable. CA certyficate could be created by using openssl library.
 