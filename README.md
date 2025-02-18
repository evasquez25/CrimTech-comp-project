# Simple Counter App

This is a simple web application that demonstrates a counter service using gRPC-Web and Svelte. The server is implemented in Go, and the client is a Svelte application.

## Features
- Increment a global counter on the server.
- Multiple clients can connect and see the updated counter value.

## Prerequisites
- Go 1.16+
- Node.js 14+

## Setup

### Server
1. Navigate to the project root directory.
2. Run the server:
   ```bash
   go run server.go
   ```
   The server will start on port `8080`.

### Client
1. Navigate to the `svelte-client` directory:
   ```bash
   cd svelte-client
   ```
2. Install the dependencies:
   ```bash
   npm install
   ```
3. Start the development server:
   ```bash
   npm run dev
   ```
   The client will be served on `http://localhost:5000`.

## Usage
- Open `http://localhost:5000` in your web browser.
- Click the "Increment Counter" button to increment the counter value on the server.

## Notes
- The server uses gRPC-Web to communicate with the client.
- Ensure both server and client are running to see the application in action.
