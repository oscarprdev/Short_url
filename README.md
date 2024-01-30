# Short URL React + Golang app 

[![Deployment pipeline](https://github.com/oscarprdev/Twitter_clone/actions/workflows/workflow.yaml/badge.svg?event=pull_request)](https://github.com/oscarprdev/Short_url/actions/workflows/workflow.yaml)

![121shots_so](https://github.com/oscarprdev/Short_url/assets/94851836/8188e643-5ca9-4ad9-a673-e8425e9c8b3a)

## Description

Short url application done with React, Zustand, React Query, Typescript and Tailwind running with Golang api using PostgreSQL database and OAuth. 

In testing, it's working for unit tests with Vitest, React Testing Library, and Mock Service Worker enabling HTTP request interception and Playwright for end-to-end tests.

## Getting Started

To run the app locally, follow the steps below:

### Backend Environment Variables

Create a `.env` file in the `backend` directory using the provided example (`backend/.env.example`). Adjust the values according to your local environment.

```env
PORT=xxx
HOST=xxx
DB_URL=xxx
SECRET=xxx

// For OAuth tokens
GOOGLE_CLIENT_ID=xxx
GOOGLE_CLIENT_SECRET=xxx

GITHUB_CLIENT_ID=xxx
GITHUB_CLIENT_SECRET=xxx
```

### Install dependencies

Run the following command to install the necessary dependencies:

```
make install_app
```

### Run the application

Start the frontend and backend servers using the following commands:

```
make start_f
make start_b
```

### Api spec

Navigate to the api_spec directory and open the API documentation using:

```
cd api_spec && make serve
```



