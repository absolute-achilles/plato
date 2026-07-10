"use client"

import { ChevronLeft, ChevronRight } from "lucide-react"

import { Button } from "@/components/ui/button"
import { cn } from "@/lib/utils"

interface PaginationControlsProps {
  currentPage: number
  totalPages: number
  onPageChange: (page: number) => void
  className?: string
}

export function PaginationControls({
  currentPage,
  totalPages,
  onPageChange,
  className,
}: PaginationControlsProps) {
  if (totalPages <= 1) return null

  const pages = Array.from({ length: totalPages }, (_, i) => i + 1)
  const maxVisible = 5

  let visiblePages = pages
  if (totalPages > maxVisible) {
    const start = Math.max(
      1,
      Math.min(currentPage - 2, totalPages - maxVisible + 1)
    )
    visiblePages = pages.slice(start - 1, start + maxVisible - 1)
  }

  return (
    <div className={cn("flex items-center justify-center gap-1", className)}>
      <Button
        variant="outline"
        size="icon"
        onClick={() => onPageChange(currentPage - 1)}
        disabled={currentPage === 1}
        className="clay-button h-9 w-9 rounded-xl border-border"
      >
        <ChevronLeft className="h-4 w-4" />
      </Button>

      {visiblePages.map((page) => (
        <Button
          key={page}
          variant={currentPage === page ? "default" : "outline"}
          size="icon"
          onClick={() => onPageChange(page)}
          className={cn(
            "clay-button h-9 w-9 rounded-xl border-border",
            currentPage === page &&
              "bg-primary text-primary-foreground hover:bg-primary/90"
          )}
        >
          {page}
        </Button>
      ))}

      <Button
        variant="outline"
        size="icon"
        onClick={() => onPageChange(currentPage + 1)}
        disabled={currentPage === totalPages}
        className="clay-button h-9 w-9 rounded-xl border-border"
      >
        <ChevronRight className="h-4 w-4" />
      </Button>
    </div>
  )
}
