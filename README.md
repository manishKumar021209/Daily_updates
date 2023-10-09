# Daily_updates

# Album Management API in Go

This Go application implements a simple RESTful API for managing a collection of albums. It allows you to perform basic operations such as creating, retrieving, updating, and deleting album records. The API is built using the Gorilla Mux router and uses JSON for data exchange.

## Features

- Retrieve a list of albums.
- Retrieve an album by its ID.
- Create a new album.
- Update an existing album.
- Delete an album by its ID.

## Getting Started

To run the Album Management API, follow these steps:

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/your-username/album-management-api.git
2. Change into the project directory:cd album-management-api
3.Build and run the Go program:go run main.go
4. The API will be accessible at http://localhost:3000.

API Endpoints
GET /albums: Retrieve a list of all albums.
GET /albums/{id}: Retrieve an album by its ID.
POST /albums: Create a new album.
PUT /albums/{id}: Update an existing album by its ID.
DELETE /albums/{id}: Delete an album by its ID.


Sample Usage
GET /albums
Retrieve a list of all albums.
curl http://localhost:3000/albums


GET /albums/{id}
Retrieve an album by its ID.
GET /albums/{id}
curl http://localhost:3000/albums/1



POST /albums
Create a new album.
curl -X POST -H "Content-Type: application/json" -d '{"title":"New Album","artist":"New Artist","price":"$19.99"}' http://localhost:3000/albums


PUT /albums/{id}
Update an existing album by its ID.
curl -X PUT -H "Content-Type: application/json" -d '{"title":"Updated Album","artist":"Updated Artist","price":"$24.99"}' http://localhost:3000/albums/1


DELETE /albums/{id}
Delete an album by its ID.
curl -X DELETE http://localhost:3000/albums/1


Author
Manish
