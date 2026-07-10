import type { Role, User } from "../types"
import { usersByRole } from "../data"

let currentRole: Role = "teacher"

export function setCurrentRole(role: Role): void {
  currentRole = role
}

export function getCurrentRole(): Role {
  return currentRole
}

export function getCurrentUser(): User {
  return usersByRole[currentRole]
}

export async function fetchCurrentUser(): Promise<User> {
  await delay(300)
  return getCurrentUser()
}

function delay(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms))
}
