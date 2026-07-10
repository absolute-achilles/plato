import { generateCourses } from "./generators"
import type { Course } from "../types"

export const allCourses: Course[] = generateCourses(24)

// Override first 4 courses to match original demo courses
allCourses[0] = {
  ...allCourses[0],
  id: "course-1",
  teacherId: "user-teacher-1",
  teacherName: "Budi Santoso",
  name: "Mathematics 101",
  description:
    "Fundamental algebra, geometry, and number theory for early secondary students.",
  category: "Mathematics",
  color: "#0D9488",
  thumbnailUrl: "https://placehold.co/600x400/0D9488/FFFFFF?text=Math",
  level: "beginner",
  status: "published",
  createdAt: new Date("2023-07-01T08:00:00Z"),
  updatedAt: new Date("2024-01-10T08:00:00Z"),
}

allCourses[1] = {
  ...allCourses[1],
  id: "course-2",
  teacherId: "user-teacher-1",
  teacherName: "Budi Santoso",
  name: "Physics Fundamentals",
  description:
    "Introduction to mechanics, energy, and waves with hands-on experiments.",
  category: "Science",
  color: "#D97706",
  thumbnailUrl: "https://placehold.co/600x400/D97706/FFFFFF?text=Physics",
  level: "beginner",
  status: "published",
  createdAt: new Date("2023-07-15T08:00:00Z"),
  updatedAt: new Date("2024-01-12T08:00:00Z"),
}

allCourses[2] = {
  ...allCourses[2],
  id: "course-3",
  teacherId: "user-teacher-1",
  teacherName: "Budi Santoso",
  name: "Bahasa Indonesia",
  description:
    "Grammar, reading comprehension, and writing skills in Indonesian language.",
  category: "Language",
  color: "#2563EB",
  thumbnailUrl: "https://placehold.co/600x400/2563EB/FFFFFF?text=BI",
  level: "beginner",
  status: "published",
  createdAt: new Date("2023-08-01T08:00:00Z"),
  updatedAt: new Date("2024-01-15T08:00:00Z"),
}

allCourses[3] = {
  ...allCourses[3],
  id: "course-4",
  teacherId: "user-teacher-1",
  teacherName: "Budi Santoso",
  name: "English Grammar",
  description:
    "Tenses, sentence structure, and vocabulary building for daily communication.",
  category: "Language",
  color: "#7C3AED",
  thumbnailUrl: "https://placehold.co/600x400/7C3AED/FFFFFF?text=English",
  level: "beginner",
  status: "published",
  createdAt: new Date("2023-08-15T08:00:00Z"),
  updatedAt: new Date("2024-01-18T08:00:00Z"),
}

export const courses: Course[] = allCourses

export function getCourseById(id: string): Course | undefined {
  return courses.find((c) => c.id === id)
}

export function getCoursesByTeacher(teacherId: string): Course[] {
  return courses.filter((c) => c.teacherId === teacherId)
}

export function getCoursesByIds(ids: string[]): Course[] {
  return courses.filter((c) => ids.includes(c.id))
}
