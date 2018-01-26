
## ToDo

Data
- Create a default data set including a user, system, and movements.
-- Notes: Don't need transitions because they are part of a system.

Security
- Lock down the api's so they can only updated by the owner.
-- Movements
 
REST API
 x - Create, update, and delete and movement.
- Create, update, and delete a system.
- Find Systems by User
- Create, update, and delete a user.
- Look into mgo transactional settings.
- Add created and last updated to the system and movements table.

Application
- List view
- Create, update, delete a movement form.
- System visualization.
- Create a new transition.
- Create a transition to a new movement.
- Edit an existing transition.
- Delete an existing transition.
- Authentication
- Authorization

Common
- Add a 500 error middleware handler.

Database
- Backups
- Import/Export

## Data Model

### Users

``` json
{
  "name": "",
    "email": "",
    "enabled": true
}
```

### Movements

``` json
{
  "name": "",
    "description": "",
    "details": ""
}
```

### Systems

``` json
{
  "name": "",
  "description": "",
  "transitions": [
  {
    "name": "",
    "parent": "",
    "child": ""
  }
  ]
}
```

