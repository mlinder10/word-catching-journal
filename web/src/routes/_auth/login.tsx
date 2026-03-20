import LoadSwap from "@/components/custom-ui/load-swap";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useAuth } from "@/context/auth";
import { createFileRoute, Link } from "@tanstack/react-router";
import { useRef, useState } from "react";
import { z } from "zod";

export const Route = createFileRoute("/_auth/login")({
  validateSearch: z.object({
    redirect: z.string().optional(),
  }),
  component: LoginPage,
});

function LoginPage() {
  const emailRef = useRef<HTMLInputElement>(null);
  const passwordRef = useRef<HTMLInputElement>(null);
  const { isPending, error, handleLogin } = useLogin();

  return (
    <main className="flex-1 place-items-center grid w-full">
      <Card className="min-w-md">
        <CardHeader>
          <CardTitle>Login</CardTitle>
          <CardDescription>
            Welcome back to your Word Catching Journal!
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="email">Email</Label>
            <Input id="email" type="email" ref={emailRef} />
          </div>
          <div className="space-y-2">
            <div className="flex justify-between">
              <Label htmlFor="password">Password</Label>
              <Link to="/reset-password">Forgot Password?</Link>
            </div>
            <Input id="password" type="password" ref={passwordRef} />
          </div>
          {error && <p className="text-destructive">{error}</p>}
          <Button
            onClick={() =>
              handleLogin(emailRef.current?.value, passwordRef.current?.value)
            }
            className="w-full"
            disabled={isPending}
          >
            <LoadSwap isLoading={isPending} label="Logging in...">
              Login
            </LoadSwap>
          </Button>
          <div className="space-x-1">
            <span className="text-muted-foreground">
              Don&apos;t have an account?
            </span>
            <Link to="/register">Register</Link>
          </div>
        </CardContent>
      </Card>
    </main>
  );
}

function useLogin() {
  const { redirect: redirectUrl } = Route.useSearch();
  const navigate = Route.useNavigate();
  const { login } = useAuth();

  const [isPending, setIsPending] = useState(false);
  const [error, setError] = useState<string | null>(null);

  async function handleLogin(
    email: string | undefined,
    password: string | undefined,
  ) {
    if (!email || !password) {
      setError("Email and password are required");
      return;
    }
    setIsPending(true);
    const error = await login(email, password);
    setIsPending(false);
    if (error === undefined) {
      navigate({ to: redirectUrl ?? "/app" });
    } else {
      setError(error);
    }
  }

  return { isPending, error, handleLogin };
}
