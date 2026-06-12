CREATE TABLE IF NOT EXISTS courses(
  id  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  teacher_id UUID,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  created_at TIMESTAMPTZ DEFAULT now(),
  updated_at TIMESTAMPTZ,

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

  CONSTRAINT fk_course
    FOREIGN KEY (course_id)
    REFERENCES courses(id)
    ON DELETE CASCADE
);

-- index for fast drag-and-drop sorting queries
CREATE INDEX IF NOT EXISTS idx_module_course_position ON modules(course_id, position);

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
