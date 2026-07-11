import type { APIEnvelope } from "./types"

export class APIError extends Error {
  code: string
  status: number

  constructor(code: string, message: string, status: number) {
    super(message)
    this.code = code
    this.status = status
    this.name = "APIError"
  }
}

export async function apiFetch<T = unknown>(
  input: string,
  init?: RequestInit
): Promise<T> {
  const res = await fetch(input, {
    ...init,
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
      ...init?.headers,
    },
  })

  const body = (await res.json().catch(() => ({}))) as APIEnvelope<T>

  if (!res.ok || !body.success) {
    throw new APIError(
      body.error?.code ?? "UNKNOWN_ERROR",
      body.error?.message ?? `Request failed with status ${res.status}`,
      res.status
    )
  }

  return body.data as T
}

export async function apiGet<T = unknown>(input: string): Promise<T> {
  return apiFetch<T>(input, { method: "GET" })
}

export async function apiPost<T = unknown, B = Record<string, unknown>>(
  input: string,
  body: B
): Promise<T> {
  return apiFetch<T>(input, {
    method: "POST",
    body: JSON.stringify(body),
  })
}

export async function apiPatch<T = unknown, B = Record<string, unknown>>(
  input: string,
  body: B
): Promise<T> {
  return apiFetch<T>(input, {
    method: "PATCH",
    body: JSON.stringify(body),
  })
}

export async function apiDelete<T = unknown>(input: string): Promise<T> {
  return apiFetch<T>(input, { method: "DELETE" })
}
