-- 1. Create the ENUM for the user roles
CREATE TYPE user_role AS ENUM ('admin', 'student', 'parent', 'teacher');

-- 2. Create the base User table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    hash_password VARCHAR(255) NOT NULL,
    role user_role NOT NULL,

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
