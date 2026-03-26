import UnifiedInput from "./components/UnifiedInput";
import Link from "next/link";

export default function Dashboard() {
  return (
    <main className="min-h-screen bg-gray-50 p-8">
      <div className="max-w-4xl mx-auto flex justify-between items-center mb-8">
        <h1 className="text-3xl font-extrabold text-gray-800">Clinic Notes & Billing</h1>
        <Link 
          href="/print" 
          className="bg-blue-800 text-white px-4 py-2 rounded hover:bg-gray-700"
        >
          Go to Print Preview
        </Link>
      </div>
      <UnifiedInput />
    </main>
  );
}