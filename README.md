# Backend_BTS_Technical_Test_Benni_Tampubolon

##  Authentication Routes

### 1. Register (POST /register) 

Method: POST * URL: http://localhost:8081/register

Content-Type: application/json * Body (raw/JSON):

Body Raw test:

`{

    "username": "testuser",
    
    "password": "testpassword",
    
    "email": "test@example.com"
    
}`

### 2. Login (POST /login) 

Method: POST * URL: http://localhost:8081/login  

Headers:

Content-Type: application/json * Body (raw/JSON):
Body Raw test:

`{

    "username": "testuser",
    
    "password": "testpassword"
    
}`

## Checklist Routes (Protected)

### 1. Get All Checklists (GET /checklists) 

Method: GET * URL: http://localhost:8081/checklists

Example Response:

[
  {
  
      "id": 1,
      
      "user_id": 1,
      
      "item_name": "My First Checklist",
      
      "created_at": "2024-01-27 12:00:00",
      
      "updated_at": "2024-01-27 12:00:00"
      
  }
  
]

### 2. Create Checklist (POST /checklists)  

Method: POST * URL: http://localhost:8081/checklists
Body Raw test:

`{
    "item_name": "My New Checklist"
}`

### 3. Get Checklist (GET /checklists/{id}) 

Method: GET * URL: http://localhost:8081/checklists/1

 Example Response:
 '
 {
 
    "id": 1,
    
    "user_id": 1,
    
    "item_name": "My First Checklist",
    
    "created_at": "2024-01-27 12:00:00",
    
    "updated_at": "2024-01-27 12:00:00"
    
}
'

### 4. Delete Checklist (DELETE /checklists/{id}

Method: DELETE * URL: http://localhost:8081/checklists/1

## 3. Item Routes (Protected)

### 1. Get Items (GET /checklists/{id}/items) 

Method: GET * URL: http://localhost:8081/checklists/1/items

Example Response:
`
[
  {
      "id": 1,
      "checklist_id": 1,
      "text": "First Item",
      "completed": false,
      "created_at": "2024-01-27 12:00:00",
      "updated_at": "2024-01-27 12:00:00"
  }
]
`

### 2. Create Item (POST /checklists/{id}/items)  

Method: POST * URL: http://localhost:8081/checklists/1/items

Body Raw test:

{

    "text": "My New Item"
    
}


### 3. Update Item (PUT /checklists/{id}/items/{item_id})  

Method: PUT * URL: http://localhost:8081/checklists/1/items/1

{
    "text": "Updated Item Text",
    "completed": true
}

### 4. Delete Item (DELETE /checklists/{id}/items/{item_id}) 

Method: DELETE * URL: http://localhost:8081/checklists/1/items/1
