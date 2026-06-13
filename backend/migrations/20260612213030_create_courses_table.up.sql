CREATE TYPE module_content_type AS ENUM (
  'lesson',
  'assignment'
);

CREATE TABLE IF NOT EXISTS courses(
  id  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  teacher_id UUID,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  created_at TIMESTAMPTZ DEFAULT now(),
  updated_at TIMESTAMPTZ DEFAULT now(),

  CONSTRAINT fk_teacher
    FOREIGN KEY (teacher_id)
    REFERENCES teachers(user_id)
    ON DELETE SET NULL
  -- don't delete courses even if the teacher is deleted
);

CREATE TABLE IF NOT EXISTS modules(
  id  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  course_id UUID NOT NULL,
  name VARCHAR(255) NOT NULL,

  position FLOAT NOT NULL,

  is_published BOOLEAN NOT NULL DEFAULT false,
  unlock_date TIMESTAMPTZ, -- If NULL, it unlocks immediately upon publishing

  created_at TIMESTAMPTZ DEFAULT now(),
  updated_at TIMESTAMPTZ DEFAULT now(),

  CONSTRAINT fk_course
    FOREIGN KEY (course_id)
    REFERENCES courses(id)
    ON DELETE CASCADE
);

-- index for fast drag-and-drop sorting queries
CREATE INDEX IF NOT EXISTS idx_module_course_position ON modules(course_id, position);

CREATE TABLE IF NOT EXISTS module_contents(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  module_id UUID NOT NULL,
  title VARCHAR(255) NOT NULL,
  type module_content_type NOT NULL,

  -- The main body text (Markdown or HTML)
  body_content TEXT,

  -- drag-and-drop ordering within a specific module
  position FLOAT NOT NULL,

  -- Allows a teacher to draft a single PDF without hiding the whole module
  is_published BOOLEAN NOT NULL DEFAULT false,

  created_at TIMESTAMPTZ DEFAULT now(),
  updated_at TIMESTAMPTZ,

  CONSTRAINT fk_module
    FOREIGN KEY (module_id)
    REFERENCES modules(id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_content_module_position ON module_contents(module_id, position);

CREATE TABLE IF NOT EXISTS content_attachments(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  module_content_id UUID NOT NULL,

  -- File metadata
  name VARCHAR(255) NOT NULL, -- e.g., "Math_Worksheet_v2.pdf"
  url TEXT NOT NULL,          -- e.g., "s3://plato-lms/files/math.pdf"
  size_bytes BIGINT,          -- e.g., "100Mb"
  type VARCHAR(100),          -- e.g., "application/pdf" or "video/mp4"

  created_at TIMESTAMPTZ DEFAULT now(),

  CONSTRAINT fk_module_content
    FOREIGN KEY (module_content_id)
    REFERENCES module_contents(id)
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS enrollments(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  student_id UUID NOT NULL,
  course_id UUID NOT NULL,
  enrolled_at TIMESTAMPTZ DEFAULT now(),

  CONSTRAINT fk_student
    FOREIGN KEY (student_id)
    REFERENCES students(user_id)
    ON DELETE CASCADE,

  CONSTRAINT fk_course
    FOREIGN KEY (course_id)
    REFERENCES courses(id)
    ON DELETE CASCADE,

  UNIQUE (student_id, course_id)
);
