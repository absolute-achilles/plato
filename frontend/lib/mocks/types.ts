export type Role = "student" | "teacher" | "parent" | "admin"

export type UserStatus = "active" | "inactive" | "pending" | "suspended"

export interface User {
  id: string
  username: string
  name: string
  displayName?: string
  firstName?: string
  lastName?: string
  email: string
  role: Role
  avatarUrl?: string
  phone?: string
  location?: string
  bio?: string
  timezone?: string
  languages?: string[]
  linkedInUrl?: string
  websiteUrl?: string
  portfolioUrl?: string
  status?: UserStatus
  lastActiveAt?: Date
  joinDate?: Date
  badges?: string[]
  createdAt: Date
}

export interface Student extends User {
  role: "student"
  parentId?: string
  grade?: string
  level?: string
  enrolledCourses?: string[]
  completedCourses?: number
  inProgressCourses?: number
  gpa?: number
  attendanceRate?: number
  averageScore?: number
  certificatesCount?: number
  interests?: string[]
  achievements?: string[]
  streakDays?: number
  totalPoints?: number
  rank?: string
  mentorId?: string
  preferredLearningStyle?: string
  notes?: string
}

export interface Teacher extends User {
  role: "teacher"
  department?: string
  specializations?: string[]
  yearsOfExperience?: number
  totalCourses?: number
  totalStudents?: number
  averageRating?: number
  twitterUrl?: string
  awards?: string[]
  availability?: string
  isVerified?: boolean
  isFeatured?: boolean
  responseTime?: string
}

export interface Parent extends User {
  role: "parent"
  childrenIds: string[]
}

export interface Admin extends User {
  role: "admin"
}

export type CourseStatus = "draft" | "published" | "archived"
export type CourseLevel = "beginner" | "intermediate" | "advanced"

export interface Course {
  id: string
  teacherId: string
  teacherName?: string
  teacherAvatar?: string
  name: string
  description: string
  shortDescription?: string
  category: string
  level?: CourseLevel
  language?: string[]
  subtitles?: string[]
  prerequisites?: string[]
  learningObjectives?: string[]
  tags?: string[]
  syllabus?: string[]
  color: string
  thumbnailUrl?: string
  coverImageUrl?: string
  duration?: number
  studentCount?: number
  moduleCount?: number
  lessonCount?: number
  rating?: number
  reviewCount?: number
  price?: number
  currency?: string
  status?: CourseStatus
  progress?: number
  startDate?: Date
  endDate?: Date
  enrollmentDeadline?: Date
  certificateOffered?: boolean
  certificateTemplate?: string
  forumEnabled?: boolean
  liveSessions?: string[]
  faq?: string[]
  refundPolicy?: string
  estimatedHoursPerWeek?: number
  lastUpdated?: Date
  createdAt: Date
  updatedAt: Date
}

export type ModuleContentType = "lesson" | "assignment"
export type ModuleDifficulty = "beginner" | "intermediate" | "advanced"
export type ModuleStatus = "draft" | "published" | "archived"

export interface Module {
  id: string
  courseId: string
  courseName?: string
  name: string
  description?: string
  shortDescription?: string
  learningObjectives?: string[]
  prerequisites?: string[]
  difficulty?: ModuleDifficulty
  category?: string
  tags?: string[]
  language?: string
  position: number
  isPublished: boolean
  unlockDate?: Date
  thumbnailUrl?: string
  videoUrl?: string
  duration?: number
  contentCount?: number
  attachmentCount?: number
  quizCount?: number
  assignmentCount?: number
  forumPostCount?: number
  totalPoints?: number
  passingScore?: number
  certificateEligible?: boolean
  isMandatory?: boolean
  estimatedHours?: number
  resources?: string[]
  sections?: string[]
  enrollmentCount?: number
  averageRating?: number
  completionRate?: number
  status?: ModuleStatus
  createdBy?: string
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
