import {
  Card,
  CardAction,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"

import { Button } from "@/components/ui/button"
import {
  Field,
  FieldDescription,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field"
import { Input } from "@/components/ui/input"

export const LoginCard = () => {
  return (
    <Card className="w-[50%]">
      <CardHeader>
        <CardTitle>Login</CardTitle>
        <CardDescription>Welcome Back! Ready to learn?</CardDescription>
        <CardAction>Sign Up</CardAction>
      </CardHeader>
      <CardContent>
        <FieldGroup>
          <Field>
            <FieldLabel htmlFor="fieldgroup-name">Name</FieldLabel>
            <Input id="fieldgroup-name" placeholder="Jordan Lee" />
          </Field>
          <Field>
            <FieldLabel htmlFor="fieldgroup-email">Email</FieldLabel>
            <Input
              id="fieldgroup-email"
              type="email"
              placeholder="name@example.com"
            />
            <FieldDescription>
              We&apos;ll send updates to this address.
            </FieldDescription>
          </Field>
          <Field orientation="horizontal">
            <Button type="submit">Login</Button>
          </Field>
        </FieldGroup>
      </CardContent>
      <CardFooter>
        <Button>Sign in with Google</Button>
      </CardFooter>
    </Card>
  )
}
