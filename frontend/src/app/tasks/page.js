'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import Navbar from '@/components/Navbar'

export default function Tasks() {
  const [tasks, setTasks] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')
  const router = useRouter()

  useEffect(() => {
    const email = localStorage.getItem('userEmail')
    if (!email) {
      router.push('/login')
      return
    }
    fetchTasks()
  }, [router])

  const fetchTasks = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/task/getAllTasks')
      const data = await response.json()

      if (response.ok) {
        setTasks(Array.isArray(data) ? data : [])
      } else {
        setError('Failed to fetch tasks')
      }
    } catch (err) {
      setError('Failed to connect to server')
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
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '2rem' }}>
          <h1 style={{ fontSize: '2.5rem' }}>Tasks</h1>
          <button
            onClick={() => router.push('/tasks/create')}
            className="btn btn-primary"
          >
            Create Task
          </button>
        </div>

        {error && <div className="error">{error}</div>}

        {tasks.length === 0 ? (
          <div className="card" style={{ textAlign: 'center', padding: '3rem' }}>
            <p style={{ color: 'var(--text-secondary)', fontSize: '1.125rem' }}>
              No tasks found. Create your first task!
            </p>
          </div>
        ) : (
          <div className="grid grid-2">
            {tasks.map((task) => (
              <div key={task._id} className="card">
                <h3 style={{ fontSize: '1.5rem', marginBottom: '0.5rem' }}>{task.title}</h3>
                <p style={{ color: 'var(--text-secondary)', marginBottom: '1rem' }}>
                  {task.description}
                </p>
                <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                  <span className="badge badge-primary">ID: {task._id}</span>
                  {task.user_id && (
                    <span className="badge badge-success">User: {task.user_id}</span>
                  )}
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </>
  )
}
