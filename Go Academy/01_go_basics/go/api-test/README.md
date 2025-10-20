| Endpoint                          | Method | Description                                               |
|-----------------------------------|--------|-----------------------------------------------------------|
| `/movies`                         | GET    | Retrieve a list of all movies                             |
| `/movies`                         | POST   | Create a new movie with title and release year            |
| `/movies`                         | DELETE | Delete a movie by its ID (passed as query parameter)      |
| `/characters`                     | GET    | Retrieve a list of all characters                         |
| `/characters`                     | POST   | Create a new character with name, description, and movie  |
| `/characters`                     | PUT    | Update an existing character’s details                    |
| `/characters/{id}`                | DELETE | Delete a character by their unique ID                     |
| `/appearances`                    | POST   | Link a character to a movie (record their appearance)     |
| `/characters/by-movie`            | GET    | Get characters that appear in a movie by its title        |
| `/movies/by-character`            | GET    | Get movies in which a character appears by their name     |
| `/certificates`                   | GET    | Retrieve a list of all issued certificates                |