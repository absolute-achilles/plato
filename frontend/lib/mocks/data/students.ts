import { generateStudents } from "./generators"
import type { Student } from "../types"

export const allStudents: Student[] = generateStudents(100)

// Override the first student to match the original demo user
allStudents[0] = {
  ...allStudents[0],
  id: "user-student-1",
  username: "anisa.wulandari",
  name: "Anisa Wulandari",
  displayName: "Anisa Wulandari",
  firstName: "Anisa",
  lastName: "Wulandari",
  email: "anisa.w@plato.edu",
  parentId: "user-parent-1",
  bio: "Curious student passionate about science and mathematics.",
  grade: "10",
  level: "High School",
  status: "active",
  gpa: 3.8,
  attendanceRate: 96,
  averageScore: 88,
  streakDays: 12,
  totalPoints: 2450,
  rank: "Silver",
}

export const student: Student = allStudents[0]
