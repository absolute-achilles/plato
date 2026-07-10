"use client"

import { useMemo, useState } from "react"
import {
  BookOpen,
  CheckCircle,
  Clock,
  FileText,
  Library,
  Trophy,
} from "lucide-react"

import { Button } from "@/components/ui/button"
import { Card, CardContent, CardFooter } from "@/components/ui/card"
import { allCourses, allModules } from "@/lib/mocks/data"
import type { Module } from "@/lib/mocks/types"

import { DetailDialog } from "../_components/DetailDialog"
import { PageHeader, StatPill } from "../_components/PageHeader"
import { PaginationControls } from "../_components/PaginationControls"
import { SearchFilters } from "../_components/SearchFilters"
import { StatusBadge } from "../_components/StatusBadge"

const PAGE_SIZE = 12

const courses = ["All", ...allCourses.map((c) => c.name)]
const difficulties = ["All", "beginner", "intermediate", "advanced"]
const statuses = ["All", "draft", "published", "archived"]

const sortOptions = [
  { label: "Position", value: "position" },
  { label: "Name A-Z", value: "name-asc" },
  { label: "Name Z-A", value: "name-desc" },
  { label: "Most Content", value: "content" },
  { label: "Highest Rated", value: "rating" },
]

export default function ModulesPage() {
  const [search, setSearch] = useState("")
  const [courseFilter, setCourseFilter] = useState("All")
  const [difficultyFilter, setDifficultyFilter] = useState("All")
  const [statusFilter, setStatusFilter] = useState("All")
  const [sortBy, setSortBy] = useState("position")
  const [page, setPage] = useState(1)
  const [selectedModule, setSelectedModule] = useState<Module | null>(null)

  const courseMap = useMemo(() => {
    const map = new Map<string, string>()
    allCourses.forEach((c) => map.set(c.id, c.name))
    return map
  }, [])

  const filtered = useMemo(() => {
    let data = [...allModules]

    if (search.trim()) {
      const q = search.toLowerCase()
      data = data.filter(
        (m) =>
          m.name.toLowerCase().includes(q) ||
          (m.description || "").toLowerCase().includes(q) ||
          (courseMap.get(m.courseId) || "").toLowerCase().includes(q)
      )
    }

    if (courseFilter !== "All") {
      const courseId = allCourses.find((c) => c.name === courseFilter)?.id
      if (courseId) {
        data = data.filter((m) => m.courseId === courseId)
      }
    }
    if (difficultyFilter !== "All") {
      data = data.filter((m) => m.difficulty === difficultyFilter)
    }
    if (statusFilter !== "All") {
      data = data.filter((m) => m.status === statusFilter)
    }

    switch (sortBy) {
      case "name-asc":
        data.sort((a, b) => a.name.localeCompare(b.name))
        break
      case "name-desc":
        data.sort((a, b) => b.name.localeCompare(a.name))
        break
      case "content":
        data.sort((a, b) => (b.contentCount || 0) - (a.contentCount || 0))
        break
      case "rating":
        data.sort((a, b) => (b.averageRating || 0) - (a.averageRating || 0))
        break
      case "position":
      default:
        data.sort((a, b) => a.position - b.position)
    }

    return data
  }, [search, courseFilter, difficultyFilter, statusFilter, sortBy, courseMap])

  const totalPages = Math.ceil(filtered.length / PAGE_SIZE)
  const paginated = filtered.slice((page - 1) * PAGE_SIZE, page * PAGE_SIZE)

  return (
    <div className="mx-auto max-w-7xl space-y-6">
      <PageHeader
        title="Modules"
        description="Explore and manage learning modules across all courses."
      >
        <StatPill label="Total" value={filtered.length} />
      </PageHeader>

      <SearchFilters
        search={search}
        onSearchChange={(v) => {
          setSearch(v)
          setPage(1)
        }}
        searchPlaceholder="Search modules, courses..."
        filters={[
          {
            key: "course",
            label: "Course",
            value: courseFilter,
            onChange: (v) => {
              setCourseFilter(v)
              setPage(1)
            },
            options: courses.map((c) => ({ label: c, value: c })),
          },
          {
            key: "difficulty",
            label: "Difficulty",
            value: difficultyFilter,
            onChange: (v) => {
              setDifficultyFilter(v)
              setPage(1)
            },
            options: difficulties.map((d) => ({ label: d, value: d })),
          },
          {
            key: "status",
            label: "Status",
            value: statusFilter,
            onChange: (v) => {
              setStatusFilter(v)
              setPage(1)
            },
            options: statuses.map((s) => ({ label: s, value: s })),
          },
        ]}
        sortOptions={sortOptions}
        sortValue={sortBy}
        onSortChange={setSortBy}
      />

      {paginated.length === 0 ? (
        <div className="clay-card py-16 text-center">
          <Library className="mx-auto h-12 w-12 text-muted-foreground" />
          <p className="mt-4 text-muted-foreground">No modules found.</p>
        </div>
      ) : (
        <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
          {paginated.map((module) => (
            <Card
              key={module.id}
              className="clay-card flex cursor-pointer flex-col border-0 transition-all hover:-translate-y-1"
              onClick={() => setSelectedModule(module)}
            >
              <CardContent className="flex flex-1 flex-col gap-3 p-5">
                <div className="flex items-start justify-between gap-2">
                  <div className="shadow-clay-sm flex h-10 w-10 items-center justify-center rounded-2xl bg-primary text-primary-foreground">
                    <BookOpen className="h-5 w-5" />
                  </div>
                  <StatusBadge status={module.status || "published"} />
                </div>

                <div>
                  <span className="text-xs font-semibold tracking-wide text-muted-foreground uppercase">
                    {courseMap.get(module.courseId) || "Unknown Course"}
                  </span>
                  <h3 className="font-heading text-lg font-bold text-foreground">
                    {module.name}
                  </h3>
                </div>

                <p className="line-clamp-2 text-sm text-muted-foreground">
                  {module.shortDescription || module.description}
                </p>

                <div className="mt-auto flex flex-wrap gap-2 pt-3">
                  <span className="inline-flex items-center gap-1 rounded-full bg-muted px-2 py-1 text-xs font-semibold text-muted-foreground">
                    <FileText className="h-3 w-3" />
                    {module.contentCount} items
                  </span>
                  {module.isMandatory && (
                    <span className="inline-flex items-center gap-1 rounded-full bg-amber-100 px-2 py-1 text-xs font-semibold text-amber-700">
                      <Trophy className="h-3 w-3" />
                      Required
                    </span>
                  )}
                  {module.certificateEligible && (
                    <span className="inline-flex items-center gap-1 rounded-full bg-emerald-100 px-2 py-1 text-xs font-semibold text-emerald-700">
                      <CheckCircle className="h-3 w-3" />
                      Cert
                    </span>
                  )}
                </div>

                <div className="flex items-center justify-between text-sm text-muted-foreground">
                  <div className="flex items-center gap-2">
                    <Clock className="h-4 w-4" />
                    {module.duration}m
                  </div>
                  <div className="font-semibold text-foreground">
                    Pos {module.position}
                  </div>
                </div>
              </CardContent>
              <CardFooter className="border-t border-border p-4">
                <Button
                  variant="outline"
                  className="clay-button w-full rounded-xl border-border"
                  onClick={(e) => {
                    e.stopPropagation()
                    setSelectedModule(module)
                  }}
                >
                  View Details
                </Button>
              </CardFooter>
            </Card>
          ))}
        </div>
      )}

      <PaginationControls
        currentPage={page}
        totalPages={totalPages}
        onPageChange={setPage}
        className="pt-4"
      />

      <DetailDialog
        open={!!selectedModule}
        onOpenChange={(open) => !open && setSelectedModule(null)}
        title={selectedModule?.name || "Module Details"}
        description={selectedModule?.description}
        fields={
          selectedModule
            ? [
                {
                  key: "name",
                  label: "Module Name",
                  value: selectedModule.name,
                },
                {
                  key: "course",
                  label: "Course",
                  value: courseMap.get(selectedModule.courseId) || "",
                },
                {
                  key: "difficulty",
                  label: "Difficulty",
                  value: selectedModule.difficulty || "",
                },
                {
                  key: "position",
                  label: "Position",
                  value: String(selectedModule.position),
                },
                {
                  key: "duration",
                  label: "Duration (minutes)",
                  value: String(selectedModule.duration || 0),
                },
                {
                  key: "description",
                  label: "Description",
                  value: selectedModule.description || "",
                  type: "textarea",
                },
              ]
            : []
        }
      />
    </div>
  )
}
