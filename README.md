# SecureSessionManager

A web service for managing secure user sessions with authentication and authorization functionality enabled and real-time timeouts. 

## Features

- Manage users sessions and make them secure
- Role-based authetication and authorization 
- Sessions managed by timeout timer for a 60 seconds 
- Cookie-based session information storage
- Notifies users about the session timeout at real-time 

## Technology Stack

- Go - backend functionality and http routing
- HTML - Frontend UI
- JavaScript - for enabling session timeout as static function called by the Go route
- Gorilla - library for enabling cookie-based session information storage and key-based authorization

## Running the Project

- Download the project or ``git clone`` the repo in your IDE (preferably VSCode)
- Open the terminal and run the command - ``go get github.com/gorilla/sessions``
- Run ``app.go`` and "Allow Access" in the following dialogue box 
- Start ``localhost:3002`` - you can change the port as you like
- Test the functionality 
