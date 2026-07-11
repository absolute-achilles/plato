import { apiGet, apiPost } from "./client"
import type { LoginRequest, TokenResponse, UserResponse } from "./types"

export async function login(req: LoginRequest): Promise<TokenResponse> {
  return apiPost<TokenResponse, LoginRequest>("/api/v1/auth/login", req)
}

export async function fetchCurrentUser(): Promise<UserResponse> {
  return apiGet<UserResponse>("/api/v1/auth/me")
}

export async function changePassword(req: {
  old_password: string
  new_password: string
}): Promise<void> {
  await apiPost("/api/v1/auth/change-password", req)
}
