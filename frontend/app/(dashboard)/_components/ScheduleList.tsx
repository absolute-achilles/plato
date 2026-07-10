import { CalendarCheck, Clock, FileText, FlaskConical } from "lucide-react"

import type { ScheduleEvent } from "@/lib/mocks/types"
import { getCourseById } from "@/lib/mocks/data"
import { cn } from "@/lib/utils"

interface ScheduleListProps {
  events: ScheduleEvent[]
}

const typeConfig = {
  "live-class": {
    icon: Clock,
    label: "Live Class",
    color: "bg-blue-100 text-blue-700",
  },
  "assignment-due": {
    icon: FileText,
    label: "Due",
    color: "bg-amber-100 text-amber-700",
  },
  exam: {
    icon: FlaskConical,
    label: "Exam",
    color: "bg-rose-100 text-rose-700",
  },
}

export function ScheduleList({ events }: ScheduleListProps) {
  if (events.length === 0) {
    return (
      <div className="flex flex-col items-center justify-center rounded-2xl border border-dashed border-border py-10 text-center">
        <CalendarCheck className="h-10 w-10 text-muted-foreground" />
        <p className="mt-3 text-sm font-medium text-muted-foreground">
          No upcoming events
        </p>
      </div>
    )
  }

  return (
    <ul className="space-y-3">
      {events.map((event) => {
        const course = getCourseById(event.courseId)
        const config = typeConfig[event.type]
        const Icon = config.icon

        return (
          <li
            key={event.id}
            className="flex items-center gap-4 rounded-2xl border border-border bg-card p-4 shadow-clay-sm transition-all hover:shadow-clay"
          >
            <div
              className={cn(
                "flex h-11 w-11 shrink-0 items-center justify-center rounded-xl",
                config.color
              )}
            >
              <Icon className="h-5 w-5" />
            </div>
            <div className="min-w-0 flex-1">
              <p className="truncate font-heading text-sm font-bold text-foreground">
                {event.title}
              </p>
              <p className="truncate text-xs text-muted-foreground">
                {course?.name ?? "General"}
              </p>
            </div>
            <div className="text-right">
              <p className="text-xs font-bold text-foreground">
                {event.startAt.toLocaleTimeString([], {
                  hour: "2-digit",
                  minute: "2-digit",
                })}
              </p>
              <p className="text-xs text-muted-foreground">
                {event.startAt.toLocaleDateString([], {
                  weekday: "short",
                  day: "numeric",
                })}
              </p>
            </div>
          </li>
        )
      })}
    </ul>
  )
}
