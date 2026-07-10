import type { Admin, Parent, User } from "../types"
import { allStudents, student } from "./students"
import { allTeachers, teacher } from "./teachers"

export const parent: Parent = {
  id: "user-parent-1",
  username: "siti.aminah",
  name: "Siti Aminah",
  email: "siti.aminah@email.com",
  role: "parent",
  childrenIds: [student.id],
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

export const allUsers: User[] = [
  teacher,
  student,
  parent,
  admin,
  ...allTeachers.slice(1),
  ...allStudents.slice(1),
]

export const usersByRole: Record<string, User> = {
  teacher,
  student,
  parent,
  admin,
}

export { teacher, student, allTeachers, allStudents }
