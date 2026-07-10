"use client"

import { useMemo, useState } from "react"
import { BookOpen, Clock, GraduationCap, Star, UsersRound } from "lucide-react"
import Image from "next/image"

import { Button } from "@/components/ui/button"
import { Card, CardContent, CardFooter } from "@/components/ui/card"
import { allCourses } from "@/lib/mocks/data"
import type { Course } from "@/lib/mocks/types"

import { DetailDialog } from "../_components/DetailDialog"
import { PageHeader, StatPill } from "../_components/PageHeader"
import { PaginationControls } from "../_components/PaginationControls"
import { SearchFilters } from "../_components/SearchFilters"
import { StatusBadge } from "../_components/StatusBadge"

const PAGE_SIZE = 8

const categories = [
  "All",
  "Mathematics",
  "Science",
  "Language",
  "Technology",
  "Arts",
  "Humanities",
  "Business",
  "Health",
]
const levels = ["All", "beginner", "intermediate", "advanced"]
const statuses = ["All", "draft", "published", "archived"]

const sortOptions = [
  { label: "Newest", value: "newest" },
  { label: "Name A-Z", value: "name-asc" },
  { label: "Name Z-A", value: "name-desc" },
  { label: "Most Students", value: "students" },
  { label: "Highest Rated", value: "rating" },
]

export default function CoursesPage() {
  const [search, setSearch] = useState("")
  const [categoryFilter, setCategoryFilter] = useState("All")
  const [levelFilter, setLevelFilter] = useState("All")
  const [statusFilter, setStatusFilter] = useState("All")
  const [sortBy, setSortBy] = useState("newest")
  const [page, setPage] = useState(1)
  const [selectedCourse, setSelectedCourse] = useState<Course | null>(null)

  const filtered = useMemo(() => {
    let data = [...allCourses]

    if (search.trim()) {
      const q = search.toLowerCase()
      data = data.filter(
        (c) =>
          c.name.toLowerCase().includes(q) ||
          c.description.toLowerCase().includes(q) ||
          c.category.toLowerCase().includes(q) ||
          c.teacherName?.toLowerCase().includes(q)
      )
    }

    if (categoryFilter !== "All") {
      data = data.filter((c) => c.category === categoryFilter)
    }
    if (levelFilter !== "All") {
      data = data.filter((c) => c.level === levelFilter)
    }
    if (statusFilter !== "All") {
      data = data.filter((c) => c.status === statusFilter)
    }

    switch (sortBy) {
      case "name-asc":
        data.sort((a, b) => a.name.localeCompare(b.name))
        break
      case "name-desc":
        data.sort((a, b) => b.name.localeCompare(a.name))
        break
      case "students":
        data.sort((a, b) => (b.studentCount || 0) - (a.studentCount || 0))
        break
      case "rating":
        data.sort((a, b) => (b.rating || 0) - (a.rating || 0))
        break
      case "newest":
      default:
        data.sort((a, b) => b.createdAt.getTime() - a.createdAt.getTime())
    }

    return data
  }, [search, categoryFilter, levelFilter, statusFilter, sortBy])

  const totalPages = Math.ceil(filtered.length / PAGE_SIZE)
  const paginated = filtered.slice((page - 1) * PAGE_SIZE, page * PAGE_SIZE)

  return (
    <div className="mx-auto max-w-7xl space-y-6">
      <PageHeader
        title="Courses"
        description="Browse and manage all available learning courses."
      >
        <StatPill label="Total" value={filtered.length} />
      </PageHeader>

      <SearchFilters
        search={search}
        onSearchChange={(v) => {
          setSearch(v)
          setPage(1)
        }}
        searchPlaceholder="Search courses, categories, teachers..."
        filters={[
          {
            key: "category",
            label: "Category",
            value: categoryFilter,
            onChange: (v) => {
              setCategoryFilter(v)
              setPage(1)
            },
            options: categories.map((c) => ({ label: c, value: c })),
          },
          {
            key: "level",
            label: "Level",
            value: levelFilter,
            onChange: (v) => {
              setLevelFilter(v)
              setPage(1)
            },
            options: levels.map((l) => ({ label: l, value: l })),
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
          <BookOpen className="mx-auto h-12 w-12 text-muted-foreground" />
          <p className="mt-4 text-muted-foreground">No courses found.</p>
        </div>
      ) : (
        <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
          {paginated.map((course) => (
            <Card
              key={course.id}
              className="clay-card flex cursor-pointer flex-col overflow-hidden border-0 transition-all hover:-translate-y-1"
              onClick={() => setSelectedCourse(course)}
            >
              <div className="relative aspect-video w-full overflow-hidden bg-muted">
                <Image
                  src={course.thumbnailUrl || course.coverImageUrl || ""}
                  alt={course.name}
                  fill
                  className="object-cover"
                  sizes="(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 25vw"
                />
                <div className="absolute top-3 left-3">
                  <StatusBadge status={course.status || "published"} />
                </div>
              </div>
              <CardContent className="flex flex-1 flex-col gap-3 p-5">
                <div className="flex items-start justify-between gap-2">
                  <span className="rounded-full bg-muted px-2.5 py-1 text-xs font-semibold text-muted-foreground">
                    {course.category}
                  </span>
                  <div className="flex items-center gap-1 text-xs font-semibold text-amber-600">
                    <Star className="h-3.5 w-3.5 fill-current" />
                    {course.rating}
                  </div>
                </div>

                <h3 className="font-heading text-lg font-bold text-foreground">
                  {course.name}
                </h3>
                <p className="line-clamp-2 text-sm text-muted-foreground">
                  {course.shortDescription || course.description}
                </p>

                <div className="mt-auto space-y-2 pt-3">
                  <div className="flex items-center gap-2 text-sm text-muted-foreground">
                    <GraduationCap className="h-4 w-4" />
                    <span>{course.teacherName}</span>
                  </div>
                  <div className="flex items-center justify-between text-sm text-muted-foreground">
                    <div className="flex items-center gap-2">
                      <UsersRound className="h-4 w-4" />
                      {course.studentCount} students
                    </div>
                    <div className="flex items-center gap-2">
                      <Clock className="h-4 w-4" />
                      {course.duration}m
                    </div>
                  </div>
                </div>
              </CardContent>
              <CardFooter className="border-t border-border p-4">
                <Button
                  variant="outline"
                  className="clay-button w-full rounded-xl border-border"
                  onClick={(e) => {
                    e.stopPropagation()
                    setSelectedCourse(course)
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
        open={!!selectedCourse}
        onOpenChange={(open) => !open && setSelectedCourse(null)}
        title={selectedCourse?.name || "Course Details"}
        description={selectedCourse?.description}
        fields={
          selectedCourse
            ? [
                {
                  key: "name",
                  label: "Course Name",
                  value: selectedCourse.name,
                },
                {
                  key: "category",
                  label: "Category",
                  value: selectedCourse.category,
                },
                {
                  key: "level",
                  label: "Level",
                  value: selectedCourse.level || "",
                },
                {
                  key: "teacher",
                  label: "Teacher",
                  value: selectedCourse.teacherName || "",
                },
                {
                  key: "description",
                  label: "Description",
                  value: selectedCourse.description,
                  type: "textarea",
                },
              ]
            : []
        }
      />
    </div>
  )
}
