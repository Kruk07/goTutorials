Test files:

- create-movie.http  
  Method: POST  
  Endpoint: /movies  
  Parameters: { "title": string, "year": int }  
  Description: Creates a new movie in the repository  
  Response: JSON object with the created movie

- create-character.http  
  Method: POST  
  Endpoint: /characters  
  Parameters: { "name": string }  
  Description: Creates a new character  
  Response: JSON object with the created character

- add-appearance.http  
  Method: POST  
  Endpoint: /appearance  
  Parameters: { "movie_id": uuid, "character_id": uuid }  
  Description: Links a character to a movie  
  Response: No content (204)

- list-all-movies.http  
  Method: GET  
  Endpoint: /movies/list  
  Parameters: none  
  Description: Returns all movies in the repository  
  Response: JSON array of movies

- list-all-characters.http  
  Method: GET  
  Endpoint: /characters/list  
  Parameters: none  
  Description: Returns all characters in the repository  
  Response: JSON array of characters

- get-characters-by-movie-title.http  
  Method: GET  
  Endpoint: /characters/by-movie?title={string}  
  Parameters: title (query string)  
  Description: Returns characters appearing in the specified movie  
  Response: JSON array of characters

- get-movies-by-character-name.http  
  Method: GET  
  Endpoint: /movies/by-character?name={string}  
  Parameters: name (query string)  
  Description: Returns movies featuring the specified character  
  Response: JSON array of movie titles

- update-character.http  
  Method: PUT  
  Endpoint: /characters/update  
  Parameters: { "id": uuid, "name": string }  
  Description: Updates the name of a character  
  Response: No content (204)

- delete-movie.http  
  Method: DELETE  
  Endpoint: /movies/delete?id={uuid}  
  Parameters: id (query string)  
  Description: Deletes a movie by UUID  
  Response: No content (204)

  - delete-character.http  
  Method: DELETE  
  Endpoint: /characters/delete?id={uuid}  
  Parameters: id (query string)  
  Description: Deletes a movie by UUID  
  Response: No content (204)