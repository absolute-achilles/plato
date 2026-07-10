import type {
  Course,
  CourseLevel,
  CourseStatus,
  Module,
  ModuleDifficulty,
  ModuleStatus,
  Student,
  Teacher,
  UserStatus,
} from "../types"

export function createAvatarUrl(seed: string): string {
  return `https://api.dicebear.com/7.x/avataaars/svg?seed=${encodeURIComponent(seed)}`
}

export function createThumbnailUrl(text: string, color: string): string {
  const cleanText = encodeURIComponent(text.replace(/\s+/g, "+"))
  const cleanColor = color.replace("#", "")
  return `https://placehold.co/600x400/${cleanColor}/FFFFFF?text=${cleanText}`
}

export function createCoverImageUrl(seed: string, color: string): string {
  const cleanColor = color.replace("#", "")
  return `https://placehold.co/1200x400/${cleanColor}/FFFFFF?text=${encodeURIComponent(seed)}`
}

function hashString(str: string): number {
  let hash = 0
  for (let i = 0; i < str.length; i++) {
    const char = str.charCodeAt(i)
    hash = (hash << 5) - hash + char
    hash = hash & hash
  }
  return Math.abs(hash)
}

function createRng(seed: string): () => number {
  let state = hashString(seed) || 1
  return () => {
    state = (state * 1664525 + 1013904223) % 2 ** 32
    return state / 2 ** 32
  }
}

const courseNames = [
  "Mathematics 101",
  "Physics Fundamentals",
  "Bahasa Indonesia",
  "English Grammar",
  "Chemistry Basics",
  "Biology Introduction",
  "World History",
  "Geography Essentials",
  "Computer Science 101",
  "Introduction to Programming",
  "Web Development",
  "Mobile App Development",
  "Data Science Basics",
  "Artificial Intelligence",
  "Machine Learning 101",
  "Digital Literacy",
  "Entrepreneurship",
  "Financial Literacy",
  "Public Speaking",
  "Creative Writing",
  "Music Theory",
  "Visual Arts",
  "Physical Education",
  "Health and Wellness",
]

const courseCategories = [
  "Mathematics",
  "Science",
  "Language",
  "Technology",
  "Arts",
  "Humanities",
  "Business",
  "Health",
]

const courseColors = [
  "#0D9488",
  "#D97706",
  "#2563EB",
  "#7C3AED",
  "#DC2626",
  "#059669",
  "#0891B2",
  "#BE123C",
  "#4338CA",
  "#A16207",
]

const courseLevels: CourseLevel[] = ["beginner", "intermediate", "advanced"]
const courseStatuses: CourseStatus[] = ["draft", "published", "archived"]

const moduleNames = [
  "Introduction",
  "Core Concepts",
  "Practical Applications",
  "Advanced Topics",
  "Case Studies",
  "Assessment",
  "Review",
  "Final Project",
  "Lab Session",
  "Group Discussion",
  "Workshop",
  "Quiz",
  "Deep Dive",
  "Supplementary Reading",
  "Practice Problems",
  "Real-world Examples",
  "Mastery Check",
  "Capstone",
]

const moduleDifficulties: ModuleDifficulty[] = [
  "beginner",
  "intermediate",
  "advanced",
]
const moduleStatuses: ModuleStatus[] = ["draft", "published", "archived"]

const userStatuses: UserStatus[] = ["active", "inactive", "pending", "suspended"]

const departments = [
  "Mathematics",
  "Science",
  "Language",
  "Technology",
  "Arts",
  "Humanities",
  "Business",
  "Health",
]

const specializations = [
  "Algebra",
  "Calculus",
  "Physics",
  "Chemistry",
  "Biology",
  "Literature",
  "Grammar",
  "Programming",
  "Web Development",
  "Data Science",
  "Machine Learning",
  "History",
  "Geography",
  "Music",
  "Visual Arts",
]

const grades = ["7", "8", "9", "10", "11", "12"]
const levels = ["Middle School", "High School", "Undergraduate"]
const learningStyles = ["Visual", "Auditory", "Kinesthetic", "Reading/Writing"]
const timezones = ["WIB", "WITA", "WIT", "UTC+7", "UTC+8"]
const languagesList = ["Indonesian", "English", "Mandarin", "Japanese", "Korean"]
const locations = ["Jakarta", "Bandung", "Surabaya", "Yogyakarta", "Medan", "Makassar", "Denpasar", "Malang"]
const interests = ["Mathematics", "Science", "Coding", "Art", "Music", "Sports", "Reading", "Writing", "History", "Geography"]
const badges = ["Early Bird", "Top Performer", "Helper", "Consistent", "Innovator", "Leader", "Scholar"]

const firstNames = [
  "Ahmad", "Budi", "Citra", "Dewi", "Eka", "Fajar", "Gita", "Hadi", "Indra", "Joko",
  "Kirana", "Lestari", "Maya", "Nadia", "Olivia", "Putri", "Qori", "Raka", "Sari", "Tina",
  "Umar", "Vina", "Wahyu", "Xena", "Yani", "Zaki", "Adi", "Bayu", "Cici", "Dian",
  "Elang", "Fitri", "Gede", "Hana", "Ivan", "Jeni", "Krisna", "Lina", "Maman", "Nina",
  "Oka", "Puspita", "Qais", "Rani", "Sandi", "Tari", "Udin", "Vera", "Wulan", "Yoga",
]

const lastNames = [
  "Santoso", "Wulandari", "Setiawan", "Kusuma", "Pratama", "Sari", "Hidayat", "Saputra", "Utami", "Lestari",
  "Handoko", "Siregar", "Mulyani", "Wibowo", "Nugroho", "Purnama", "Kurniawan", "Susanto", "Rahayu", "Ananda",
  "Putra", "Dewi", "Agustina", "Sihombing", "Simanjuntak", "Nasution", "Harahap", "Sinaga", "Tambunan", "Nainggolan",
  "Situmorang", "Marpaung", "Lumbantobing", "Panggabean", "Samosir", "Aritonang", "Hutapea", "Siagian", "Pane", "Sitorus",
]

function pickWithRng<T>(arr: T[], rng: () => number): T {
  return arr[Math.floor(rng() * arr.length)]
}

function pickManyWithRng<T>(arr: T[], count: number, rng: () => number): T[] {
  const result: T[] = []
  const available = [...arr]
  while (result.length < count && available.length > 0) {
    const index = Math.floor(rng() * available.length)
    result.push(available.splice(index, 1)[0])
  }
  return result
}

function randomDateWithRng(start: Date, end: Date, rng: () => number): Date {
  return new Date(start.getTime() + rng() * (end.getTime() - start.getTime()))
}

export function generateName(rng: () => number): { firstName: string; lastName: string; fullName: string } {
  const firstName = pickWithRng(firstNames, rng)
  const lastName = pickWithRng(lastNames, rng)
  return { firstName, lastName, fullName: `${firstName} ${lastName}` }
}

export function generateTeacher(id: string, index: number): Teacher {
  const rng = createRng(`${id}-${index}`)
  const { firstName, lastName, fullName } = generateName(rng)
  const username = `${firstName.toLowerCase()}.${lastName.toLowerCase()}-${index}`
  const chosenSpecializations = pickManyWithRng(specializations, 1 + Math.floor(rng() * 3), rng)
  const department = pickWithRng(departments, rng)
  const status = pickWithRng(userStatuses, rng)
  const createdAt = randomDateWithRng(new Date("2022-01-01"), new Date("2024-06-01"), rng)
  const lastActiveAt = randomDateWithRng(createdAt, new Date(), rng)
  const totalCourses = Math.floor(rng() * 8) + 1
  const totalStudents = Math.floor(rng() * 200) + 20
  const averageRating = Math.round((3 + rng() * 2) * 10) / 10

  return {
    id,
    username,
    name: fullName,
    displayName: fullName,
    firstName,
    lastName,
    email: `${username}@plato.edu`,
    role: "teacher",
    avatarUrl: createAvatarUrl(username),
    phone: `+62${Math.floor(8000000000 + rng() * 999999999)}`,
    location: pickWithRng(locations, rng),
    bio: `${department} educator with a passion for ${chosenSpecializations.join(", ")}.`,
    department,
    specializations: chosenSpecializations,
    yearsOfExperience: Math.floor(rng() * 15) + 1,
    totalCourses,
    totalStudents,
    averageRating,
    status,
    lastActiveAt,
    joinDate: createdAt,
    linkedInUrl: `https://linkedin.com/in/${username}`,
    twitterUrl: `https://twitter.com/${username}`,
    websiteUrl: `https://${username}.plato.edu`,
    awards: pickManyWithRng(["Best Educator 2023", "Innovation Award", "Student Choice", "Top Mentor"], Math.floor(rng() * 3), rng),
    languages: pickManyWithRng(languagesList, Math.floor(rng() * 2) + 1, rng),
    timezone: pickWithRng(timezones, rng),
    availability: pickWithRng(["Weekdays", "Weekends", "Evenings", "Flexible"], rng),
    isVerified: rng() > 0.3,
    isFeatured: rng() > 0.8,
    responseTime: `${Math.floor(rng() * 24) + 1}h`,
    createdAt,
  }
}

export function generateStudent(id: string, index: number): Student {
  const rng = createRng(`${id}-${index}`)
  const { firstName, lastName, fullName } = generateName(rng)
  const username = `${firstName.toLowerCase()}.${lastName.toLowerCase()}-${index}`
  const status = pickWithRng(userStatuses, rng)
  const createdAt = randomDateWithRng(new Date("2022-01-01"), new Date("2024-06-01"), rng)
  const lastActiveAt = randomDateWithRng(createdAt, new Date(), rng)
  const enrolledCount = Math.floor(rng() * 5) + 1
  const enrolledCourses: string[] = []
  for (let i = 0; i < enrolledCount; i++) {
    enrolledCourses.push(`course-${Math.floor(rng() * 24) + 1}`)
  }
  const uniqueEnrolledCourses = [...new Set(enrolledCourses)]

  return {
    id,
    username,
    name: fullName,
    displayName: fullName,
    firstName,
    lastName,
    email: `${username}@student.plato.edu`,
    role: "student",
    avatarUrl: createAvatarUrl(username),
    phone: `+62${Math.floor(8000000000 + rng() * 999999999)}`,
    location: pickWithRng(locations, rng),
    bio: `Passionate learner from ${pickWithRng(locations, rng)}.`,
    grade: pickWithRng(grades, rng),
    level: pickWithRng(levels, rng),
    enrolledCourses: uniqueEnrolledCourses,
    completedCourses: Math.floor(rng() * 3),
    inProgressCourses: uniqueEnrolledCourses.length - Math.floor(rng() * 2),
    gpa: Math.round((2 + rng() * 2) * 10) / 10,
    attendanceRate: Math.round((80 + rng() * 20) * 10) / 10,
    averageScore: Math.round((70 + rng() * 30) * 10) / 10,
    certificatesCount: Math.floor(rng() * 5),
    interests: pickManyWithRng(interests, Math.floor(rng() * 3) + 1, rng),
    achievements: pickManyWithRng(badges, Math.floor(rng() * 3), rng),
    streakDays: Math.floor(rng() * 30),
    totalPoints: Math.floor(rng() * 5000),
    rank: pickWithRng(["Bronze", "Silver", "Gold", "Platinum"], rng),
    preferredLearningStyle: pickWithRng(learningStyles, rng),
    notes: `Student shows strong interest in ${pickWithRng(interests, rng)}.`,
    status,
    lastActiveAt,
    joinDate: createdAt,
    linkedInUrl: `https://linkedin.com/in/${username}`,
    portfolioUrl: `https://${username}.portfolio.dev`,
    languages: pickManyWithRng(languagesList, Math.floor(rng() * 2) + 1, rng),
    timezone: pickWithRng(timezones, rng),
    badges: pickManyWithRng(badges, Math.floor(rng() * 3), rng),
    createdAt,
  }
}

export function generateCourse(id: string, index: number): Course {
  const rng = createRng(`${id}-${index}`)
  const name = courseNames[index % courseNames.length]
  const category = pickWithRng(courseCategories, rng)
  const color = courseColors[index % courseColors.length]
  const level = pickWithRng(courseLevels, rng)
  const status = pickWithRng(courseStatuses, rng)
  const teacherId = `teacher-${(index % 50) + 1}`
  const teacherName = `Teacher ${(index % 50) + 1}`
  const createdAt = randomDateWithRng(new Date("2022-01-01"), new Date("2024-06-01"), rng)
  const updatedAt = randomDateWithRng(createdAt, new Date(), rng)
  const moduleCount = Math.floor(rng() * 8) + 3
  const lessonCount = moduleCount * 2 + Math.floor(rng() * 4)
  const duration = moduleCount * 60 + Math.floor(rng() * 120)

  return {
    id,
    teacherId,
    teacherName,
    teacherAvatar: createAvatarUrl(teacherId),
    name,
    description: `A comprehensive ${level} course on ${name}. Students will explore core concepts, practical applications, and real-world problem solving.`,
    shortDescription: `Master ${name} with hands-on projects and assessments.`,
    category,
    level,
    language: ["Indonesian", "English"],
    subtitles: ["English", "Indonesian"],
    prerequisites: [`Basic ${category} knowledge`, "Willingness to learn"],
    learningObjectives: [
      `Understand core ${name} concepts`,
      "Apply knowledge to real-world problems",
      "Complete hands-on assessments",
    ],
    tags: [category, level, name.split(" ")[0]],
    syllabus: Array.from({ length: moduleCount }, (_, i) => `Module ${i + 1}: ${pickWithRng(moduleNames, rng)}`),
    color,
    thumbnailUrl: createThumbnailUrl(name, color),
    coverImageUrl: createCoverImageUrl(name, color),
    duration,
    studentCount: Math.floor(rng() * 200) + 10,
    moduleCount,
    lessonCount,
    rating: Math.round((3 + rng() * 2) * 10) / 10,
    reviewCount: Math.floor(rng() * 100),
    price: Math.floor(rng() * 500) * 1000,
    currency: "IDR",
    status,
    progress: Math.floor(rng() * 100),
    startDate: randomDateWithRng(new Date("2024-01-01"), new Date("2024-12-01"), rng),
    endDate: randomDateWithRng(new Date("2025-01-01"), new Date("2025-12-01"), rng),
    enrollmentDeadline: randomDateWithRng(new Date("2024-12-01"), new Date("2025-03-01"), rng),
    certificateOffered: rng() > 0.3,
    certificateTemplate: "Standard Certificate",
    forumEnabled: rng() > 0.5,
    liveSessions: ["Weekly Q&A", "Office Hours"],
    faq: ["How long is the course?", "Is there a certificate?", "Can I join late?"],
    refundPolicy: "7-day refund policy",
    estimatedHoursPerWeek: Math.floor(rng() * 6) + 2,
    lastUpdated: updatedAt,
    createdAt,
    updatedAt,
  }
}

export function generateModule(id: string, index: number, courseId: string): Module {
  const rng = createRng(`${id}-${index}-${courseId}`)
  const name = `${pickWithRng(moduleNames, rng)} ${index + 1}`
  const difficulty = pickWithRng(moduleDifficulties, rng)
  const status = pickWithRng(moduleStatuses, rng)
  const courseName = `Course ${courseId}`
  const createdAt = randomDateWithRng(new Date("2022-01-01"), new Date("2024-06-01"), rng)
  const updatedAt = randomDateWithRng(createdAt, new Date(), rng)
  const contentCount = Math.floor(rng() * 6) + 2
  const quizCount = Math.floor(rng() * 2)
  const assignmentCount = Math.floor(rng() * 2) + 1
  const duration = contentCount * 15 + Math.floor(rng() * 30)
  const color = courseColors[parseInt(courseId.split("-")[1] || "1") % courseColors.length]

  return {
    id,
    courseId,
    courseName,
    name,
    description: `This module covers ${name.toLowerCase()} concepts in depth.`,
    shortDescription: `Learn ${name} through lessons and assessments.`,
    learningObjectives: [
      `Understand ${name} fundamentals`,
      `Apply ${name} techniques`,
      `Analyze ${name} scenarios`,
    ],
    prerequisites: ["Previous module completed"],
    difficulty,
    category: pickWithRng(courseCategories, rng),
    tags: [difficulty, name.split(" ")[0]],
    language: "Indonesian",
    position: index + 1,
    isPublished: status === "published",
    unlockDate: randomDateWithRng(new Date("2024-01-01"), new Date("2025-06-01"), rng),
    thumbnailUrl: createThumbnailUrl(name, color),
    videoUrl: "https://example.com/video.mp4",
    duration,
    contentCount,
    attachmentCount: Math.floor(rng() * 3),
    quizCount,
    assignmentCount,
    forumPostCount: Math.floor(rng() * 20),
    totalPoints: (quizCount + assignmentCount) * 100,
    passingScore: 70,
    certificateEligible: rng() > 0.5,
    isMandatory: rng() > 0.3,
    estimatedHours: Math.ceil(duration / 60),
    resources: ["Reading PDF", "Video Lecture", "Practice Worksheet"],
    sections: ["Lesson", "Quiz", "Assignment"],
    enrollmentCount: Math.floor(rng() * 200) + 10,
    averageRating: Math.round((3 + rng() * 2) * 10) / 10,
    completionRate: Math.floor(rng() * 100),
    status,
    createdBy: `teacher-${(index % 50) + 1}`,
    createdAt,
    updatedAt,
  }
}

export function generateTeachers(count: number): Teacher[] {
  return Array.from({ length: count }, (_, i) => generateTeacher(`teacher-${i + 1}`, i + 1))
}

export function generateStudents(count: number): Student[] {
  return Array.from({ length: count }, (_, i) => generateStudent(`student-${i + 1}`, i + 1))
}

export function generateCourses(count: number): Course[] {
  return Array.from({ length: count }, (_, i) => generateCourse(`course-${i + 1}`, i))
}

export function generateModules(count: number, courses: Course[]): Module[] {
  const modules: Module[] = []
  let moduleIndex = 0
  const courseRng = createRng("generateModules-courses")

  courses.forEach((course) => {
    const courseRng2 = createRng(`generateModules-${course.id}`)
    const courseModuleCount = Math.floor(courseRng2() * 4) + 2
    for (let i = 0; i < courseModuleCount; i++) {
      modules.push(generateModule(`module-${moduleIndex + 1}`, i, course.id))
      moduleIndex++
      if (moduleIndex >= count) break
    }
  })

  while (moduleIndex < count) {
    const course = pickWithRng(courses, courseRng)
    modules.push(
      generateModule(
        `module-${moduleIndex + 1}`,
        modules.filter((m) => m.courseId === course.id).length,
        course.id
      )
    )
    moduleIndex++
  }

  return modules
}
