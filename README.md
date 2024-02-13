# EventHub

This project is a Gin-based web application designed to facilitate event creation and registration. It provides a set of RESTful APIs for managing events, allowing users to create new events, view existing events, and register for events.

## Routes

- **POST /signup**: User Signup
- **POST /login**: User Login
- **POST /events**: Create a new event
- **GET /events**: Retrieve all events
- **GET /events/:id**: Retrieve a specific event by ID
- **POST /events/:id/register**: Register for a specific event
- **DELETE /events/:id**: Delete a specific event
- **DELETE /events/:id/register**: Cancel registration for a specific event

## How to Install

To install and run this project locally, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/event-registration.git
   ```

2. Navigate to the project directory:

   ```bash
   cd event-registration
   ```

3. Install dependencies:

   ```bash
   go mod download
   ```

## How to Start the Project

To start the project, run the following command:

```bash
go run .
```

The server will start running on `localhost:8080` by default.

## What It Does

This project provides a backend system for creating and managing events. It allows users to perform the following actions:

- Create new events by providing details such as event name, date, location, and description.
- View a list of all existing events.
- Retrieve details of a specific event by its unique ID.
- Register for an event by providing necessary information.
- Delete an event if it's no longer needed.

This application can serve as a foundation for building a comprehensive event management system, offering flexibility and scalability for various event-based applications.
