CREATE TYPE attendance_status AS ENUM ('present', 'absent', 'late', 'excused');

CREATE TABLE IF NOT EXISTS attendances (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  student_id UUID NOT NULL,
  module_id UUID NOT NULL,
  status attendance_status NOT NULL,
  recorded_at TIMESTAMPTZ DEFAULT now(),
  notes TEXT,

  CONSTRAINT fk_student
    FOREIGN KEY (student_id)
    REFERENCES students(user_id)
    ON DELETE CASCADE,

  CONSTRAINT fk_module
    FOREIGN KEY (module_id)
    REFERENCES modules(id)
    ON DELETE CASCADE,

  UNIQUE (student_id, module_id)
);
