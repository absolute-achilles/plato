export type Role = "student" | "teacher" | "parent" | "admin"

export type Department =
  | "Mathematics"
  | "Science"
  | "English"
  | "History"
  | "Arts"
  | "Physical Education"
  | "Computer Science"
  | "Other"

export type GradeLevel =
  | "Grade 1"
  | "Grade 2"
  | "Grade 3"
  | "Grade 4"
  | "Grade 5"
  | "Grade 6"
  | "Grade 7"
  | "Grade 8"
  | "Grade 9"
  | "Grade 10"
  | "Grade 11"
  | "Grade 12"

export type ParentType = "father" | "mother" | "guardian" | "other"

export interface APIError {
  code: string
  message: string
}

export interface APIEnvelope<T> {
  success: boolean
  data?: T
  error?: APIError
  meta?: {
    page?: number
    page_size?: number
    total_items?: number
    total_pages?: number
  }
}

export interface UserResponse {
  id: string
  username: string
  name: string
  email: string
  role: Role
  phone_number?: string
  created_at: string
}

export interface TokenResponse {
  access_token: string
  refresh_token: string
  expires_in: number
  user: UserResponse
}

export interface TeacherResponse extends UserResponse {
  department: Department
}

export interface StudentResponse extends UserResponse {
  grade_level: GradeLevel
}

export interface ParentResponse extends UserResponse {
  type: ParentType
  student_ids: string[]
}

export interface LoginRequest {
  email: string
  password: string
}

export interface CreateTeacherRequest {
  username: string
  email: string
  password: string
  phone_number?: string
  department: Department
}

export interface ChangePasswordRequest {
  old_password: string
  new_password: string
}
