import { apiPost } from "./client"
import type {
  CreateTeacherRequest,
  ParentResponse,
  StudentResponse,
  TeacherResponse,
} from "./types"

export async function createTeacher(
  req: CreateTeacherRequest
): Promise<TeacherResponse> {
  return apiPost<TeacherResponse, CreateTeacherRequest>(
    "/api/v1/admin/teachers",
    req
  )
}

export async function createStudent(
  req: Omit<StudentResponse, "id" | "role" | "created_at" | "name"> & { password: string }
): Promise<StudentResponse> {
  return apiPost<StudentResponse, typeof req>("/api/v1/admin/students", req)
}

export async function createParent(
  req: Omit<ParentResponse, "id" | "role" | "created_at" | "name"> & {
    password: string
  }
): Promise<ParentResponse> {
  return apiPost<ParentResponse, typeof req>("/api/v1/admin/parents", req)
}
