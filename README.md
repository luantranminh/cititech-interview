## Team management app

Written in Golang, clean architecture, RESTful API.

Time spent: 16h

#### Prerequisite
- Already setup: Golang, docker, docker-compose.

### 1. Architecture
Based on this article [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) of Uncle Bob

### 2. How to start
- Clone this project, read [this article](https://luantranminh.github.io/blog/l%C3%A0m-sao-%C4%91%C3%B3ng-g%C3%B3p-cho-1-project-vi%E1%BA%BFt-b%E1%BA%B1ng-go-v%E1%BB%9Bi-git/) for better experiment.
- Install golang dependency management tool `go dep` by `go get -u github.com/golang/dep/cmd/dep`
- `dep ensure`: install the project's dependencies.
- `make local-db`: To create local database(port: 5432), test database(port:5439).
- `make run`: To start server (default port is 8000). Your server now available at **localhost:8000**
- `make unit-test`: To run test.

### 3. List of API
- **GET** `http://35.240.149.142:8080/projects/{uuidv4}`: return project with corresponding id.
- **POST** `http://35.240.149.142:8080/projects/`: create new project, body request form:
    ```json
    {
	    "project": {
		    "name": "name of this project"
	    }
    }
    ```

- **POST**   `http://35.240.149.142:8080/members`: create new member, body request form:
  ```json
    {
	    "member": {
		    "name": "Trần Minh Luân",
		    "phone": "+84332275305"
	    }
    }
  ```
- **POST**  `http://35.240.149.142:8080/members/assignments`: assign a member to project, body request form:
  ```json
    {
	    "assignment": {
		    "member_id": "{uuidv4}",
		    "project_id": "{uuidv4}"
	    }
    }
  ```

### 3. To do 
- [x] Add a member (name, phone).
- [x] Add a project (name).
- [x] Validate input.
- [x] Get project's detail (name and member belong this project).
- [x] Assign a member to a project.
- [x] Add mock for testable code.
- [ ] Unit test for "member".
- [x] Unit test for "project".