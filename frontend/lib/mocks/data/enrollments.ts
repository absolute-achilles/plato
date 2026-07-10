import type { Enrollment } from "../types"
import { student } from "./users"

export const enrollments: Enrollment[] = [
  {
    id: "enroll-1",
    studentId: student.id,
    courseId: "course-1",
    enrolledAt: new Date("2024-01-15T08:00:00Z"),
    progressPercent: 72,
  },
  {
    id: "enroll-2",
    studentId: student.id,
    courseId: "course-2",
    enrolledAt: new Date("2024-01-16T08:00:00Z"),
    progressPercent: 45,
  },
  {
    id: "enroll-3",
    studentId: student.id,
    courseId: "course-3",
    enrolledAt: new Date("2024-01-17T08:00:00Z"),
    progressPercent: 88,
  },
  {
    id: "enroll-4",
    studentId: student.id,
    courseId: "course-4",
    enrolledAt: new Date("2024-01-18T08:00:00Z"),
    progressPercent: 30,
  },
]

export function getEnrollmentsByStudent(studentId: string): Enrollment[] {
  return enrollments.filter((e) => e.studentId === studentId)
}

export function getEnrollmentsByCourse(courseId: string): Enrollment[] {
  return enrollments.filter((e) => e.courseId === courseId)
}
