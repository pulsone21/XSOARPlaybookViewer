import React from "react"
import { Link } from "react-router-dom"

export default function NavBar() {
  return (
    <nav>
      <Link to="/">Home</Link>
      <Link to="/playbooks">Playbooks</Link>
      <Link to="/viewer">Viewer!</Link>
    </nav>
  )
}
