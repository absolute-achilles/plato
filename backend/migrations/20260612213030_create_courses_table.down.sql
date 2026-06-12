DROP TABLE IF EXISTS enrollments;

DROP INDEX IF EXISTS idx_module_course_position ON modules(course_id, position);

DROP TABLE IF EXISTS content_attachments;

DROP INDEX IF EXISTS idx_content_module_position ON module_contents(module_id, position);

DROP TABLE IF EXISTS module_contents;

DROP TABLE IF EXISTS modules;

DROP TABLE IF EXISTS courses;

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
