import { ArrowDown, ArrowUp, Minus } from "lucide-react"

import { cn } from "@/lib/utils"
import type { DashboardStats } from "@/lib/mocks/types"

interface StatCardProps {
  stat: DashboardStats
}

export function StatCard({ stat }: StatCardProps) {
  const Icon =
    stat.trend === "up" ? ArrowUp : stat.trend === "down" ? ArrowDown : Minus

  return (
    <div className="clay-card flex flex-col justify-between p-5">
      <span className="text-sm font-medium text-muted-foreground">
        {stat.label}
      </span>
      <div className="mt-2 flex items-end justify-between">
        <span className="font-heading text-3xl font-bold text-foreground">
          {stat.value}
        </span>
        {stat.change && (
          <span
            className={cn(
              "flex items-center gap-0.5 rounded-full px-2 py-0.5 text-xs font-bold",
              stat.trend === "up" && "bg-emerald-100 text-emerald-700",
              stat.trend === "down" && "bg-rose-100 text-rose-700",
              stat.trend === "neutral" && "bg-muted text-muted-foreground"
            )}
          >
            <Icon className="h-3 w-3" />
            {stat.change}
          </span>
        )}
      </div>
    </div>
  )
}
