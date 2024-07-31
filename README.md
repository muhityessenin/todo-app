Todo List API
=============

This repository provides a RESTful API for managing a todo list. The API includes endpoints for creating, updating, deleting, and querying tasks. Below is a brief overview of the available functionality:

1\. Create a New Task
---------------------

**Endpoint:** `POST /api/todo-list/tasks`

**Request Body:**

    {
      "title": "Buy a book",
      "activeAt": "2023-08-04"
    }

**Fields:**

*   `title` (required, max 200 characters)
*   `activeAt` (required, valid date)

**Responses:**

*   `201 Created` - If the task is created successfully, returns the ID of the task.
*   `404 Not Found` - If a task with the same title and date already exists.

2\. Update an Existing Task
---------------------------

**Endpoint:** `PUT /api/todo-list/tasks/{ID}`

**Request Body:**

    {
      "title": "Buy a book - High-load applications",
      "activeAt": "2023-08-05"
    }

**Parameters:**

*   `{ID}` (required)

**Fields:**

*   `title` (required, max 200 characters)
*   `activeAt` (required, valid date)

**Responses:**

*   `204 No Content` - If the task is updated successfully.
*   `404 Not Found` - If the task with the specified ID does not exist.

3\. Delete a Task
-----------------

**Endpoint:** `DELETE /api/todo-list/tasks/{ID}`

**Parameters:**

*   `{ID}` (required)

**Responses:**

*   `204 No Content` - If the task is deleted successfully.
*   `404 Not Found` - If the task with the specified ID does not exist.

4\. Mark Task as Done
---------------------

**Endpoint:** `PUT /api/todo-list/tasks/{ID}/done`

**Parameters:**

*   `{ID}` (required)

**Responses:**

*   `204 No Content` - If the task is marked as done successfully.
*   `404 Not Found` - If the task with the specified ID does not exist.

5\. List Tasks by Status
------------------------

**Endpoint:** `GET /api/todo-list/tasks`

**Query Parameters:**

*   `status` (optional, can be `active` or `done`. Default is `active`)

**Responses:**

*   `200 OK` - Returns a list of tasks.
*   Tasks with `activeAt` <= current date for status `active`.
*   Tasks are sorted by creation date.
*   If the task's `activeAt` date falls on a weekend, the title is prefixed with "WEEKEND - ".
*   If no tasks are found, returns an empty array `[]`.

**Example Response:**

    [
      {
        "id": "65f19340848f4be025160391",
        "title": "Buy a book - High-load applications",
        "activeAt": "2023-08-05"
      },
      {
        "id": "75f19340848f4be025160392",
        "title": "Buy an apartment",
        "activeAt": "2023-08-05"
      },
      {
        "id": "45f19340848f4be025160394",
        "title": "Buy a car",
        "activeAt": "2023-08-05"
      }
    ]

How to run it locally
------------------------
**Required Tools: Docker**

**git clone https://github.com/muhityessenin/todo-app.git**

**make build**

**make up**

**Those commands will create docker containers and run the program**
