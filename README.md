# CTC Coding Challenge
## Info
### Tech Stack:
- Backend: `Go`
- Database: `PostgreSQL`
- Frontend: `NextJS + TypeScript`

### How to use:
- Navigate to the root folder and run the command: `docker compose up --build`
- The website should be, once the docker container is up and running, at `http://localhost:3000`

#### Requirements
1. User authentication completed through `username` and `password`. Secure handling is done by password hasing using `bcrypt`. If this were a production environment, then `ssl` would also be used.
2. Users can send and receive messages in real time and those messages are displayed through the custom chat interface. That being said, I believe, for a small application such as this, sockets would be a better answer for real-time updates, but the requirement stated RESTful API.
3. Frontend is built using `NextJS` and `TypeScript`, built and served using `Nginx`
4. Backend is built using `Go` and uses REST API to handle all functions with the frontend. Error handling and data validation - to include logging - is also managed.
5. Data persistence is maintained via the database: `PostgreSQL`.
6. Docker compose is used to containerize the application and it can be run simply with `docker compose up --build` (or `docker compose up` if this is a fresh build).

#### Guidelines
1. The code follows clean code practices and documentation is provided.
2. This readme file explains all necessary details of setting up and using the application.
3. Testing: Ideally can test through docker but unsure on how to make that work properly - must look further into profiles.
4. Docker compose works as expected and clear instructions are provided above.

## Requirements
1. User Authentication:
    - [x] Implement a user sign-up and login functionality.
    - [x] Ensure secure handling of user credentials.
2. Chat Functionality:
    - [x] Users should be able to send and receive messages in real-time.
    - [x] Display messages in a chat interface.
3. Front-End:
    - [x] Build the front-end using a framework of your choice (e.g., React, Angular, Vue.js).
    - [x] The front-end should be user-friendly and responsive.
4. Back-End:
    - [x] Develop the back-end using any language you prefer (e.g., Node.js, Python, Go).
    - [x] Provide a RESTful API to handle user authentication and chat functionalities.
    - [x] Ensure proper error handling and data validation.
5. Data Persistence:
    - [x] Implement a way to persist data (e.g., messages, user information).
    - [x] You may use any database (e.g., PostgreSQL, MongoDB, MySQL).
6. Deployment:
    - [x] Use Docker Compose to containerize the application.
    - [x] Ensure the application can be easily set up and run with a single command.

## Submission Guidelines
1. Code Quality:
    - Ensure your code is clean, well-documented, and follows best practices.
    - Include comments where necessary to explain your logic.
2. Documentation:
    - Provide a README file that explains how to set up and run the application.
    - Include any assumptions made and instructions for using the application.
3. Testing:
    - Write unit tests for critical parts of your application.
        - Tests written for backend, I would prefer to have testing through docker as well but I was unable to find a perfect solution to that
    - Provide instructions on how to run the tests.
        - Navigate to the `/app/backend/test` directory then use `go test`
        - Ideally, this would all be handled through docker
4. Deployment:
    - Ensure your Docker Compose setup works as expected.
    - Provide clear instructions on how to start the application using Docker Compose.

## Evaluation Criteria
- Functionality: Does the application meet the requirements and work as expected?
- Code Quality: Is the code clean, maintainable, and well-documented?
- User Experience: Is the front-end user-friendly and responsive?
- Security: Are user credentials and data handled securely?
- Deployment: Is the application easy to set up and run using Docker Compose?
- Testing: Are there sufficient tests, and do they cover critical parts of the application?

## Deadline
- Please submit your completed project within two weeks of receiving this challenge