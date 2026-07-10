CREATE TYPE IF NOT EXISTS teacher_department AS ENUM (
  'Mathematics',
  'Science',
  'English',
  'History',
  'Arts',
  'Physical Education',
  'Computer Science',
  'Other'
);

CREATE TYPE IF NOT EXISTS student_grade_level AS ENUM (
  'Grade 1',
  'Grade 2',
  'Grade 3',
  'Grade 4',
  'Grade 5',
  'Grade 6',
  'Grade 7',
  'Grade 8',
  'Grade 9',
  'Grade 10',
  'Grade 11',
  'Grade 12'
);

ALTER TABLE teachers
  ADD COLUMN IF NOT EXISTS department teacher_department NOT NULL DEFAULT 'Other';

ALTER TABLE students
  ADD COLUMN IF NOT EXISTS grade_level student_grade_level NOT NULL DEFAULT 'Grade 1';

-- Seed the first admin user.
-- Temporary password: admin12345 (must be changed immediately via /api/v1/auth/change-password).
INSERT INTO users (id, username, email, hash_password, role, phone_number)
VALUES (
  '00000000-0000-0000-0000-000000000001',
  'admin',
  'admin@plato.local',
  '$2a$10$ExW51KWIj3f8ddpF7HB8Suo1Uzgo22UOF69mBxghdAy9MIVEWyesq',
  'admin',
  NULL
)
ON CONFLICT (id) DO NOTHING;

INSERT INTO admins (user_id)
VALUES ('00000000-0000-0000-0000-000000000001')
ON CONFLICT (user_id) DO NOTHING;
