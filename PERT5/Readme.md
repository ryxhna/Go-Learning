### Import package yang diperlukan
- encoding/json : konversi data ke/dari format JSON
- fmt, log      : output dan logging
- net/http      : membuat HTTP server
- os            : file handling
- strconv       : konversi string ke int

### Endpoint 
- handleFilms = http/localhost:8080/films           --> GET, POST
- handleFilmByID = http/localhost:8080/films/{id}   --> GET, PUT, DELETE
