"use client";
import { useState } from "react";
import apiClient from "../api/apiClient";

export default function UnifiedInput() {
  const [text, setText] = useState("");
  const [parsed, setParsed] = useState<any>({ drugs: [], lab_tests: [], notes: [] });
  const [loading, setLoading] = useState(false);

  const handleSubmit = async () => {
    setLoading(true);
    try {
      const response = await apiClient.post("/api/unified-input", { text });
      setParsed(response.data);
    } catch (error) {
      console.error("Error parsing input:", error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="p-4 max-w-2xl mx-auto border rounded-lg shadow-sm bg-white">
      <textarea
        className="w-full h-40 p-2 border text-gray-900 rounded focus:ring-2 focus:ring-blue-500 outline-none"
        placeholder="Enter or speak notes, drugs, and tests..."
        value={text}
        onChange={(e) => setText(e.target.value)}
      />
      <button 
        onClick={handleSubmit} 
        disabled={loading}
        className="mt-2 bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded transition"
      >
        {loading ? "Parsing..." : "Parse Input"}
      </button>

      <div className="mt-6 space-y-4">
        <div>
          <h3 className="font-bold text-red-600">Drugs:</h3>
          <ul className="list-disc pl-5">{parsed.drugs?.map((d: string, i: number) => <li key={i}>{d}</li>)}</ul>
        </div>

        <div>
          <h3 className="font-bold text-green-600">Lab Tests:</h3>
          <ul className="list-disc pl-5">{parsed.lab_tests?.map((l: string, i: number) => <li key={i}>{l}</li>)}</ul>
        </div>

        <div>
          <h3 className="font-bold text-gray-700">Notes:</h3>
          <ul className="list-disc pl-5">{parsed.notes?.map((n: string, i: number) => <li key={i}>{n}</li>)}</ul>
        </div>
      </div>
    </div>
  );
}