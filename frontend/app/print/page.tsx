"use client";
import jsPDF from "jspdf";
import Link from "next/link";

export default function PrintPreview() {
  const generatePDF = () => {
    const doc = new jsPDF();
    doc.setFontSize(20);
    doc.text("Clinic Report", 10, 20);
    doc.setFontSize(12);
    doc.text("Generated on: " + new Date().toLocaleString(), 10, 30);
    doc.save("clinic-report.pdf");
  };

  return (
    <div className="p-8 max-w-2xl mx-auto bg-white py-20">
      <Link href="/" className="text-blue-500 hover:underline mb-4 inline-block">
        ← Back to Dashboard
      </Link>
      <h1 className="text-2xl font-bold mb-6 text-black">Print Preview</h1>
      <div className="border-2 border-dashed border-gray-300 p-10 text-center rounded-lg">
        <p className="text-gray-500">PDF Preview Content Goes Here</p>
        <button 
          onClick={generatePDF} 
          className="mt-6 bg-green-600 hover:bg-green-700 text-white px-6 py-2 rounded shadow-md transition"
        >
          Download PDF Report
        </button>
      </div>
    </div>
  );
}