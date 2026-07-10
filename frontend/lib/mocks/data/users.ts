import type { Admin, Parent, Student, Teacher, User } from "../types"

export const teacher: Teacher = {
  id: "user-teacher-1",
  username: "budi.santoso",
  name: "Budi Santoso",
  email: "budi.santoso@plato.edu",
  role: "teacher",
  bio: "Mathematics and Physics educator with 8 years of experience.",
  avatarUrl: "https://api.dicebear.com/7.x/avataaars/svg?seed=Budi",
  createdAt: new Date("2023-06-01T08:00:00Z"),
}

export const student: Student = {
  id: "user-student-1",
  username: "anisa.wulandari",
  name: "Anisa Wulandari",
  email: "anisa.w@plato.edu",
  role: "student",
  parentId: "user-parent-1",
  avatarUrl: "https://api.dicebear.com/7.x/avataaars/svg?seed=Anisa",
  createdAt: new Date("2023-06-15T08:00:00Z"),
}

export const parent: Parent = {
  id: "user-parent-1",
  username: "siti.aminah",
  name: "Siti Aminah",
  email: "siti.aminah@email.com",
  role: "parent",
  childrenIds: ["user-student-1"],
  avatarUrl: "https://api.dicebear.com/7.x/avataaars/svg?seed=Siti",
  createdAt: new Date("2023-06-10T08:00:00Z"),
}

export const admin: Admin = {
  id: "user-admin-1",
  username: "admin",
  name: "Plato Administrator",
  email: "admin@plato.edu",
  role: "admin",
  avatarUrl: "https://api.dicebear.com/7.x/avataaars/svg?seed=Admin",
  createdAt: new Date("2023-01-01T08:00:00Z"),
}

export const allUsers: User[] = [teacher, student, parent, admin]

export const usersByRole: Record<string, User> = {
  teacher,
  student,
  parent,
  admin,
}
