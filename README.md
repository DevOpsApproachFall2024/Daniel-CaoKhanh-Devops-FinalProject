# Daniel-CaoKhanh-Devops-FinalProject

A DevOps Project with the goal to implement DevOps SDLC

## App

Created a simple button that leads to pages on specfied routes

## Components

- Backend Golang
- Frontend TypeScript (React)
- Dockerfile
- Github Action

## Backend File Structure

    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── golangci.yml
    ├── main.go
    ├── routes
    │   ├── fiboRoutes.go
    │   ├── rickRoutes.go
    │   └── tonyRoutes.go
    └── tests
        ├── fibo_test.go
        ├── rick_test.go
        └── tony_test.go

## Frontend File Structure

    ├── Dockerfile
    ├── jest.config.ts
    ├── package-lock.json
    ├── package.json
    ├── public/
    ├── readme.md
    ├── src
    │   ├── App.css
    │   ├── App.test.tsx
    │   ├── App.tsx
    │   ├── Fibonacci.css
    │   ├── Fibonacci.tsx
    │   ├── FibonacciInput.tsx
    │   ├── RickVideo.tsx
    │   ├── TonyVideo.tsx
    │   ├── WelcomePage.css
    │   ├── WelcomePage.tsx
    │   ├── index.css
    │   ├── index.tsx
    │   ├── logo.svg
    │   ├── reportWebVitals.ts
    │   ├── setupTests.ts
    │   └── ytvideos.css
    └── tsconfig.json

## Github Workflow File Structure

    └── workflows
        ├── backend.yml
        ├── frontend-ci.yml
        └── semantic-versioning.yml

# How To Run 

1. Git clone the repo
2. Make sure your docker is running
3. Run the command

        docker compose up --build -d

4. And then you are good to go!!