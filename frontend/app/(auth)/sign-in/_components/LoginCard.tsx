"use client"

import { GraduationCap } from "lucide-react"
import { useState } from "react"

import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {
  Field,
  FieldDescription,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { cn } from "@/lib/utils"

export function LoginCard() {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")

  return (
    <Card className="shadow-clay-lg w-full max-w-md overflow-hidden rounded-[2rem] border-border bg-card">
      <CardHeader className="space-y-1 bg-primary p-8 text-primary-foreground">
        <div className="mx-auto flex h-16 w-16 items-center justify-center rounded-2xl bg-white/20 backdrop-blur-sm">
          <GraduationCap className="h-8 w-8" />
        </div>
        <CardTitle className="pt-2 text-center font-heading text-2xl font-bold">
          Welcome back to Plato
        </CardTitle>
        <CardDescription className="text-center text-primary-foreground/90">
          Sign in to continue your learning journey
        </CardDescription>
      </CardHeader>

      <CardContent className="p-8">
        <FieldGroup className="space-y-5">
          <Field>
            <FieldLabel htmlFor="email">Email</FieldLabel>
            <Input
              id="email"
              type="email"
              placeholder="name@example.com"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="clay-input h-12 px-4"
            />
          </Field>
          <Field>
            <FieldLabel htmlFor="password">Password</FieldLabel>
            <Input
              id="password"
              type="password"
              placeholder="••••••••"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="clay-input h-12 px-4"
            />
            <FieldDescription>
              Use a strong password to keep your account safe.
            </FieldDescription>
          </Field>
          <Button
            type="submit"
            className={cn(
              "clay-button h-12 w-full bg-accent font-heading text-base font-bold text-accent-foreground hover:bg-accent/90"
            )}
          >
            Sign In
          </Button>
        </FieldGroup>
      </CardContent>

      <CardFooter className="flex flex-col gap-3 border-t border-border bg-muted/30 px-8 py-6">
        <Button
          variant="outline"
          className="clay-button h-11 w-full border-border bg-card font-semibold text-foreground hover:bg-muted"
        >
          Sign in with Google
        </Button>
        <p className="text-center text-sm text-muted-foreground">
          Don&apos;t have an account?{" "}
          <a
            href="#"
            className="font-bold text-primary underline-offset-2 hover:underline"
          >
            Sign up
          </a>
        </p>
      </CardFooter>
    </Card>
  )
}
