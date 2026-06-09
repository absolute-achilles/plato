-- 1. Create the ENUM for the user roles
CREATE TYPE user_role AS ENUM ('ADMIN', 'STUDENT', 'PARENT', 'TEACHER');

-- 2. Create the base User table
CREATE TABLE IF NOT EXISTS "user" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL,
    email TEXT NOT NULL UNIQUE, 
    password VARCHAR(255) NOT NULL,
    role user_role NOT NULL,

    UNIQUE (id, role)
);

CREATE INDEX IF NOT EXISTS idx_user_role ON "user"(role);

-- 3. Create the Admin subclass table
CREATE TABLE IF NOT EXISTS admin (
    user_id UUID PRIMARY KEY,
    role user_role DEFAULT 'ADMIN' CHECK (role = 'ADMIN'),
    CONSTRAINT fk_admin_user 
        FOREIGN KEY (user_id, role) 
        REFERENCES "user"(id, role) 
        ON DELETE CASCADE
);

-- 4. Create the Student subclass table
CREATE TABLE IF NOT EXISTS student (
    user_id UUID PRIMARY KEY,
    role user_role DEFAULT 'STUDENT' CHECK (role = 'STUDENT'),
    CONSTRAINT fk_student_user 
        FOREIGN KEY (user_id, role) 
        REFERENCES "user"(id, role) 
        ON DELETE CASCADE
);

-- 5. Create the Parent subclass table
CREATE TABLE IF NOT EXISTS parent (
    user_id UUID PRIMARY KEY,
    role user_role DEFAULT 'PARENT' CHECK (role = 'PARENT'),
    CONSTRAINT fk_parent_user 
        FOREIGN KEY (user_id, role) 
        REFERENCES "user"(id, role) 
        ON DELETE CASCADE
);

-- 6. Create the Teacher subclass table
CREATE TABLE IF NOT EXISTS teacher (
    user_id UUID PRIMARY KEY,
    role user_role DEFAULT 'TEACHER' CHECK (role = 'TEACHER'),
    CONSTRAINT fk_teacher_user 
        FOREIGN KEY (user_id, role) 
        REFERENCES "user"(id, role) 
        ON DELETE CASCADE
);

