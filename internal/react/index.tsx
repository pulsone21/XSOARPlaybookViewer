import { createRoot } from "react-dom/client";
import React from "react"
import { BrowserRouter, Routes, Route } from "react-router-dom"
import Home from "./pages/home"
import Playbooks from "./pages/playbooks"
import Viewer from "./pages/viewer"



const root = createRoot(document.getElementById("root")!);


root.render(
  <BrowserRouter>
    <Routes>
      <Route index element={<Home />} />
      <Route path="/playbooks" element={<Playbooks />} />
      <Route path="/viewer" element={<Viewer />} />
    </ Routes>

  </ BrowserRouter>
);
