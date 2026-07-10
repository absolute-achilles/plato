"use client"

import { useMemo, useState } from "react"
import { Eye, Mail, MoreHorizontal, TrendingUp } from "lucide-react"

import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import { allStudents } from "@/lib/mocks/data"
import type { Student } from "@/lib/mocks/types"

import { DetailDialog } from "../_components/DetailDialog"
import { PageHeader, StatPill } from "../_components/PageHeader"
import { PaginationControls } from "../_components/PaginationControls"
import { SearchFilters } from "../_components/SearchFilters"
import { StatusBadge } from "../_components/StatusBadge"

const PAGE_SIZE = 10

const grades = ["All", "7", "8", "9", "10", "11", "12"]
const statuses = ["All", "active", "inactive", "pending", "suspended"]
const ranks = ["All", "Bronze", "Silver", "Gold", "Platinum"]

const sortOptions = [
  { label: "Name A-Z", value: "name-asc" },
  { label: "Name Z-A", value: "name-desc" },
  { label: "Highest GPA", value: "gpa" },
  { label: "Best Attendance", value: "attendance" },
  { label: "Most Points", value: "points" },
]

export default function StudentsPage() {
  const [search, setSearch] = useState("")
  const [gradeFilter, setGradeFilter] = useState("All")
  const [statusFilter, setStatusFilter] = useState("All")
  const [rankFilter, setRankFilter] = useState("All")
  const [sortBy, setSortBy] = useState("name-asc")
  const [page, setPage] = useState(1)
  const [selectedStudent, setSelectedStudent] = useState<Student | null>(null)

  const filtered = useMemo(() => {
    let data = [...allStudents]

    if (search.trim()) {
      const q = search.toLowerCase()
      data = data.filter(
        (s) =>
          s.name.toLowerCase().includes(q) ||
          s.email.toLowerCase().includes(q) ||
          s.username.toLowerCase().includes(q)
      )
    }

    if (gradeFilter !== "All") {
      data = data.filter((s) => s.grade === gradeFilter)
    }
    if (statusFilter !== "All") {
      data = data.filter((s) => s.status === statusFilter)
    }
    if (rankFilter !== "All") {
      data = data.filter((s) => s.rank === rankFilter)
    }

    switch (sortBy) {
      case "name-desc":
        data.sort((a, b) => b.name.localeCompare(a.name))
        break
      case "gpa":
        data.sort((a, b) => (b.gpa || 0) - (a.gpa || 0))
        break
      case "attendance":
        data.sort((a, b) => (b.attendanceRate || 0) - (a.attendanceRate || 0))
        break
      case "points":
        data.sort((a, b) => (b.totalPoints || 0) - (a.totalPoints || 0))
        break
      case "name-asc":
      default:
        data.sort((a, b) => a.name.localeCompare(b.name))
    }

    return data
  }, [search, gradeFilter, statusFilter, rankFilter, sortBy])

  const totalPages = Math.ceil(filtered.length / PAGE_SIZE)
  const paginated = filtered.slice((page - 1) * PAGE_SIZE, page * PAGE_SIZE)

  return (
    <div className="mx-auto max-w-7xl space-y-6">
      <PageHeader
        title="Students"
        description="View and manage all enrolled students."
      >
        <StatPill label="Total" value={filtered.length} />
      </PageHeader>

      <SearchFilters
        search={search}
        onSearchChange={(v) => {
          setSearch(v)
          setPage(1)
        }}
        searchPlaceholder="Search students, emails..."
        filters={[
          {
            key: "grade",
            label: "Grade",
            value: gradeFilter,
            onChange: (v) => {
              setGradeFilter(v)
              setPage(1)
            },
            options: grades.map((g) => ({ label: g, value: g })),
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
          {
            key: "rank",
            label: "Rank",
            value: rankFilter,
            onChange: (v) => {
              setRankFilter(v)
              setPage(1)
            },
            options: ranks.map((r) => ({ label: r, value: r })),
          },
        ]}
        sortOptions={sortOptions}
        sortValue={sortBy}
        onSortChange={setSortBy}
      />

      <div className="clay-card overflow-hidden border-0 p-0">
        <Table>
          <TableHeader>
            <TableRow className="border-b border-border hover:bg-transparent">
              <TableHead className="w-[250px]">Student</TableHead>
              <TableHead>Grade</TableHead>
              <TableHead>Status</TableHead>
              <TableHead>GPA</TableHead>
              <TableHead>Attendance</TableHead>
              <TableHead>Rank</TableHead>
              <TableHead className="text-right">Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {paginated.length === 0 ? (
              <TableRow>
                <TableCell
                  colSpan={7}
                  className="py-16 text-center text-muted-foreground"
                >
                  No students found.
                </TableCell>
              </TableRow>
            ) : (
              paginated.map((student) => (
                <TableRow
                  key={student.id}
                  className="cursor-pointer border-b border-border transition-colors hover:bg-sidebar-accent"
                  onClick={() => setSelectedStudent(student)}
                >
                  <TableCell>
                    <div className="flex items-center gap-3">
                      <Avatar className="h-10 w-10 border-2 border-border">
                        <AvatarImage
                          src={student.avatarUrl}
                          alt={student.name}
                        />
                        <AvatarFallback className="bg-primary text-primary-foreground">
                          {student.name.charAt(0)}
                        </AvatarFallback>
                      </Avatar>
                      <div className="min-w-0">
                        <p className="truncate font-semibold text-foreground">
                          {student.name}
                        </p>
                        <p className="truncate text-sm text-muted-foreground">
                          {student.email}
                        </p>
                      </div>
                    </div>
                  </TableCell>
                  <TableCell className="text-muted-foreground">
                    Grade {student.grade}
                  </TableCell>
                  <TableCell>
                    <StatusBadge status={student.status || "active"} />
                  </TableCell>
                  <TableCell className="font-semibold text-foreground">
                    {student.gpa}
                  </TableCell>
                  <TableCell>
                    <div className="flex items-center gap-1 text-emerald-600">
                      <TrendingUp className="h-3.5 w-3.5" />
                      {student.attendanceRate}%
                    </div>
                  </TableCell>
                  <TableCell className="text-muted-foreground">
                    {student.rank}
                  </TableCell>
                  <TableCell className="text-right">
                    <DropdownMenu>
                      <DropdownMenuTrigger asChild>
                        <Button
                          variant="ghost"
                          size="icon"
                          className="h-8 w-8 rounded-lg"
                          onClick={(e) => e.stopPropagation()}
                        >
                          <MoreHorizontal className="h-4 w-4" />
                        </Button>
                      </DropdownMenuTrigger>
                      <DropdownMenuContent align="end">
                        <DropdownMenuItem
                          onClick={() => setSelectedStudent(student)}
                        >
                          <Eye className="mr-2 h-4 w-4" />
                          View Details
                        </DropdownMenuItem>
                        <DropdownMenuItem>
                          <Mail className="mr-2 h-4 w-4" />
                          Send Email
                        </DropdownMenuItem>
                      </DropdownMenuContent>
                    </DropdownMenu>
                  </TableCell>
                </TableRow>
              ))
            )}
          </TableBody>
        </Table>
      </div>

      <PaginationControls
        currentPage={page}
        totalPages={totalPages}
        onPageChange={setPage}
        className="pt-4"
      />

      <DetailDialog
        open={!!selectedStudent}
        onOpenChange={(open) => !open && setSelectedStudent(null)}
        title={selectedStudent?.name || "Student Details"}
        description={selectedStudent?.bio}
        fields={
          selectedStudent
            ? [
                {
                  key: "name",
                  label: "Full Name",
                  value: selectedStudent.name,
                },
                {
                  key: "email",
                  label: "Email",
                  value: selectedStudent.email,
                  type: "email",
                },
                {
                  key: "grade",
                  label: "Grade",
                  value: selectedStudent.grade || "",
                },
                {
                  key: "level",
                  label: "Level",
                  value: selectedStudent.level || "",
                },
                {
                  key: "gpa",
                  label: "GPA",
                  value: String(selectedStudent.gpa || 0),
                },
                {
                  key: "phone",
                  label: "Phone",
                  value: selectedStudent.phone || "",
                  type: "tel",
                },
                {
                  key: "notes",
                  label: "Notes",
                  value: selectedStudent.notes || "",
                  type: "textarea",
                },
              ]
            : []
        }
      />
    </div>
  )
}
