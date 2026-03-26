import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Dashboard from "./pages/Dashboard";
import PrintPreview from "./pages/PrintPreview";

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Dashboard />} />
        <Route path="/print" element={<PrintPreview />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;