## Usage

All responses will have the form

```json
{
	"data": "Mixed type holding the contnt of response",
	"message": "Description of what happened",
}
```

Subsequent response definations will only detail the expected value of the `sdata field`

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