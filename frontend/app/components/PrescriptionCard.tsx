export default function PrescriptionCard({ text }: { text: string }) {
  return (
    <div className="border p-2 rounded bg-red-50">
      {text}
    </div>
  );
}