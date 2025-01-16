# Backend_BTS_Technical_Test_Benni_Tampubolon

##  Authentication Routes

### 1. Register (POST /register) 

Method: POST * URL: http://localhost:8081/register

Content-Type: application/json * Body (raw/JSON):

`{

    "username": "testuser",
    
    "password": "testpassword",
    
    "email": "test@example.com"
    
}`

### 2. Login (POST /login) 

Method: POST * URL: http://localhost:8081/login  

Headers:

Content-Type: application/json * Body (raw/JSON):
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

