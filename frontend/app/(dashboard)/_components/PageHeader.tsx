import { cn } from "@/lib/utils"

interface PageHeaderProps {
  title: string
  description?: string
  children?: React.ReactNode
}

export function PageHeader({ title, description, children }: PageHeaderProps) {
  return (
    <div className="flex flex-col justify-between gap-4 sm:flex-row sm:items-center">
      <div>
        <h1 className="font-heading text-2xl font-bold text-foreground sm:text-3xl">
          {title}
        </h1>
        {description && (
          <p className="mt-1 text-sm text-muted-foreground">{description}</p>
        )}
      </div>
      {children && (
        <div className="flex flex-wrap items-center gap-2">{children}</div>
      )}
    </div>
  )
}

interface StatPillProps {
  label: string
  value: string | number
  className?: string
}

export function StatPill({ label, value, className }: StatPillProps) {
  return (
    <div
      className={cn(
        "shadow-clay-sm rounded-2xl border border-border bg-card px-4 py-2 text-sm",
        className
      )}
    >
      <span className="text-muted-foreground">{label}</span>
      <span className="ml-2 font-bold text-foreground">{value}</span>
    </div>
  )
}
