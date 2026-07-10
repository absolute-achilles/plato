import { generateTeachers } from "./generators"
import type { Teacher } from "../types"

export const allTeachers: Teacher[] = generateTeachers(50)

// Override the first teacher to match the original demo user
allTeachers[0] = {
  ...allTeachers[0],
  id: "user-teacher-1",
  username: "budi.santoso",
  name: "Budi Santoso",
  displayName: "Budi Santoso",
  firstName: "Budi",
  lastName: "Santoso",
  email: "budi.santoso@plato.edu",
  bio: "Mathematics and Physics educator with 8 years of experience.",
  department: "Mathematics",
  specializations: ["Algebra", "Physics", "Calculus"],
  yearsOfExperience: 8,
  totalCourses: 4,
  totalStudents: 120,
  averageRating: 4.8,
  status: "active",
  isVerified: true,
  isFeatured: true,
  responseTime: "2h",
}

export const teacher: Teacher = allTeachers[0]
