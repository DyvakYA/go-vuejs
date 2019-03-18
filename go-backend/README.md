## Usage

All responses will have the form

```json
{
	"message": "Description of what happened",
	"data": "Mixed type holding the contnt of response"
}
```

Subsequent response definations will only detail the expected value of the `data field`

### List all users

**Defination**

`GET /users`

**Response**

- `200 OK` on succes

```json
[
	{
		"id": "1",
		"name": "Nockolas",
		"description": "Business Partner"
	},
	{
		"id": "2",
		"name": "Robert",
		"description": "Teammate"
	}	
]