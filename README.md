# macaron-sample

A prototype to see how it feels to work with Macaron and MongoDB.  
It's some kind of simple web service which handles CRUD operations over HTTP.  

    $ docker-compose up --build

Create an entry.  

    $ curl -H "Content-Type: application/json" -X POST -d '{"name":"Macaron"}' http://localhost:4000/api/persons

## Feedback
Star this repo if you found it useful. Use the github issue tracker to give feedback on this repo.
