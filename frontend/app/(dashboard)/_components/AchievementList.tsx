import { Compass, Footprints, Trophy } from "lucide-react"

import type { Achievement } from "@/lib/mocks/types"

interface AchievementListProps {
  achievements: Achievement[]
}

const iconMap: Record<string, React.ElementType> = {
  Footprints,
  CalendarCheck: Trophy,
  Trophy,
  Compass,
}

export function AchievementList({ achievements }: AchievementListProps) {
  if (achievements.length === 0) {
    return (
      <div className="rounded-2xl border border-dashed border-border py-8 text-center text-sm text-muted-foreground">
        No achievements yet
      </div>
    )
  }

  return (
    <div className="grid grid-cols-2 gap-3">
      {achievements.map((achievement) => {
        const Icon = iconMap[achievement.icon] ?? Trophy

        return (
          <div
            key={achievement.id}
            className="shadow-clay-sm hover:shadow-clay flex flex-col items-center rounded-2xl border border-border bg-card p-4 text-center transition-all"
          >
            <div className="flex h-10 w-10 items-center justify-center rounded-full bg-primary/10 text-primary">
              <Icon className="h-5 w-5" />
            </div>
            <p className="mt-2 font-heading text-xs font-bold text-foreground">
              {achievement.title}
            </p>
            <p className="text-xs text-muted-foreground">
              {achievement.description}
            </p>
          </div>
        )
      })}
    </div>
  )
}
