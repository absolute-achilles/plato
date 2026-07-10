"use client"

import {
  BookOpen,
  CalendarCheck,
  GraduationCap,
  LayoutDashboard,
  Settings,
  Shield,
  Trophy,
  UsersRound,
} from "lucide-react"
import Link from "next/link"

import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarSeparator,
} from "@/components/ui/sidebar"
import { useRole } from "@/components/role-provider"
import { cn } from "@/lib/utils"
import { usersByRole } from "@/lib/mocks/data"
import type { Role } from "@/lib/mocks/types"

const roleLabels: Record<Role, string> = {
  student: "Student",
  teacher: "Teacher",
  parent: "Parent",
  admin: "Admin",
}

const roleColors: Record<Role, string> = {
  student: "bg-blue-500",
  teacher: "bg-primary",
  parent: "bg-accent",
  admin: "bg-purple-500",
}

const navigation: Record<
  Role,
  { title: string; url: string; icon: React.ElementType }[]
> = {
  teacher: [
    { title: "Dashboard", url: "/", icon: LayoutDashboard },
    { title: "My Courses", url: "/", icon: BookOpen },
    { title: "Students", url: "/", icon: UsersRound },
    { title: "Attendance", url: "/", icon: CalendarCheck },
    { title: "Schedule", url: "/", icon: GraduationCap },
  ],
  student: [
    { title: "Dashboard", url: "/", icon: LayoutDashboard },
    { title: "My Courses", url: "/", icon: BookOpen },
    { title: "Schedule", url: "/", icon: GraduationCap },
    { title: "Achievements", url: "/", icon: Trophy },
  ],
  parent: [
    { title: "Dashboard", url: "/", icon: LayoutDashboard },
    { title: "My Children", url: "/", icon: UsersRound },
    { title: "Progress", url: "/", icon: Trophy },
    { title: "Schedule", url: "/", icon: GraduationCap },
  ],
  admin: [
    { title: "Dashboard", url: "/", icon: LayoutDashboard },
    { title: "Users", url: "/", icon: UsersRound },
    { title: "Courses", url: "/", icon: BookOpen },
    { title: "Analytics", url: "/", icon: Shield },
  ],
}

export default function DashboardSidebar() {
  const { role, setRole } = useRole()
  const user = usersByRole[role]
  const items = navigation[role]

  return (
    <Sidebar className="border-r border-sidebar-border bg-sidebar">
      <SidebarHeader className="p-4">
        <Link
          href="/"
          className="flex items-center gap-3 rounded-2xl bg-primary px-4 py-3 text-primary-foreground shadow-clay transition-transform hover:scale-[1.02]"
        >
          <GraduationCap className="h-6 w-6" />
          <span className="font-heading text-xl font-bold">Plato</span>
        </Link>

        <div className="mt-4 rounded-2xl border border-sidebar-border bg-sidebar-accent p-3">
          <span className="mb-2 block text-xs font-semibold uppercase tracking-wide text-muted-foreground">
            Preview role
          </span>
          <div className="grid grid-cols-2 gap-2">
            {(Object.keys(roleLabels) as Role[]).map((r) => (
              <button
                key={r}
                onClick={() => setRole(r)}
                className={cn(
                  "rounded-xl px-2 py-1.5 text-xs font-semibold transition-all",
                  role === r
                    ? "bg-primary text-primary-foreground shadow-clay-sm"
                    : "bg-transparent text-sidebar-foreground hover:bg-sidebar-border"
                )}
              >
                {roleLabels[r]}
              </button>
            ))}
          </div>
        </div>
      </SidebarHeader>

      <SidebarSeparator className="bg-sidebar-border" />

      <SidebarContent className="p-3">
        <SidebarGroup>
          <SidebarGroupLabel className="text-muted-foreground">
            Menu
          </SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu className="gap-1.5">
              {items.map((item) => (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton asChild>
                    <Link
                      href={item.url}
                      className="rounded-xl text-sidebar-foreground transition-all hover:bg-sidebar-accent hover:text-sidebar-foreground"
                    >
                      <item.icon className="h-5 w-5" />
                      <span className="font-medium">{item.title}</span>
                    </Link>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>

      <SidebarFooter className="mt-auto border-t border-sidebar-border p-4">
        <div className="flex items-center gap-3 rounded-2xl bg-sidebar-accent p-3">
          <div
            className={cn(
              "flex h-10 w-10 items-center justify-center rounded-full text-sm font-bold text-white",
              roleColors[role]
            )}
          >
            {user.name.charAt(0)}
          </div>
          <div className="flex min-w-0 flex-1 flex-col">
            <span className="truncate text-sm font-bold text-sidebar-foreground">
              {user.name}
            </span>
            <span className="truncate text-xs text-muted-foreground">
              {roleLabels[role]}
            </span>
          </div>
          <button className="rounded-lg p-2 text-sidebar-foreground transition-colors hover:bg-sidebar-border">
            <Settings className="h-4 w-4" />
          </button>
        </div>
      </SidebarFooter>
    </Sidebar>
  )
}
