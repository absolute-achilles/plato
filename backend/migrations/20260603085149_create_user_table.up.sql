-- ENUMS
CREATE TYPE user_role AS ENUM (
  'admin',
  'student',
  'parent',
  'teacher'
);

CREATE TYPE parent_type AS ENUM (
  'father',
  'mother',
  'guardian',
  'other'
);

-- 2. Create the base User table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    hash_password VARCHAR(255) NOT NULL,
    role user_role NOT NULL,

    -- Additional informations
    phone_number VARCHAR(20),

    created_at TIMESTAMPTZ DEFAULT now(),

    UNIQUE (id, role)
);

CREATE INDEX IF NOT EXISTS idx_user_role ON users(role);

-- 3. Create the Admin subclass table
CREATE TABLE IF NOT EXISTS admins (
    user_id UUID PRIMARY KEY,
    role user_role DEFAULT 'admin' CHECK (role = 'admin'),
    CONSTRAINT fk_admin_user
        FOREIGN KEY (user_id, role)
        REFERENCES users(id, role)
        ON DELETE CASCADE
);

-- 4. Create the Student subclass table
CREATE TABLE IF NOT EXISTS students (
    user_id UUID PRIMARY KEY,
    role user_role DEFAULT 'student' CHECK (role = 'student'),
    CONSTRAINT fk_student_user
        FOREIGN KEY (user_id, role)
        REFERENCES users(id, role)
        ON DELETE CASCADE
);

-- 5. Create the Parent subclass table
CREATE TABLE IF NOT EXISTS parents (
    user_id UUID PRIMARY KEY,
    role user_role DEFAULT 'parent' CHECK (role = 'parent'),
    type parent_type,
    CONSTRAINT fk_parent_user
        FOREIGN KEY (user_id, role)
        REFERENCES users(id, role)
        ON DELETE CASCADE
);

-- 6. Create the Teacher subclass table
CREATE TABLE IF NOT EXISTS teachers (
    user_id UUID PRIMARY KEY,
    role user_role DEFAULT 'teacher' CHECK (role = 'teacher'),
    CONSTRAINT fk_teacher_user
        FOREIGN KEY (user_id, role)
        REFERENCES users(id, role)
        ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS parent_student_links(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    parent_id UUID NOT NULL,
    student_id UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),

    CONSTRAINT fk_parent
        FOREIGN KEY (parent_id)
        REFERENCES parents(user_id)
        ON DELETE CASCADE,

    CONSTRAINT fk_student
        FOREIGN KEY (student_id)
        REFERENCES students(user_id)
        ON DELETE CASCADE,

    UNIQUE(parent_id, student_id)
);
