"use client"

import { useMemo, useState } from "react"
import { Eye, Mail, MoreHorizontal, Star } from "lucide-react"

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
import { allTeachers } from "@/lib/mocks/data"
import type { Teacher } from "@/lib/mocks/types"

import { DetailDialog } from "../../_components/DetailDialog"
import { PageHeader, StatPill } from "../../_components/PageHeader"
import { PaginationControls } from "../../_components/PaginationControls"
import { SearchFilters } from "../../_components/SearchFilters"
import { StatusBadge } from "../../_components/StatusBadge"

const PAGE_SIZE = 10

const departments = [
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
const statuses = ["All", "active", "inactive", "pending", "suspended"]
const verifications = ["All", "verified", "unverified"]

const sortOptions = [
  { label: "Name A-Z", value: "name-asc" },
  { label: "Name Z-A", value: "name-desc" },
  { label: "Most Experienced", value: "experience" },
  { label: "Highest Rated", value: "rating" },
  { label: "Most Students", value: "students" },
]

export default function TeachersPage() {
  const [search, setSearch] = useState("")
  const [departmentFilter, setDepartmentFilter] = useState("All")
  const [statusFilter, setStatusFilter] = useState("All")
  const [verifiedFilter, setVerifiedFilter] = useState("All")
  const [sortBy, setSortBy] = useState("name-asc")
  const [page, setPage] = useState(1)
  const [selectedTeacher, setSelectedTeacher] = useState<Teacher | null>(null)

  const filtered = useMemo(() => {
    let data = [...allTeachers]

    if (search.trim()) {
      const q = search.toLowerCase()
      data = data.filter(
        (t) =>
          t.name.toLowerCase().includes(q) ||
          t.email.toLowerCase().includes(q) ||
          t.username.toLowerCase().includes(q) ||
          (t.department || "").toLowerCase().includes(q)
      )
    }

    if (departmentFilter !== "All") {
      data = data.filter((t) => t.department === departmentFilter)
    }
    if (statusFilter !== "All") {
      data = data.filter((t) => t.status === statusFilter)
    }
    if (verifiedFilter !== "All") {
      data = data.filter((t) =>
        verifiedFilter === "verified" ? t.isVerified : !t.isVerified
      )
    }

    switch (sortBy) {
      case "name-desc":
        data.sort((a, b) => b.name.localeCompare(a.name))
        break
      case "experience":
        data.sort(
          (a, b) => (b.yearsOfExperience || 0) - (a.yearsOfExperience || 0)
        )
        break
      case "rating":
        data.sort((a, b) => (b.averageRating || 0) - (a.averageRating || 0))
        break
      case "students":
        data.sort((a, b) => (b.totalStudents || 0) - (a.totalStudents || 0))
        break
      case "name-asc":
      default:
        data.sort((a, b) => a.name.localeCompare(b.name))
    }

    return data
  }, [search, departmentFilter, statusFilter, verifiedFilter, sortBy])

  const totalPages = Math.ceil(filtered.length / PAGE_SIZE)
  const paginated = filtered.slice((page - 1) * PAGE_SIZE, page * PAGE_SIZE)

  return (
    <div className="mx-auto max-w-7xl space-y-6">
      <PageHeader
        title="Teachers"
        description="Manage and review all platform teachers."
      >
        <StatPill label="Total" value={filtered.length} />
      </PageHeader>

      <SearchFilters
        search={search}
        onSearchChange={(v) => {
          setSearch(v)
          setPage(1)
        }}
        searchPlaceholder="Search teachers, emails, departments..."
        filters={[
          {
            key: "department",
            label: "Department",
            value: departmentFilter,
            onChange: (v) => {
              setDepartmentFilter(v)
              setPage(1)
            },
            options: departments.map((d) => ({ label: d, value: d })),
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
            key: "verified",
            label: "Verified",
            value: verifiedFilter,
            onChange: (v) => {
              setVerifiedFilter(v)
              setPage(1)
            },
            options: verifications.map((v) => ({ label: v, value: v })),
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
              <TableHead className="w-[250px]">Teacher</TableHead>
              <TableHead>Department</TableHead>
              <TableHead>Status</TableHead>
              <TableHead>Experience</TableHead>
              <TableHead>Students</TableHead>
              <TableHead>Rating</TableHead>
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
                  No teachers found.
                </TableCell>
              </TableRow>
            ) : (
              paginated.map((teacher) => (
                <TableRow
                  key={teacher.id}
                  className="cursor-pointer border-b border-border transition-colors hover:bg-sidebar-accent"
                  onClick={() => setSelectedTeacher(teacher)}
                >
                  <TableCell>
                    <div className="flex items-center gap-3">
                      <Avatar className="h-10 w-10 border-2 border-border">
                        <AvatarImage
                          src={teacher.avatarUrl}
                          alt={teacher.name}
                        />
                        <AvatarFallback className="bg-primary text-primary-foreground">
                          {teacher.name.charAt(0)}
                        </AvatarFallback>
                      </Avatar>
                      <div className="min-w-0">
                        <p className="truncate font-semibold text-foreground">
                          {teacher.name}
                        </p>
                        <p className="truncate text-sm text-muted-foreground">
                          {teacher.email}
                        </p>
                      </div>
                    </div>
                  </TableCell>
                  <TableCell className="text-muted-foreground">
                    {teacher.department || "—"}
                  </TableCell>
                  <TableCell>
                    <StatusBadge status={teacher.status || "active"} />
                  </TableCell>
                  <TableCell className="text-muted-foreground">
                    {teacher.yearsOfExperience} years
                  </TableCell>
                  <TableCell className="text-muted-foreground">
                    {teacher.totalStudents}
                  </TableCell>
                  <TableCell>
                    <div className="flex items-center gap-1 font-semibold text-amber-600">
                      <Star className="h-3.5 w-3.5 fill-current" />
                      {teacher.averageRating}
                    </div>
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
                          onClick={() => setSelectedTeacher(teacher)}
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
        open={!!selectedTeacher}
        onOpenChange={(open) => !open && setSelectedTeacher(null)}
        title={selectedTeacher?.name || "Teacher Details"}
        description={selectedTeacher?.bio}
        fields={
          selectedTeacher
            ? [
                {
                  key: "name",
                  label: "Full Name",
                  value: selectedTeacher.name,
                },
                {
                  key: "email",
                  label: "Email",
                  value: selectedTeacher.email,
                  type: "email",
                },
                {
                  key: "department",
                  label: "Department",
                  value: selectedTeacher.department || "",
                },
                {
                  key: "phone",
                  label: "Phone",
                  value: selectedTeacher.phone || "",
                  type: "tel",
                },
                {
                  key: "experience",
                  label: "Years of Experience",
                  value: String(selectedTeacher.yearsOfExperience || 0),
                },
                {
                  key: "bio",
                  label: "Bio",
                  value: selectedTeacher.bio || "",
                  type: "textarea",
                },
              ]
            : []
        }
      />
    </div>
  )
}
