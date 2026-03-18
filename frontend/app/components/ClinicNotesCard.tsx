export default function ClinicNotesCard({ text }: { text: string }) {
  return (
    <div className="border p-2 rounded bg-yellow-50">
      {text}
    </div>
  );
}