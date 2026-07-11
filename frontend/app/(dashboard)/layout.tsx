import { TooltipProvider } from "@/components/ui/tooltip"
import { SidebarProvider } from "@/components/ui/sidebar"
import { AuthProvider } from "@/components/auth-provider"
import DashboardSidebar from "@/components/DashboardSidebar"
import { RoleProvider } from "@/components/role-provider"

export default function DashboardLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <RoleProvider>
      <AuthProvider>
        <SidebarProvider>
          <TooltipProvider>
            <DashboardSidebar />
            <main className="min-h-screen flex-1 bg-background p-6 lg:p-8">
              {children}
            </main>
          </TooltipProvider>
        </SidebarProvider>
      </AuthProvider>
    </RoleProvider>
  )
}
