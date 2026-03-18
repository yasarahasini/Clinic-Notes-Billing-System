import React from "react";

interface Props {
  labTests: string[];
}

const LabTestInput: React.FC<Props> = ({ labTests }) => {
  return (
    <div className="border p-3 rounded">
      <h2 className="font-semibold mb-2">Lab Tests / Investigations</h2>
      {labTests.length === 0 ? (
        <p>No lab tests detected.</p>
      ) : (
        <ul className="list-disc pl-5">
          {labTests.map((test, index) => (
            <li key={index}>{test}</li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default LabTestInput;