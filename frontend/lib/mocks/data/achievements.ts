import type { Achievement } from "../types"
import { student } from "./users"

export const achievements: Achievement[] = [
  {
    id: "ach-1",
    studentId: student.id,
    title: "First Steps",
    description: "Completed your first lesson",
    icon: "Footprints",
    earnedAt: new Date("2024-01-20T08:00:00Z"),
  },
  {
    id: "ach-2",
    studentId: student.id,
    title: "Perfect Attendance",
    description: "Attended 5 classes in a row",
    icon: "CalendarCheck",
    earnedAt: new Date("2024-02-15T08:00:00Z"),
  },
  {
    id: "ach-3",
    studentId: student.id,
    title: "Quiz Master",
    description: "Scored 100% on a quiz",
    icon: "Trophy",
    earnedAt: new Date("2024-03-10T08:00:00Z"),
  },
  {
    id: "ach-4",
    studentId: student.id,
    title: "Course Explorer",
    description: "Enrolled in 4 courses",
    icon: "Compass",
    earnedAt: new Date("2024-03-20T08:00:00Z"),
  },
]

export function getAchievementsByStudent(studentId: string): Achievement[] {
  return achievements
    .filter((a) => a.studentId === studentId)
    .sort((a, b) => b.earnedAt.getTime() - a.earnedAt.getTime())
}
