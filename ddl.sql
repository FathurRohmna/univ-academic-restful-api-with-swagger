CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE students (
    student_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    address TEXT NOT NULL,
    date_of_birth DATE NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE professors (
    professor_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    address TEXT NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE departments (
    department_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE courses (
    course_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    department_id UUID NOT NULL REFERENCES departments(department_id) ON DELETE CASCADE,
    credits INT NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE enrollments (
    student_id UUID NOT NULL REFERENCES students(student_id) ON DELETE CASCADE,
    course_id UUID NOT NULL REFERENCES courses(course_id) ON DELETE CASCADE,
    enrollment_date DATE NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,

    PRIMARY KEY (student_id, course_id)
);

CREATE TABLE teachings (
    professor_id UUID NOT NULL REFERENCES professors(professor_id) ON DELETE CASCADE,
    course_id UUID NOT NULL REFERENCES courses(course_id) ON DELETE CASCADE,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,

    PRIMARY KEY (professor_id, course_id)
);

----------------------------

INSERT INTO departments (name, description) VALUES
('Ilmu Komputer', 'Departemen Ilmu Komputer'),
('Matematika', 'Departemen Matematika');

INSERT INTO courses (name, description, department_id, credits) VALUES
('Struktur Data', 'Mempelajari pengorganisasian data', 
    (SELECT department_id FROM departments WHERE name = 'Ilmu Komputer'), 4),
('Kalkulus', 'Konsep-konsep matematika tingkat lanjut', 
    (SELECT department_id FROM departments WHERE name = 'Matematika'), 3);

INSERT INTO professors (first_name, last_name, email, address) VALUES
('Fathur', 'Rohman', 'fr@gmail.com', 'Jalan haleluyah'),
('Lisa', 'Manoban', 'jlisa@gmail.com', 'Jalan Inkar janji');

INSERT INTO students (first_name, last_name, email, address, date_of_birth, password_hash) VALUES
('Alex', 'Manhattam', 'alex.@gmail.com', 'Jalan Pine', '2000-05-15', '$2a$10$C4c7kc9eQwdjdlzcLRYTs.2hcmJKNY3K9YBw0B92SQ2A7oeKWxmKO'),
('Josh', 'Brown', 'josh@gmail.com', 'Jalan Maple', '1999-03-10', '$2a$10$C4c7kc9eQwdjdlzcLRYTs.2hcmJKNY3K9YBw0B92SQ2A7oeKWxmKO');

INSERT INTO enrollments (student_id, course_id, enrollment_date) VALUES
((SELECT student_id FROM students WHERE first_name = 'Alex'), 
    (SELECT course_id FROM courses WHERE name = 'Struktur Data'), '2024-01-15'),
((SELECT student_id FROM students WHERE first_name = 'Josh'), 
    (SELECT course_id FROM courses WHERE name = 'Kalkulus'), '2024-01-16');

INSERT INTO teachings (professor_id, course_id) VALUES
((SELECT professor_id FROM professors WHERE first_name = 'Fathur'), 
    (SELECT course_id FROM courses WHERE name = 'Struktur Data')),
((SELECT professor_id FROM professors WHERE first_name = 'Lisa'), 
    (SELECT course_id FROM courses WHERE name = 'Kalkulus'));

--------------------------------------------------------

INSERT INTO enrollments (student_id, course_id, enrollment_date) VALUES
((SELECT student_id FROM students WHERE first_name = 'Alex'), 
    (SELECT course_id FROM courses WHERE name = 'Kalkulus'), '2024-01-16');