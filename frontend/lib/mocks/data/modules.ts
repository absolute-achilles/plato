import type { Module, ModuleContent } from "../types"

export const modules: Module[] = [
  // Mathematics 101
  {
    id: "mod-math-1",
    courseId: "course-1",
    name: "Introduction to Algebra",
    position: 1,
    isPublished: true,
    unlockDate: new Date("2024-02-01T08:00:00Z"),
    createdAt: new Date("2023-07-01T08:00:00Z"),
    updatedAt: new Date("2024-01-10T08:00:00Z"),
  },
  {
    id: "mod-math-2",
    courseId: "course-1",
    name: "Linear Equations",
    position: 2,
    isPublished: true,
    unlockDate: new Date("2024-02-08T08:00:00Z"),
    createdAt: new Date("2023-07-01T08:00:00Z"),
    updatedAt: new Date("2024-01-10T08:00:00Z"),
  },
  {
    id: "mod-math-3",
    courseId: "course-1",
    name: "Geometry Basics",
    position: 3,
    isPublished: true,
    unlockDate: new Date("2024-02-15T08:00:00Z"),
    createdAt: new Date("2023-07-01T08:00:00Z"),
    updatedAt: new Date("2024-01-10T08:00:00Z"),
  },
  // Physics Fundamentals
  {
    id: "mod-phys-1",
    courseId: "course-2",
    name: "Motion and Forces",
    position: 1,
    isPublished: true,
    unlockDate: new Date("2024-02-02T08:00:00Z"),
    createdAt: new Date("2023-07-15T08:00:00Z"),
    updatedAt: new Date("2024-01-12T08:00:00Z"),
  },
  {
    id: "mod-phys-2",
    courseId: "course-2",
    name: "Energy and Work",
    position: 2,
    isPublished: true,
    unlockDate: new Date("2024-02-09T08:00:00Z"),
    createdAt: new Date("2023-07-15T08:00:00Z"),
    updatedAt: new Date("2024-01-12T08:00:00Z"),
  },
  // Bahasa Indonesia
  {
    id: "mod-bi-1",
    courseId: "course-3",
    name: "Tata Bahasa Dasar",
    position: 1,
    isPublished: true,
    unlockDate: new Date("2024-02-03T08:00:00Z"),
    createdAt: new Date("2023-08-01T08:00:00Z"),
    updatedAt: new Date("2024-01-15T08:00:00Z"),
  },
  {
    id: "mod-bi-2",
    courseId: "course-3",
    name: "Membaca Naratif",
    position: 2,
    isPublished: true,
    unlockDate: new Date("2024-02-10T08:00:00Z"),
    createdAt: new Date("2023-08-01T08:00:00Z"),
    updatedAt: new Date("2024-01-15T08:00:00Z"),
  },
  // English Grammar
  {
    id: "mod-eng-1",
    courseId: "course-4",
    name: "Present Tenses",
    position: 1,
    isPublished: true,
    unlockDate: new Date("2024-02-05T08:00:00Z"),
    createdAt: new Date("2023-08-15T08:00:00Z"),
    updatedAt: new Date("2024-01-18T08:00:00Z"),
  },
  {
    id: "mod-eng-2",
    courseId: "course-4",
    name: "Past Tenses",
    position: 2,
    isPublished: true,
    unlockDate: new Date("2024-02-12T08:00:00Z"),
    createdAt: new Date("2023-08-15T08:00:00Z"),
    updatedAt: new Date("2024-01-18T08:00:00Z"),
  },
]

export const moduleContents: ModuleContent[] = [
  // Math module 1
  {
    id: "mc-math-1-1",
    moduleId: "mod-math-1",
    title: "Variables and Constants",
    type: "lesson",
    bodyContent:
      "Learn the difference between variables and constants in algebraic expressions.",
    position: 1,
    isPublished: true,
    createdAt: new Date("2023-07-01T08:00:00Z"),
    updatedAt: new Date("2024-01-10T08:00:00Z"),
  },
  {
    id: "mc-math-1-2",
    moduleId: "mod-math-1",
    title: "Quiz: Algebra Basics",
    type: "assignment",
    bodyContent: "10 multiple choice questions on variables and constants.",
    position: 2,
    isPublished: true,
    createdAt: new Date("2023-07-01T08:00:00Z"),
    updatedAt: new Date("2024-01-10T08:00:00Z"),
  },
  // Physics module 1
  {
    id: "mc-phys-1-1",
    moduleId: "mod-phys-1",
    title: "Newton's First Law",
    type: "lesson",
    bodyContent:
      "Understanding inertia and objects at rest or in uniform motion.",
    position: 1,
    isPublished: true,
    createdAt: new Date("2023-07-15T08:00:00Z"),
    updatedAt: new Date("2024-01-12T08:00:00Z"),
  },
  {
    id: "mc-phys-1-2",
    moduleId: "mod-phys-1",
    title: "Lab Report: Motion",
    type: "assignment",
    bodyContent: "Record and analyze the motion of a toy car on a ramp.",
    position: 2,
    isPublished: true,
    createdAt: new Date("2023-07-15T08:00:00Z"),
    updatedAt: new Date("2024-01-12T08:00:00Z"),
  },
  // BI module 1
  {
    id: "mc-bi-1-1",
    moduleId: "mod-bi-1",
    title: "Kata Benda dan Kata Kerja",
    type: "lesson",
    bodyContent: "Mengenal jenis kata dan penggunaannya dalam kalimat.",
    position: 1,
    isPublished: true,
    createdAt: new Date("2023-08-01T08:00:00Z"),
    updatedAt: new Date("2024-01-15T08:00:00Z"),
  },
  // English module 1
  {
    id: "mc-eng-1-1",
    moduleId: "mod-eng-1",
    title: "Simple Present",
    type: "lesson",
    bodyContent: "Habits, routines, and general truths using simple present tense.",
    position: 1,
    isPublished: true,
    createdAt: new Date("2023-08-15T08:00:00Z"),
    updatedAt: new Date("2024-01-18T08:00:00Z"),
  },
  {
    id: "mc-eng-1-2",
    moduleId: "mod-eng-1",
    title: "Exercise: Present Tenses",
    type: "assignment",
    bodyContent: "Fill in the blanks with the correct present tense form.",
    position: 2,
    isPublished: true,
    createdAt: new Date("2023-08-15T08:00:00Z"),
    updatedAt: new Date("2024-01-18T08:00:00Z"),
  },
]

export function getModulesByCourse(courseId: string): Module[] {
  return modules
    .filter((m) => m.courseId === courseId)
    .sort((a, b) => a.position - b.position)
}

export function getContentsByModule(moduleId: string): ModuleContent[] {
  return moduleContents
    .filter((c) => c.moduleId === moduleId)
    .sort((a, b) => a.position - b.position)
}
