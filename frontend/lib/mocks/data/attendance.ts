import type { Attendance } from "../types"
import { student } from "./users"

export const attendanceRecords: Attendance[] = [
  {
    id: "att-1",
    studentId: student.id,
    moduleId: "mod-math-1",
    status: "present",
    recordedAt: new Date("2024-07-07T08:00:00Z"),
    notes: "Participated actively",
  },
  {
    id: "att-2",
    studentId: student.id,
    moduleId: "mod-math-2",
    status: "present",
    recordedAt: new Date("2024-07-08T08:00:00Z"),
  },
  {
    id: "att-3",
    studentId: student.id,
    moduleId: "mod-phys-1",
    status: "late",
    recordedAt: new Date("2024-07-08T08:15:00Z"),
    notes: "Arrived 15 minutes late",
  },
  {
    id: "att-4",
    studentId: student.id,
    moduleId: "mod-bi-1",
    status: "present",
    recordedAt: new Date("2024-07-09T08:00:00Z"),
  },
  {
    id: "att-5",
    studentId: student.id,
    moduleId: "mod-eng-1",
    status: "absent",
    recordedAt: new Date("2024-07-09T08:00:00Z"),
    notes: "Sick leave",
  },
  {
    id: "att-6",
    studentId: student.id,
    moduleId: "mod-math-3",
    status: "present",
    recordedAt: new Date("2024-07-10T08:00:00Z"),
  },
]

export function getAttendanceByStudent(studentId: string): Attendance[] {
  return attendanceRecords.filter((a) => a.studentId === studentId)
}

export function getAttendanceRate(studentId: string): number {
  const records = getAttendanceByStudent(studentId)
  if (records.length === 0) return 0
  const present = records.filter((r) => r.status === "present").length
  return Math.round((present / records.length) * 100)
}
