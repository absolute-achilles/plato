"use client"

import { useRouter } from "next/navigation"
import {
  createContext,
  useContext,
  useEffect,
  useState,
  type ReactNode,
} from "react"

import { useRole } from "@/components/role-provider"
import { fetchCurrentUser } from "@/lib/api/auth"
import type { Role, UserResponse } from "@/lib/api/types"

interface AuthContextValue {
  user: UserResponse | null
  role: Role | null
  isLoading: boolean
  error: Error | null
}

const AuthContext = createContext<AuthContextValue | null>(null)

export function AuthProvider({ children }: { children: ReactNode }) {
  const router = useRouter()
  const { setRole } = useRole()
  const [user, setUser] = useState<UserResponse | null>(null)
  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState<Error | null>(null)

  useEffect(() => {
    let cancelled = false

    fetchCurrentUser()
      .then((u) => {
        if (cancelled) return
        setUser(u)
        setRole(u.role)
        setIsLoading(false)
      })
      .catch((err) => {
        if (cancelled) return
        setError(err as Error)
        setIsLoading(false)
        router.replace("/sign-in")
      })

    return () => {
      cancelled = true
    }
  }, [router, setRole])

  return (
    <AuthContext.Provider
      value={{ user, role: user?.role ?? null, isLoading, error }}
    >
      {children}
    </AuthContext.Provider>
  )
}

export function useAuth(): AuthContextValue {
  const context = useContext(AuthContext)
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider")
  }
  return context
}
