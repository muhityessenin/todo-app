Here is the README content formatted in HTML:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Todo List API</title>
</head>
<body>
    <h1>Todo List API</h1>
    <p>This repository provides a RESTful API for managing a todo list. The API includes endpoints for creating, updating, deleting, and querying tasks. Below is a brief overview of the available functionality:</p>

    <h2>1. Create a New Task</h2>
    <p><strong>Endpoint:</strong> <code>POST /api/todo-list/tasks</code></p>
    <p><strong>Request Body:</strong></p>
    <pre><code>{
  "title": "Buy a book",
  "activeAt": "2023-08-04"
}</code></pre>
    <p><strong>Fields:</strong></p>
    <ul>
        <li><code>title</code> (required, max 200 characters)</li>
        <li><code>activeAt</code> (required, valid date)</li>
    </ul>
    <p><strong>Responses:</strong></p>
    <ul>
        <li><code>201 Created</code> - If the task is created successfully, returns the ID of the task.</li>
        <li><code>404 Not Found</code> - If a task with the same title and date already exists.</li>
    </ul>

    <h2>2. Update an Existing Task</h2>
    <p><strong>Endpoint:</strong> <code>PUT /api/todo-list/tasks/{ID}</code></p>
    <p><strong>Request Body:</strong></p>
    <pre><code>{
  "title": "Buy a book - High-load applications",
  "activeAt": "2023-08-05"
}</code></pre>
    <p><strong>Parameters:</strong></p>
    <ul>
        <li><code>{ID}</code> (required)</li>
    </ul>
    <p><strong>Fields:</strong></p>
    <ul>
        <li><code>title</code> (required, max 200 characters)</li>
        <li><code>activeAt</code> (required, valid date)</li>
    </ul>
    <p><strong>Responses:</strong></p>
    <ul>
        <li><code>204 No Content</code> - If the task is updated successfully.</li>
        <li><code>404 Not Found</code> - If the task with the specified ID does not exist.</li>
    </ul>

    <h2>3. Delete a Task</h2>
    <p><strong>Endpoint:</strong> <code>DELETE /api/todo-list/tasks/{ID}</code></p>
    <p><strong>Parameters:</strong></p>
    <ul>
        <li><code>{ID}</code> (required)</li>
    </ul>
    <p><strong>Responses:</strong></p>
    <ul>
        <li><code>204 No Content</code> - If the task is deleted successfully.</li>
        <li><code>404 Not Found</code> - If the task with the specified ID does not exist.</li>
    </ul>

    <h2>4. Mark Task as Done</h2>
    <p><strong>Endpoint:</strong> <code>PUT /api/todo-list/tasks/{ID}/done</code></p>
    <p><strong>Parameters:</strong></p>
    <ul>
        <li><code>{ID}</code> (required)</li>
    </ul>
    <p><strong>Responses:</strong></p>
    <ul>
        <li><code>204 No Content</code> - If the task is marked as done successfully.</li>
        <li><code>404 Not Found</code> - If the task with the specified ID does not exist.</li>
    </ul>

    <h2>5. List Tasks by Status</h2>
    <p><strong>Endpoint:</strong> <code>GET /api/todo-list/tasks</code></p>
    <p><strong>Query Parameters:</strong></p>
    <ul>
        <li><code>status</code> (optional, can be <code>active</code> or <code>done</code>. Default is <code>active</code>)</li>
    </ul>
    <p><strong>Responses:</strong></p>
    <ul>
        <li><code>200 OK</code> - Returns a list of tasks.</li>
        <li>Tasks with <code>activeAt</code> &lt;= current date for status <code>active</code>.</li>
        <li>Tasks are sorted by creation date.</li>
        <li>If the task's <code>activeAt</code> date falls on a weekend, the title is prefixed with "WEEKEND - ".</li>
        <li>If no tasks are found, returns an empty array <code>[]</code>.</li>
    </ul>

    <p><strong>Example Response:</strong></p>
    <pre><code>[
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
]</code></pre>
</body>
</html>
```
