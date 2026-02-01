'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import Navbar from '@/components/Navbar'

export default function Projects() {
  const [projects, setProjects] = useState([])
  const [loading, setLoading] = useState(false)
  const router = useRouter()

  useEffect(() => {
    const email = localStorage.getItem('userEmail')
    if (!email) {
      router.push('/login')
      return
    }
    const savedProjects = localStorage.getItem('projects')
    if (savedProjects) {
      setProjects(JSON.parse(savedProjects))
    }
  }, [router])

  return (
    <>
      <Navbar />
      <div className="container">
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '2rem' }}>
          <h1 style={{ fontSize: '2.5rem' }}>Projects</h1>
          <button
            onClick={() => router.push('/projects/create')}
            className="btn btn-primary"
          >
            Create Project
          </button>
        </div>

        {projects.length === 0 ? (
          <div className="card" style={{ textAlign: 'center', padding: '3rem' }}>
            <p style={{ color: 'var(--text-secondary)', fontSize: '1.125rem' }}>
              No projects found. Create your first project!
            </p>
          </div>
        ) : (
          <div className="grid grid-2">
            {projects.map((project, index) => (
              <div key={index} className="card">
                <h3 style={{ fontSize: '1.5rem', marginBottom: '0.5rem' }}>{project.name}</h3>
                <p style={{ color: 'var(--text-secondary)', marginBottom: '1rem' }}>
                  {project.description}
                </p>
                <div style={{ display: 'flex', gap: '0.5rem', flexWrap: 'wrap' }}>
                  <span className="badge badge-primary">Status: {project.status}</span>
                  {project.deadline && (
                    <span className="badge badge-success">Due: {project.deadline}</span>
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
