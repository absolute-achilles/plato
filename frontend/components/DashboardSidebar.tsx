import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarSeparator,
} from "@/components/ui/sidebar"
import {
  CalendarCheck,
  GraduationCap,
  Home,
  LayoutDashboard,
  UsersRound,
} from "lucide-react"
import Image from "next/image"
import Link from "next/link"

const teacherSidebarItems = [
  {
    title: "Dashboard",
    url: "/",
    icon: LayoutDashboard,
  },
  {
    title: "Courses",
    url: "/",
    icon: GraduationCap,
  },
  {
    title: "Students",
    url: "/",
    icon: UsersRound,
  },
  {
    title: "Attendance",
    url: "/",
    icon: CalendarCheck,
  },
]

const DashboardSidebar = () => {
  return (
    <Sidebar collapsible="icon">
      {/* Header */}
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuButton asChild>
            <Link href="/">
              <span>Plato</span>
            </Link>
          </SidebarMenuButton>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarSeparator />
      {/* Content */}
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Menu</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu className="gap-2">
              {teacherSidebarItems.map((item) => (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton asChild>
                    <Link href={item.url}>
                      <item.icon />
                      <span>{item.title}</span>
                    </Link>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
        <SidebarGroup></SidebarGroup>
      </SidebarContent>
    </Sidebar>
  )
}

export default DashboardSidebar
