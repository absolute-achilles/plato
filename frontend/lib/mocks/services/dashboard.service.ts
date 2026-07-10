import type {
  AdminDashboard,
  Course,
  DashboardStats,
  ParentDashboard,
  Role,
  ScheduleEvent,
  StudentDashboard,
  TeacherDashboard,
} from "../types"
import {
  allUsers,
  courses,
  enrollments,
  getAchievementsByStudent,
  getAttendanceRate,
  getCourseById,
  getCoursesByIds,
  getEnrollmentsByStudent,
  getRecentAnnouncements,
  getScheduleForCourseIds,
  scheduleEvents,
  student,
  teacher,
} from "../data"

export async function fetchDashboard<T extends Role>(
  role: T
): Promise<DashboardForRole<T>> {
  await delay(400)

  switch (role) {
    case "teacher":
      return buildTeacherDashboard() as DashboardForRole<T>
    case "student":
      return buildStudentDashboard() as DashboardForRole<T>
    case "parent":
      return buildParentDashboard() as DashboardForRole<T>
    case "admin":
      return buildAdminDashboard() as DashboardForRole<T>
    default:
      throw new Error(`Unknown role: ${role}`)
  }
}

export type DashboardForRole<T extends Role> = T extends "teacher"
  ? TeacherDashboard
  : T extends "student"
    ? StudentDashboard
    : T extends "parent"
      ? ParentDashboard
      : T extends "admin"
        ? AdminDashboard
        : never

function buildTeacherDashboard(): TeacherDashboard {
  const teacherCourses = courses.filter((c) => c.teacherId === teacher.id)
  const courseIds = teacherCourses.map((c) => c.id)
  const enrolledStudents = new Set(
    enrollments
      .filter((e) => courseIds.includes(e.courseId))
      .map((e) => e.studentId)
  ).size

  const stats: DashboardStats[] = [
    { label: "My Courses", value: String(teacherCourses.length) },
    { label: "Students", value: String(enrolledStudents) },
    {
      label: "Avg. Attendance",
      value: "92%",
      change: "+3%",
      trend: "up",
    },
    { label: "Pending Reviews", value: "5" },
  ]

  return {
    stats,
    courses: teacherCourses,
    schedule: getScheduleForCourseIds(courseIds).slice(0, 5),
    announcements: getRecentAnnouncements(3),
  }
}

function buildStudentDashboard(): StudentDashboard {
  const studentEnrollments = getEnrollmentsByStudent(student.id)
  const courseIds = studentEnrollments.map((e) => e.courseId)
  const enrolledCourses = getCoursesByIds(courseIds)
  const avgProgress = Math.round(
    studentEnrollments.reduce((sum, e) => sum + e.progressPercent, 0) /
      Math.max(studentEnrollments.length, 1)
  )
  const attendanceRate = getAttendanceRate(student.id)
  const upcomingDeadlines = scheduleEvents.filter(
    (e) => courseIds.includes(e.courseId) && e.type === "assignment-due"
  ).length

  const stats: DashboardStats[] = [
    { label: "My Courses", value: String(studentEnrollments.length) },
    { label: "Avg. Progress", value: `${avgProgress}%` },
    {
      label: "Attendance",
      value: `${attendanceRate}%`,
      change: "-2%",
      trend: "down",
    },
    { label: "Due Soon", value: String(upcomingDeadlines) },
  ]

  return {
    stats,
    enrollments: studentEnrollments,
    courses: enrolledCourses,
    schedule: getScheduleForCourseIds(courseIds).slice(0, 5),
    announcements: getRecentAnnouncements(4),
    achievements: getAchievementsByStudent(student.id).slice(0, 4),
  }
}

function buildParentDashboard(): ParentDashboard {
  const childIds = [student.id]
  const childEnrollments = enrollments.filter((e) => childIds.includes(e.studentId))
  const courseIds = childEnrollments.map((e) => e.courseId)
  const enrolledCourses = getCoursesByIds(courseIds)
  const avgProgress = Math.round(
    childEnrollments.reduce((sum, e) => sum + e.progressPercent, 0) /
      Math.max(childEnrollments.length, 1)
  )
  const attendanceRate = getAttendanceRate(student.id)

  const stats: DashboardStats[] = [
    { label: "Children", value: String(childIds.length) },
    { label: "Enrolled Courses", value: String(childEnrollments.length) },
    {
      label: "Avg. Progress",
      value: `${avgProgress}%`,
      change: "+5%",
      trend: "up",
    },
    {
      label: "Attendance",
      value: `${attendanceRate}%`,
      change: "-2%",
      trend: "down",
    },
  ]

  return {
    stats,
    children: [student],
    childrenEnrollments: childEnrollments,
    courses: enrolledCourses,
    schedule: getScheduleForCourseIds(courseIds).slice(0, 5),
    announcements: getRecentAnnouncements(3),
  }
}

function buildAdminDashboard(): AdminDashboard {
  const stats: DashboardStats[] = [
    { label: "Total Users", value: String(allUsers.length) },
    { label: "Courses", value: String(courses.length) },
    { label: "Enrollments", value: String(enrollments.length) },
    {
      label: "Active Today",
      value: "3",
      change: "+1",
      trend: "up",
    },
  ]

  return {
    stats,
    recentUsers: allUsers.slice(0, 4),
    recentCourses: courses.slice(0, 3),
    announcements: getRecentAnnouncements(3),
  }
}

export function getCourseProgress(
  course: Course,
  studentId: string
): number {
  const enrollment = enrollments.find(
    (e) => e.courseId === course.id && e.studentId === studentId
  )
  return enrollment?.progressPercent ?? 0
}

export function getEventCourse(event: ScheduleEvent): Course | undefined {
  return getCourseById(event.courseId)
}

function delay(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms))
}
