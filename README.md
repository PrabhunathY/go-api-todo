# go-api-todo

## Get All TODO List
Request: http://localhost:9000/items
Request Method: GET
Response: [
    {
        "id": "1",
        "title": "Item1",
        "text": "Description"
    },
    {
        "id": "2",
        "title": "Title 2",
        "text": "Text 2"
    }
]

## Add New Item into TODO list
Request: http://localhost:9000/items
Request Method: POST
Request Body: 
    {
        "title": "Title 2",
        "text": "Text 2"
    }

Response: {
    "id": "2",
    "title": "Title 2",
    "text": "Text 2"
}

## Update an Item into TODO list
Request: http://localhost:9000/items
Request Method: PUT
Request Body: 
     {  "id":"2",
        "title": "Title 2 Edit",
        "text": "Text  updated"
    }

Response: [
    {
        "id": "1",
        "title": "Item1",
        "text": "Description"
    },
    {
        "id": "2",
        "title": "Title 2 Edit",
        "text": "Text  updated"
    }
]

## Remove an Item from TODO list
Request: http://localhost:9000/items/2
Request Method: DELETE

Response: [
    {
        "id": "1",
        "title": "Item1",
        "text": "Description"
    }
]