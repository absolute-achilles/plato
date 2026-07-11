"use client"

import { Loader2 } from "lucide-react"
import { useState } from "react"
import { toast } from "sonner"

import { Button } from "@/components/ui/button"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import {
  Field,
  FieldDescription,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import { createTeacher } from "@/lib/api/admin"
import type { Department, TeacherResponse } from "@/lib/api/types"

const departments: Department[] = [
  "Mathematics",
  "Science",
  "English",
  "History",
  "Arts",
  "Physical Education",
  "Computer Science",
  "Other",
]

interface CreateTeacherDialogProps {
  onCreated?: (teacher: TeacherResponse) => void
}

export function CreateTeacherDialog({ onCreated }: CreateTeacherDialogProps) {
  const [open, setOpen] = useState(false)
  const [isLoading, setIsLoading] = useState(false)
  const [username, setUsername] = useState("")
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [phone, setPhone] = useState("")
  const [department, setDepartment] = useState<Department>("Mathematics")

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    setIsLoading(true)

    try {
      const teacher = await createTeacher({
        username,
        email,
        password,
        phone_number: phone || undefined,
        department,
      })
      toast.success("Teacher created successfully")
      onCreated?.(teacher)
      setOpen(false)
      setUsername("")
      setEmail("")
      setPassword("")
      setPhone("")
      setDepartment("Mathematics")
    } catch (err) {
      toast.error(err instanceof Error ? err.message : "Failed to create teacher")
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button>Add Teacher</Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-md">
        <form onSubmit={handleSubmit}>
          <DialogHeader>
            <DialogTitle>Create Teacher</DialogTitle>
            <DialogDescription>
              Add a new teacher to the platform. They will use the temporary
              password on first sign-in.
            </DialogDescription>
          </DialogHeader>

          <FieldGroup className="space-y-4 py-4">
            <Field>
              <FieldLabel htmlFor="username">Username</FieldLabel>
              <Input
                id="username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                required
              />
            </Field>

            <Field>
              <FieldLabel htmlFor="email">Email</FieldLabel>
              <Input
                id="email"
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                required
              />
            </Field>

            <Field>
              <FieldLabel htmlFor="password">Temporary Password</FieldLabel>
              <Input
                id="password"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                required
                minLength={8}
              />
              <FieldDescription>Must be at least 8 characters.</FieldDescription>
            </Field>

            <Field>
              <FieldLabel htmlFor="phone">Phone Number</FieldLabel>
              <Input
                id="phone"
                type="tel"
                value={phone}
                onChange={(e) => setPhone(e.target.value)}
              />
            </Field>

            <Field>
              <FieldLabel htmlFor="department">Department</FieldLabel>
              <Select
                value={department}
                onValueChange={(value) => setDepartment(value as Department)}
              >
                <SelectTrigger id="department" className="w-full">
                  <SelectValue placeholder="Select department" />
                </SelectTrigger>
                <SelectContent>
                  {departments.map((d) => (
                    <SelectItem key={d} value={d}>
                      {d}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </Field>
          </FieldGroup>

          <DialogFooter>
            <Button
              type="button"
              variant="outline"
              onClick={() => setOpen(false)}
              disabled={isLoading}
            >
              Cancel
            </Button>
            <Button type="submit" disabled={isLoading}>
              {isLoading ? (
                <Loader2 className="mr-2 h-4 w-4 animate-spin" />
              ) : null}
              Create Teacher
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  )
}
