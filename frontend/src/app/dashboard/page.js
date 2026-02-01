'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import Navbar from '@/components/Navbar'

export default function Dashboard() {
  const [userEmail, setUserEmail] = useState('')
  const [stats, setStats] = useState({ users: 0, tasks: 0 })
  const [loading, setLoading] = useState(true)
  const router = useRouter()

  useEffect(() => {
    const email = localStorage.getItem('userEmail')
    if (!email) {
      router.push('/login')
      return
    }
    setUserEmail(email)
    fetchStats()
  }, [router])

  const fetchStats = async () => {
    try {
      const [usersRes, tasksRes] = await Promise.all([
        fetch('http://localhost:8080/api/user/getAllUsers'),
        fetch('http://localhost:8080/api/task/getAllTasks')
      ])

      const users = await usersRes.json()
      const tasks = await tasksRes.json()

      setStats({
        users: Array.isArray(users) ? users.length : 0,
        tasks: Array.isArray(tasks) ? tasks.length : 0
      })
    } catch (err) {
      console.error('Failed to fetch stats:', err)
    } finally {
      setLoading(false)
    }
  }

  if (loading) {
    return (
      <>
        <Navbar />
        <div className="loading">
          <div className="spinner"></div>
        </div>
      </>
    )
  }

  return (
    <>
      <Navbar />
      <div className="container">
        <h1 style={{ fontSize: '2.5rem', marginBottom: '0.5rem' }}>Dashboard</h1>
        <p style={{ color: 'var(--text-secondary)', marginBottom: '2rem' }}>
          Welcome back, {userEmail}
        </p>

        <div className="grid grid-2">
          <div className="card">
            <h3 style={{ fontSize: '1.25rem', marginBottom: '0.5rem', color: 'var(--text-secondary)' }}>
              Total Users
            </h3>
            <p style={{ fontSize: '3rem', fontWeight: '700', background: 'linear-gradient(135deg, var(--primary) 0%, var(--secondary) 100%)', WebkitBackgroundClip: 'text', WebkitTextFillColor: 'transparent' }}>
              {stats.users}
            </p>
          </div>

          <div className="card">
            <h3 style={{ fontSize: '1.25rem', marginBottom: '0.5rem', color: 'var(--text-secondary)' }}>
              Total Tasks
            </h3>
            <p style={{ fontSize: '3rem', fontWeight: '700', background: 'linear-gradient(135deg, var(--primary) 0%, var(--secondary) 100%)', WebkitBackgroundClip: 'text', WebkitTextFillColor: 'transparent' }}>
              {stats.tasks}
            </p>
          </div>
        </div>

        <div className="card" style={{ marginTop: '2rem' }}>
          <h2 style={{ fontSize: '1.5rem', marginBottom: '1rem' }}>Quick Actions</h2>
          <div className="grid grid-3">
            <button
              onClick={() => router.push('/users/create')}
              className="btn btn-primary"
              style={{ width: '100%' }}
            >
              Create User
            </button>
            <button
              onClick={() => router.push('/tasks/create')}
              className="btn btn-primary"
              style={{ width: '100%' }}
            >
              Create Task
            </button>
            <button
              onClick={() => router.push('/projects/create')}
              className="btn btn-primary"
              style={{ width: '100%' }}
            >
              Create Project
            </button>
          </div>
        </div>
      </div>
    </>
  )
}
