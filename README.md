# Authorization in Cerbos for Gorilla Applications - Sample: Article App

## Overview

This project demonstrates the implementation of authorization in a Gorilla application using Cerbos. The sample application is an article app that showcases the integration of Cerbos for securing access to resources and managing access control policies.

## Features

- Implementation of Cerbos authorization in a Gorilla application
- Sample application: Article App
- Demonstration of access control policies based on attributes and resources
- Integration of Cerbos Policy Decision Point server for secure access control

## Getting Started

To get started with this project, you can follow these steps:

1. Clone the repository:
   ```
   git clone https://github.com/Authorization-in-Cerbos.git
   ```

2. Navigate to the project directory:
   ```
   cd Authorization-in-Cerbos
   ```

3. Install the required dependencies:
   ```
   go get -u github.com/gorilla/mux
   go get -u github.com/gorilla/sessions
   go get -u github.com/cerbos/cerbos-sdk-go/cerbos
   ```

4. Build the application:
   ```
   go build
   ```

5. Run the application:
   ```
   go run main.go
   ```

## Testing

To test the authorization implementation, you can use tools like Postman or other API testing tools to send requests to the API endpoints and validate the responses.

## Additional Tools for Testing API Calls and Deployment of Go Files

In addition to Postman, developers can leverage tools like Insomnia, Paw, or cURL for testing API calls. These tools provide a user-friendly interface for sending requests, inspecting responses, and debugging API interactions.
