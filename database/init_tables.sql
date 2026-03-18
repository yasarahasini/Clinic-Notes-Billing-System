CREATE TABLE IF NOT EXISTS patients (
    id SERIAL PRIMARY KEY,
    name TEXT,
    age INT
);

CREATE TABLE IF NOT EXISTS visits (
    id SERIAL PRIMARY KEY,
    patient_id INT REFERENCES patients(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS drugs (
    id SERIAL PRIMARY KEY,
    visit_id INT REFERENCES visits(id),
    name TEXT,
    dosage TEXT,
    price NUMERIC
);

CREATE TABLE IF NOT EXISTS lab_tests (
    id SERIAL PRIMARY KEY,
    visit_id INT REFERENCES visits(id),
    test_name TEXT,
    price NUMERIC
);

CREATE TABLE IF NOT EXISTS notes (
    id SERIAL PRIMARY KEY,
    visit_id INT REFERENCES visits(id),
    content TEXT
);