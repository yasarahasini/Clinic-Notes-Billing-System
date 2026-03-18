"use client";
import jsPDF from "jspdf";

export default function PrintPreview() {
  const generatePDF = () => {
    const doc = new jsPDF();
    doc.text("Clinic Report", 10, 10);
    doc.save("report.pdf");
  };

  return (
    <div className="p-4">
      <h1 className="text-xl font-bold">Print Preview</h1>
      <button onClick={generatePDF} className="mt-4 bg-green-500 text-white px-4 py-2 rounded">
        Generate PDF
      </button>
    </div>
  );
}