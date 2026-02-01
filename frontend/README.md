# Task Manager - Next.js Frontend

A modern, full-stack task management application with a Next.js frontend and Go backend.

## Features

- **Authentication**: Login and Register pages with backend integration
- **User Management**: Create and view all users
- **Task Management**: Create and view all tasks with user assignment
- **Project Management**: Create and view projects with status tracking
- **Modern UI**: Dark theme with gradients, animations, and responsive design

## Project Structure

```
frontend/
├── src/
│   ├── app/
│   │   ├── dashboard/          # Dashboard with statistics
│   │   ├── login/              # Login page
│   │   ├── register/           # Registration page
│   │   ├── users/              # User list and create pages
│   │   ├── tasks/              # Task list and create pages
│   │   ├── projects/           # Project list and create pages
│   │   ├── globals.css         # Global styles and design system
│   │   ├── layout.js           # Root layout
│   │   └── page.js             # Home page with auto-routing
│   └── components/
│       └── Navbar.js           # Navigation component
```

## API Endpoints Used

### Authentication
- `POST /api/loguser/register` - Register new account
- `POST /api/loguser/login` - Login to account

### Users
- `POST /api/user/saveUser` - Create new user
- `GET /api/user/getAllUsers` - Get all users
- `GET /api/user/getUserByEmail/:email` - Get user by email

### Tasks
- `POST /api/task/saveTask` - Create new task
- `GET /api/task/getAllTasks` - Get all tasks
- `GET /api/task/getTaskByID/:id` - Get task by ID

## Setup Instructions

### Prerequisites
- Node.js 18+ installed
- Go backend running on `http://localhost:8080`
- MongoDB instance connected to backend

### Installation

1. Navigate to the frontend directory:
```bash
cd frontend
```

2. Install dependencies:
```bash
npm install
```

3. Start the development server:
```bash
npm run dev
```

The application will be available at `http://localhost:3000`

**Note**: If you encounter a permission error on port 3000, you can run on a different port:
```bash
PORT=3001 npm run dev
```

Or update the CORS configuration in the backend to allow the new port.

### Backend Setup

1. Navigate to the backend directory:
```bash
cd backend
```

2. Install the CORS package:
```bash
go get github.com/gin-contrib/cors
```

3. Start the backend server:
```bash
go run main.go
```

The backend will run on `http://localhost:8080`

## Usage Flow

1. **Register**: Create a new account at `/register`
2. **Login**: Sign in at `/login`
3. **Dashboard**: View statistics and quick actions
4. **Create User**: Add users to the system
5. **Create Task**: Create tasks and assign them to users
6. **Create Project**: Manage projects with status and deadlines

## Design Features

- **Modern Dark Theme**: Sleek dark background with vibrant gradients
- **Smooth Animations**: Hover effects and transitions
- **Responsive Layout**: Works on all screen sizes
- **Premium UI**: Card-based design with glassmorphism effects
- **Color Palette**: 
  - Primary: Indigo (#6366f1)
  - Secondary: Purple (#8b5cf6)
  - Background: Slate (#0f172a)
  - Success: Green (#10b981)
  - Error: Red (#ef4444)

## Technologies

- **Next.js 16**: React framework with App Router
- **React 19**: Latest React version
- **CSS**: Custom design system with CSS variables
- **Inter Font**: Modern typography from Google Fonts

## Notes

- Projects are stored in localStorage (frontend only)
- Users and Tasks are persisted in MongoDB via the backend
- Authentication state is managed with localStorage
- All pages have authentication guards

## Build for Production

```bash
npm run build
npm start
```

This will create an optimized production build.
