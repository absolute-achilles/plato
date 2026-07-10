"use client"

import { useEffect, useState } from "react"
import { Loader2 } from "lucide-react"

import { useRole } from "@/components/role-provider"
import {
  fetchDashboard,
  getCourseProgress,
} from "@/lib/mocks/services/dashboard.service"
import { student, usersByRole } from "@/lib/mocks/data"
import type {
  AdminDashboard,
  ParentDashboard,
  Role,
  StudentDashboard,
  TeacherDashboard,
} from "@/lib/mocks/types"

import { AchievementList } from "./_components/AchievementList"
import { AnnouncementList } from "./_components/AnnouncementList"
import { CourseCard } from "./_components/CourseCard"
import { ScheduleList } from "./_components/ScheduleList"
import { StatCard } from "./_components/StatCard"

const roleGreetings: Record<Role, string> = {
  teacher: "Ready to teach today?",
  student: "Keep up the great learning!",
  parent: "Here's how your children are doing.",
  admin: "Platform overview at a glance.",
}

export default function DashboardPage() {
  const { role } = useRole()
  const user = usersByRole[role]
  const [state, setState] = useState<{
    role: Role
    data:
      | TeacherDashboard
      | StudentDashboard
      | ParentDashboard
      | AdminDashboard
      | null
    loaded: boolean
  }>(() => ({ role, data: null, loaded: false }))

  useEffect(() => {
    let cancelled = false
    fetchDashboard(role).then((result) => {
      if (!cancelled) {
        setState({ role, data: result, loaded: true })
      }
    })
    return () => {
      cancelled = true
    }
  }, [role])

  const data = state.role === role ? state.data : null
  const loading = state.role !== role || !state.loaded

  if (loading || !data) {
    return (
      <div className="flex h-[60vh] flex-col items-center justify-center gap-3">
        <Loader2 className="h-10 w-10 animate-spin text-primary" />
        <p className="text-sm font-medium text-muted-foreground">
          Loading {role} dashboard...
        </p>
      </div>
    )
  }

  return (
    <div className="mx-auto max-w-7xl space-y-8">
      {/* Header */}
      <section className="clay-card flex flex-col justify-between gap-4 bg-gradient-to-br from-primary to-teal-600 p-6 text-primary-foreground sm:flex-row sm:items-center lg:p-8">
        <div>
          <p className="text-sm font-semibold opacity-90">
            {new Date().toLocaleDateString(undefined, {
              weekday: "long",
              year: "numeric",
              month: "long",
              day: "numeric",
            })}
          </p>
          <h1 className="mt-1 font-heading text-2xl font-bold sm:text-3xl">
            Hello, {user.name.split(" ")[0]}!
          </h1>
          <p className="mt-1 text-sm opacity-90">{roleGreetings[role]}</p>
        </div>
        <div className="flex h-14 w-14 items-center justify-center rounded-2xl bg-white/20 text-2xl font-bold backdrop-blur-sm">
          {user.name.charAt(0)}
        </div>
      </section>

      {/* Stats */}
      <section>
        <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
          {data.stats.map((stat) => (
            <StatCard key={stat.label} stat={stat} />
          ))}
        </div>
      </section>

      {/* Role-specific content */}
      {role === "teacher" && (
        <TeacherContent dashboard={data as TeacherDashboard} />
      )}
      {role === "student" && (
        <StudentContent dashboard={data as StudentDashboard} />
      )}
      {role === "parent" && (
        <ParentContent dashboard={data as ParentDashboard} />
      )}
      {role === "admin" && <AdminContent dashboard={data as AdminDashboard} />}
    </div>
  )
}

function TeacherContent({ dashboard }: { dashboard: TeacherDashboard }) {
  return (
    <div className="grid gap-8 lg:grid-cols-3">
      <div className="space-y-6 lg:col-span-2">
        <div>
          <h2 className="font-heading text-xl font-bold text-foreground">
            My Courses
          </h2>
          <div className="mt-4 grid gap-4 sm:grid-cols-2">
            {dashboard.courses.map((course) => (
              <CourseCard key={course.id} course={course} studentCount={12} />
            ))}
          </div>
        </div>
      </div>

      <div className="space-y-8">
        <div>
          <h2 className="font-heading text-xl font-bold text-foreground">
            Today&apos;s Schedule
          </h2>
          <div className="mt-4">
            <ScheduleList events={dashboard.schedule} />
          </div>
        </div>
        <div>
          <h2 className="font-heading text-xl font-bold text-foreground">
            Announcements
          </h2>
          <div className="mt-4">
            <AnnouncementList announcements={dashboard.announcements} />
          </div>
        </div>
      </div>
    </div>
  )
}

function StudentContent({ dashboard }: { dashboard: StudentDashboard }) {
  return (
    <div className="grid gap-8 lg:grid-cols-3">
      <div className="space-y-6 lg:col-span-2">
        <div>
          <h2 className="font-heading text-xl font-bold text-foreground">
            My Courses
          </h2>
          <div className="mt-4 grid gap-4 sm:grid-cols-2">
            {dashboard.courses.map((course) => (
              <CourseCard
                key={course.id}
                course={course}
                progress={getCourseProgress(course, student.id)}
              />
            ))}
          </div>
        </div>
      </div>

      <div className="space-y-8">
        <div>
          <h2 className="font-heading text-xl font-bold text-foreground">
            Today&apos;s Schedule
          </h2>
          <div className="mt-4">
            <ScheduleList events={dashboard.schedule} />
          </div>
        </div>
        <div>
          <h2 className="font-heading text-xl font-bold text-foreground">
            Achievements
          </h2>
          <div className="mt-4">
            <AchievementList achievements={dashboard.achievements} />
          </div>
        </div>
        <div>
          <h2 className="font-heading text-xl font-bold text-foreground">
            Announcements
          </h2>
          <div className="mt-4">
            <AnnouncementList announcements={dashboard.announcements} />
          </div>
        </div>
      </div>
    </div>
  )
}

function ParentContent({ dashboard }: { dashboard: ParentDashboard }) {
  return (
    <div className="grid gap-8 lg:grid-cols-3">
      <div className="space-y-6 lg:col-span-2">
        <div>
          <h2 className="font-heading text-xl font-bold text-foreground">
            Children&apos;s Courses
          </h2>
          <div className="mt-4 grid gap-4 sm:grid-cols-2">
            {dashboard.courses.map((course) => (
              <CourseCard
                key={course.id}
                course={course}
                progress={
                  dashboard.childrenEnrollments.find(
                    (e) => e.courseId === course.id
                  )?.progressPercent
                }
              />
            ))}
          </div>
        </div>
      </div>

      <div className="space-y-8">
        <div>
          <h2 className="font-heading text-xl font-bold text-foreground">
            Upcoming Schedule
          </h2>
          <div className="mt-4">
            <ScheduleList events={dashboard.schedule} />
          </div>
        </div>
        <div>
          <h2 className="font-heading text-xl font-bold text-foreground">
            Announcements
          </h2>
          <div className="mt-4">
            <AnnouncementList announcements={dashboard.announcements} />
          </div>
        </div>
      </div>
    </div>
  )
}

function AdminContent({ dashboard }: { dashboard: AdminDashboard }) {
  return (
    <div className="grid gap-8 lg:grid-cols-3">
      <div className="space-y-6 lg:col-span-2">
        <div>
          <h2 className="font-heading text-xl font-bold text-foreground">
            Recent Courses
          </h2>
          <div className="mt-4 grid gap-4 sm:grid-cols-2">
            {dashboard.recentCourses.map((course) => (
              <CourseCard key={course.id} course={course} />
            ))}
          </div>
        </div>
      </div>

      <div className="space-y-8">
        <div>
          <h2 className="font-heading text-xl font-bold text-foreground">
            Recent Users
          </h2>
          <div className="mt-4 space-y-3">
            {dashboard.recentUsers.map((user) => (
              <div
                key={user.id}
                className="shadow-clay-sm flex items-center gap-3 rounded-2xl border border-border bg-card p-3"
              >
                <div className="flex h-10 w-10 items-center justify-center rounded-full bg-primary text-sm font-bold text-primary-foreground">
                  {user.name.charAt(0)}
                </div>
                <div className="min-w-0 flex-1">
                  <p className="truncate text-sm font-bold text-foreground">
                    {user.name}
                  </p>
                  <p className="truncate text-xs text-muted-foreground capitalize">
                    {user.role}
                  </p>
                </div>
              </div>
            ))}
          </div>
        </div>
        <div>
          <h2 className="font-heading text-xl font-bold text-foreground">
            Announcements
          </h2>
          <div className="mt-4">
            <AnnouncementList announcements={dashboard.announcements} />
          </div>
        </div>
      </div>
    </div>
  )
}
