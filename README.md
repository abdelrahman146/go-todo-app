# Todo App API Documentation
My First Go Application

## API Endpoints

### Create Todo
- **Endpoint:** POST /todos
- **Description:** Creates a new todo item.
- **Request Body:** JSON object containing the todo item's title and status.
- **Response:** JSON object with the newly created todo item.

### Fetch All Todos
- **Endpoint:** GET /todos
- **Description:** Retrieves all todo items.
- **Response:** JSON array of all todo items.

### Fetch Single Todo
- **Endpoint:** GET /todos/:id
- **Description:** Retrieves a single todo item based on its ID.
- **Response:** JSON object of the requested todo item.

### Update Todo
- **Endpoint:** PUT /todos/:id
- **Description:** Updates the details of an existing todo item.
- **Request Body:** JSON object containing the updated title and/or status.
- **Response:** JSON object with the updated todo item.

### Delete Todo
- **Endpoint:** DELETE /todos/:id
- **Description:** Deletes a todo item based on its ID.
- **Response:** Confirmation message.

### Mark Todo as Done
- **Endpoint:** PUT /todos/:id/mark-done
- **Description:** Marks a todo item as done.
- **Response:** JSON object with the updated todo item.

### Mark Todo as Undone
- **Endpoint:** PUT /todos/:id/mark-udone
- **Description:** Marks a todo item as undone.
- **Response:** JSON object with the updated todo item.

### Delete All Todos
- **Endpoint:** DELETE /todos
- **Description:** Deletes all todo items.
- **Response:** Confirmation message.
