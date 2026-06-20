import { TooltipProvider } from "@/components/ui/tooltip"

import { SidebarProvider } from "@/components/ui/sidebar"
import DashboardSidebar from "@/components/DashboardSidebar"

export default function DashboardLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <SidebarProvider>
      <TooltipProvider>
        <DashboardSidebar />
        <main>{children}</main>
      </TooltipProvider>
    </SidebarProvider>
  )
}
