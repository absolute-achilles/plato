import { generateModules } from "./generators"
import { allCourses } from "./courses"
import type { Module, ModuleContent } from "../types"

export const allModules: Module[] = generateModules(60, allCourses)

// Override first 8 modules to align with original demo names
const originalModuleNames = [
  "Introduction to Algebra",
  "Linear Equations",
  "Geometry Basics",
  "Motion and Forces",
  "Energy and Work",
  "Tata Bahasa Dasar",
  "Membaca Naratif",
  "Present Tenses",
  "Past Tenses",
]

originalModuleNames.forEach((name, index) => {
  if (allModules[index]) {
    allModules[index] = {
      ...allModules[index],
      name,
      courseId: allCourses[index % 4]?.id ?? allModules[index].courseId,
      courseName: allCourses[index % 4]?.name ?? allModules[index].courseName,
      isPublished: true,
      status: "published",
    }
  }
})

export const modules: Module[] = allModules

export const moduleContents: ModuleContent[] = modules
  .slice(0, 8)
  .map((module, index) => ({
    id: `mc-${module.id}`,
    moduleId: module.id,
    title: `${module.name} Content`,
    type: index % 2 === 0 ? "lesson" : "assignment",
    bodyContent: `Content for ${module.name}`,
    position: 1,
    isPublished: true,
    createdAt: module.createdAt,
    updatedAt: module.updatedAt,
  }))

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
