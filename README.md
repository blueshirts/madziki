
## ToDo

REST API
- Create, update, and delete and movement.
- Create, update, and delete a system.
- Create, update, and delete a user.

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

