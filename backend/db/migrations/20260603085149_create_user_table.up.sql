-- 1. Create the ENUM for the user roles
CREATE TYPE user_role AS ENUM ('ADMIN', 'STUDENT', 'GUARDIAN', 'TEACHER');

-- 2. Create the base User table
CREATE TABLE IF NOT EXISTS "user" (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL,
    email TEXT NOT NULL UNIQUE, 
    password VARCHAR(255) NOT NULL,
    role user_role NOT NULL
);

-- 3. Create the Admin subclass table
CREATE TABLE IF NOT EXISTS admin (
    admin_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE,
    CONSTRAINT fk_admin_user 
        FOREIGN KEY (user_id) 
        REFERENCES "user"(user_id) 
        ON DELETE CASCADE
);

-- 4. Create the Student subclass table
CREATE TABLE IF NOT EXIST student (
    stud_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE,
    CONSTRAINT fk_student_user 
        FOREIGN KEY (user_id) 
        REFERENCES "user"(user_id) 
        ON DELETE CASCADE
);

-- 5. Create the Guardian subclass table
CREATE TABLE IF NOT EXISTS guardian (
    guard_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE,
    CONSTRAINT fk_guardian_user 
        FOREIGN KEY (user_id) 
        REFERENCES "user"(user_id) 
        ON DELETE CASCADE
);

-- 6. Create the Teacher subclass table
CREATE TABLE IF NOT EXISTS teacher (
    teach_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE,
    CONSTRAINT fk_teacher_user 
        FOREIGN KEY (user_id) 
        REFERENCES "user"(user_id) 
        ON DELETE CASCADE
);
