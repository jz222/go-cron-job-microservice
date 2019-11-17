# Cron Job Microservice

A microservice that manages cron jobs written in Go with the Gin HTTP framework and MongoDB for persisting jobs. The service sends webhooks to given systems at a given frequency. The service exposes three endpoints to add and delete cron jobs and receive their status. The service was written as POC to compare performance with an equivalent service written NodeJs.

## Usage

Execute `go get` in the root folder to install all dependencies and `go run main.go` to start the server.

## API

# `POST` `/add`
Saves and launches a new cron job.

| Property    | Type   | Explanation                                                                                                                                | Required |
|-------------|--------|--------------------------------------------------------------------------------------------------------------------------------------------|----------|
| frequency   | string | Determines when the cron job should run. Accepts valid cron syntax.                                                                        | yes      |
| url         | string | The URL the cron job should send a request to.                                                                                             | yes      |
| headers     | array  | The headers will be attached to the requests. Accepts an array of objects with the properties `key` and `value`.                           | no       |
| parameter   | string | Will be attached to the URL that the request is sent to.                                                                                   | no       |
| payload     | object | Will be sent as body of the request. Accepts valid JSON.                                                                                   | no       |
| method      | string | Determines the method of the request. Accepts "GET", "POST", "UPDATE" and "DELETE". If not provided, it will use the default value "POST". | no       |
| environment | string | Applies environment specific settings.                                                                                                     | no       |

*Example Request:*

```json
{
	"frequency": "* * * * *",
	"url": "https://webhook.example.com",
	"payload": {
		"data": "example data"
	}
}
```

*Example Response:*

```json
{
  "cronID": "5dd170ae13c0d3f137b2bec2"
}
```

# `DELETE` `/delete/:id`

Removes a cron job with the given ID from the queue and deletes it from the database.

*Example Response:*

```json
{
  "ok": true,
  "stopped": true,
  "deletedFromDatabase": true
}
```

# `GET` `/status/:id`

Returns the status for a cron job with the given ID.

*Example Response:*

```json
{
  "ok": true,
  "loadedAndRunning": true,
  "persistedInDatabase": true
}
```