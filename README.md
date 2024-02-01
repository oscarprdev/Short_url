# Short URL React + Golang app 

[![Deployment pipeline](https://github.com/oscarprdev/Twitter_clone/actions/workflows/workflow.yaml/badge.svg?event=pull_request)](https://github.com/oscarprdev/Short_url/actions/workflows/workflow.yaml)

![image](https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB)
![image](https://img.shields.io/badge/TypeScript-007ACC?style=for-the-badge&logo=typescript&logoColor=white)
![image](https://img.shields.io/badge/React_Query-FF4154?style=for-the-badge&logo=ReactQuery&logoColor=white)
![image](https://img.shields.io/badge/Tailwind_CSS-38B2AC?style=for-the-badge&logo=tailwind-css&logoColor=white)
![image](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![image](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![image](https://img.shields.io/badge/Playwright-45ba4b?style=for-the-badge&logo=Playwright&logoColor=white)

![121shots_so](https://github.com/oscarprdev/Short_url/assets/94851836/8188e643-5ca9-4ad9-a673-e8425e9c8b3a)

## Description

Short url application done with React, Zustand, React Query, Typescript and Tailwind running with Golang api using PostgreSQL database and OAuth authentication. 

In testing, unit tests with Vitest, React Testing Library and Playwright for end-to-end tests.

## System Design

This short URL application simplifies the process of creating and managing short URLs. Users input their desired URLs, and the app generates a concise, user-friendly short URL based on a predefined format. The shortened URL is then displayed to the user, enabling efficient storage and easy retrieval.

A key feature of the app is the ability for users to personalize their experience by managing titles associated with each URL. This customization ensures a clear and organized overview of the user's stored URLs within their profile.

Moreover, the application seamlessly handles redirections. When a user accesses a short URL, the app queries the database to retrieve the original URL associated with it, effortlessly redirecting users to the intended destination. This robust redirection mechanism ensures a smooth and expected user experience.

A simple layout of the system design:

![Captura de pantalla 2024-01-31 a las 8 25 28](https://github.com/oscarprdev/Short_url/assets/94851836/3c20ac54-efd1-4bc5-aceb-fa4b2d2de7bb)


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



