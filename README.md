# Learning Management System (LMS) API

A simple Learning Management System (LMS) API that allows users to manage courses, enrollments, and track student progress with role-based access for admins, instructors, and students.

## Features
- User Registration and Authentication
- Role-based access control (Admin, Instructor, Student)
- Course management (create, update, delete)
- Course enrollment for students
- Progress tracking for students

## Authentication
- JWT authentication for all roles.
- Roles: `admin`, `instructor`, `student`.

## Endpoints

### User Authentication
#### Public Routes
- `POST /api/v1/register` - Register a new user
  - Request Body:
    ```json
    {
      "name": "string",
      "email": "string",
      "password": "string",
      "role": "admin|instructor|student"
    }
    ```

- `POST /api/v1/login` - Login a user
  - Request Body:
    ```json
    {
      "email": "string",
      "password": "string"
    }
    ```

### Courses
#### Public and Protected Routes
- `GET /api/v1/courses` - Get all available courses (Public)
- `GET /api/v1/courses/{id}` - Get details of a specific course

- `POST /api/v1/admin/courses` - Create a new course (Admin/Instructor)
  - Request Body:
    ```json
    {
      "title": "string",
      "description": "string",
      "instructor_id": 1
    }
    ```

- `PUT /api/v1/admin/courses/{id}` - Update course (Instructor/Admin)
  - Request Body:
    ```json
    {
      "title": "string",
      "description": "string",
      "instructor_id": 1
    }
    ```

- `DELETE /api/v1/admin/courses/{id}` - Delete course (Instructor/Admin)

### Enrollments
#### Student Routes
- `POST /api/v1/student/courses/{id}/enroll` - Enroll a student in a course
- `GET /api/v1/student/courses/{id}/progress` - Get student progress for a specific course

## Database Schema

### Users
- `id` - Integer (Primary Key)
- `name` - String
- `email` - String (Unique)
- `password_hash` - String
- `role` - Enum (admin, instructor, student)

### Courses
- `id` - Integer (Primary Key)
- `title` - String
- `description` - String
- `instructor_id` - Integer (Foreign Key to Users)

### Enrollments
- `id` - Integer (Primary Key)
- `course_id` - Integer (Foreign Key to Courses)
- `student_id` - Integer (Foreign Key to Users)
- `progress` - Integer (Percentage)

## Authentication Flow
- Upon registration or login, users receive a JWT token.
- This token is required to access protected endpoints (courses, enrollments).
- Admins and instructors have higher privileges, such as managing courses.

## Roles
- **Admin**: Can manage all users and courses.
- **Instructor**: Can create, update, and delete courses.
- **Student**: Can enroll in courses and track progress.
