"use client"; 

import { BrowserRouter, Routes, Route } from "react-router-dom";
import Dashboard from "./pages/Dashboard";
import PrintPreview from "./pages/PrintPreview";

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Dashboard />} />
        <Route path="/print" element={<PrintPreview />} />
      </Routes>
    </BrowserRouter>
  );
}