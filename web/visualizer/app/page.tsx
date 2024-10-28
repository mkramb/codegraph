"use client"

import { useEffect, useState } from "react"
import { GraphComponent } from "./components/graph"

export default function App() {
  const [data, setData] = useState(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch("http://localhost:3000/data.json")
      const result = await response.json()

      setData(result)
      setLoading(false)
    }

    fetchData()
  }, [])

  if (loading) return <div>Loading...</div>
  if (!data) return <div>Error loading data.</div>

  return (
    <div className="p-20">
      <GraphComponent />
    </div>
  )
}
