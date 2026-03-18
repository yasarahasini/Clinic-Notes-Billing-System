"use client"
import { useState } from "react";
import apiClient from "../api/apiClient";

export default function UnifiedInput() {
  const [text, setText] = useState("");
  const [parsed, setParsed] = useState<any>({ drugs: [], lab_tests: [], notes: [] });

  const handleSubmit = async () => {
    const response = await apiClient.post("/api/unified-input", { text });
    setParsed(response.data);
  };

  return (
    <div className="p-4 max-w-2xl mx-auto">
      <textarea
        className="w-full h-40 p-2 border rounded"
        placeholder="Enter or speak notes, drugs, and tests..."
        value={text}
        onChange={(e) => setText(e.target.value)}
      />
      <button onClick={handleSubmit} className="mt-2 bg-blue-500 text-white px-4 py-2 rounded">
        Parse Input
      </button>

      <div className="mt-6">
        <h3 className="font-bold">Drugs:</h3>
        <ul>{parsed.drugs.map((d: string, i: number) => <li key={i}>{d}</li>)}</ul>

        <h3 className="font-bold mt-4">Lab Tests:</h3>
        <ul>{parsed.lab_tests.map((l: string, i: number) => <li key={i}>{l}</li>)}</ul>

        <h3 className="font-bold mt-4">Notes:</h3>
        <ul>{parsed.notes.map((n: string, i: number) => <li key={i}>{n}</li>)}</ul>
      </div>
    </div>
  );
}