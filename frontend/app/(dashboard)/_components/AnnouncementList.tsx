import { Megaphone } from "lucide-react"

import type { Announcement } from "@/lib/mocks/types"
import { usersByRole } from "@/lib/mocks/data"

interface AnnouncementListProps {
  announcements: Announcement[]
}

export function AnnouncementList({ announcements }: AnnouncementListProps) {
  if (announcements.length === 0) {
    return (
      <div className="rounded-2xl border border-dashed border-border py-8 text-center text-sm text-muted-foreground">
        No announcements yet
      </div>
    )
  }

  return (
    <ul className="space-y-3">
      {announcements.map((announcement) => {
        const author = usersByRole[announcement.authorId] ?? {
          name: "Plato Team",
        }

        return (
          <li
            key={announcement.id}
            className="shadow-clay-sm hover:shadow-clay rounded-2xl border border-border bg-card p-4 transition-all"
          >
            <div className="flex items-start gap-3">
              <div className="flex h-9 w-9 shrink-0 items-center justify-center rounded-xl bg-accent text-accent-foreground">
                <Megaphone className="h-4 w-4" />
              </div>
              <div className="min-w-0 flex-1">
                <p className="font-heading text-sm font-bold text-foreground">
                  {announcement.title}
                </p>
                <p className="mt-1 line-clamp-2 text-xs text-muted-foreground">
                  {announcement.body}
                </p>
                <p className="mt-2 text-xs font-semibold text-muted-foreground">
                  {author.name} ·{" "}
                  {announcement.createdAt.toLocaleDateString([], {
                    month: "short",
                    day: "numeric",
                  })}
                </p>
              </div>
            </div>
          </li>
        )
      })}
    </ul>
  )
}
