'use client'

import { useRouter } from 'next/navigation'
import Link from 'next/link'

export default function Navbar() {
  const router = useRouter()

  const handleLogout = () => {
    localStorage.removeItem('userEmail')
    router.push('/login')
  }

  return (
    <nav className="navbar">
      <div className="navbar-content">
        <div className="navbar-brand">Task Manager</div>
        <div className="navbar-links">
          <Link href="/dashboard">Dashboard</Link>
          <Link href="/users">Users</Link>
          <Link href="/tasks">Tasks</Link>
          <Link href="/projects">Projects</Link>
          <button onClick={handleLogout} className="btn btn-secondary">
            Logout
          </button>
        </div>
      </div>
    </nav>
  )
}
