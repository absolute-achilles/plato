export type Role = "student" | "teacher" | "parent" | "admin"

export interface User {
  id: string
  username: string
  name: string
  email: string
  role: Role
  avatarUrl?: string
  createdAt: Date
}

export interface Student extends User {
  role: "student"
  parentId?: string
}

export interface Teacher extends User {
  role: "teacher"
  bio?: string
}

export interface Parent extends User {
  role: "parent"
  childrenIds: string[]
}

export interface Admin extends User {
  role: "admin"
}

export interface Course {
  id: string
  teacherId: string
  name: string
  description: string
  category: string
  color: string
  thumbnailUrl?: string
  createdAt: Date
  updatedAt: Date
}

export type ModuleContentType = "lesson" | "assignment"

export interface Module {
  id: string
  courseId: string
  name: string
  position: number
  isPublished: boolean
  unlockDate?: Date
  createdAt: Date
  updatedAt: Date
}

export interface ModuleContent {
  id: string
  moduleId: string
  title: string
  type: ModuleContentType
  bodyContent: string
  position: number
  isPublished: boolean
  createdAt: Date
  updatedAt: Date
}

export interface ContentAttachment {
  id: string
  moduleContentId: string
  name: string
  url: string
  sizeBytes: number
  type: string
  createdAt: Date
}

export interface Enrollment {
  id: string
  studentId: string
  courseId: string
  enrolledAt: Date
  progressPercent: number
}

export type AttendanceStatus = "present" | "absent" | "late" | "excused"

export interface Attendance {
  id: string
  studentId: string
  moduleId: string
  status: AttendanceStatus
  recordedAt: Date
  notes?: string
}

export interface Announcement {
  id: string
  title: string
  body: string
  authorId: string
  courseId?: string
  createdAt: Date
}

export interface Achievement {
  id: string
  studentId: string
  title: string
  description: string
  icon: string
  earnedAt: Date
}

export interface ScheduleEvent {
  id: string
  courseId: string
  moduleId: string
  title: string
  startAt: Date
  endAt: Date
  type: "live-class" | "assignment-due" | "exam"
}

export interface DashboardStats {
  label: string
  value: string
  change?: string
  trend?: "up" | "down" | "neutral"
}

export interface TeacherDashboard {
  stats: DashboardStats[]
  courses: Course[]
  schedule: ScheduleEvent[]
  announcements: Announcement[]
}

export interface StudentDashboard {
  stats: DashboardStats[]
  enrollments: Enrollment[]
  courses: Course[]
  schedule: ScheduleEvent[]
  announcements: Announcement[]
  achievements: Achievement[]
}

export interface ParentDashboard {
  stats: DashboardStats[]
  children: Student[]
  childrenEnrollments: Enrollment[]
  courses: Course[]
  schedule: ScheduleEvent[]
  announcements: Announcement[]
}

export interface AdminDashboard {
  stats: DashboardStats[]
  recentUsers: User[]
  recentCourses: Course[]
  announcements: Announcement[]
}
