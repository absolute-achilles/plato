import { LoginCard } from "./_components/LoginCard"

export default function SignInPage() {
  return (
    <div className="relative flex min-h-screen items-center justify-center overflow-hidden bg-background p-4">
      {/* Decorative clay blobs */}
      <div className="absolute -left-20 -top-20 h-72 w-72 rounded-full bg-secondary/60 blur-3xl" />
      <div className="absolute -bottom-24 -right-24 h-80 w-80 rounded-full bg-accent/20 blur-3xl" />
      <div className="absolute left-1/3 top-1/4 h-40 w-40 rounded-full bg-primary/20 blur-2xl" />

      <div className="relative z-10 w-full max-w-md">
        <LoginCard />
      </div>
    </div>
  )
}
