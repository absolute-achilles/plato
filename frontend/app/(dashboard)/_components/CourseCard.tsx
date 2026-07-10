import { BookOpen, UsersRound } from "lucide-react"

import type { Course } from "@/lib/mocks/types"

interface CourseCardProps {
  course: Course
  progress?: number
  studentCount?: number
}

export function CourseCard({
  course,
  progress,
  studentCount,
}: CourseCardProps) {
  return (
    <div className="clay-card flex flex-col gap-4 p-5">
      <div className="flex items-start justify-between gap-3">
        <div
          className="shadow-clay-sm flex h-12 w-12 shrink-0 items-center justify-center rounded-2xl text-white"
          style={{ backgroundColor: course.color }}
        >
          <BookOpen className="h-6 w-6" />
        </div>
        <span className="rounded-full bg-muted px-2.5 py-1 text-xs font-bold text-muted-foreground">
          {course.category}
        </span>
      </div>

      <div>
        <h3 className="font-heading text-lg font-bold text-foreground">
          {course.name}
        </h3>
        <p className="mt-1 line-clamp-2 text-sm text-muted-foreground">
          {course.description}
        </p>
      </div>

      {typeof progress === "number" && (
        <div className="space-y-1.5">
          <div className="flex items-center justify-between text-xs font-semibold">
            <span className="text-muted-foreground">Progress</span>
            <span className="text-foreground">{progress}%</span>
          </div>
          <div className="h-2.5 w-full overflow-hidden rounded-full bg-muted">
            <div
              className="h-full rounded-full bg-primary transition-all duration-500"
              style={{ width: `${progress}%` }}
            />
          </div>
        </div>
      )}

      {typeof studentCount === "number" && (
        <div className="flex items-center gap-2 text-sm font-semibold text-muted-foreground">
          <UsersRound className="h-4 w-4" />
          <span>{studentCount} students</span>
        </div>
      )}
    </div>
  )
}
