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

export const Route = createFileRoute("/_auth/register")({
  component: RegisterPage,
});

function RegisterPage() {
  const emailRef = useRef<HTMLInputElement>(null);
  const usernameRef = useRef<HTMLInputElement>(null);
  const displaynameRef = useRef<HTMLInputElement>(null);
  const passwordRef = useRef<HTMLInputElement>(null);
  const { isPending, error, handleRegister } = useRegister();

  return (
    <main className="place-items-center grid h-screen">
      <Card className="min-w-md">
        <CardHeader>
          <CardTitle>Register</CardTitle>
          <CardDescription>
            Welcome to the Word Catching Journal
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="email">Email</Label>
            <Input id="email" type="email" ref={emailRef} />
          </div>
          <div className="space-y-2">
            <Label htmlFor="username">Username</Label>
            <Input id="username" type="text" ref={usernameRef} />
          </div>
          <div className="space-y-2">
            <Label htmlFor="displayname">Display Name</Label>
            <Input id="displayname" type="text" ref={displaynameRef} />
          </div>
          <div className="space-y-2">
            <Label htmlFor="password">Password</Label>
            <Input id="password" type="password" ref={passwordRef} />
          </div>
          {error && <p className="text-destructive">{error}</p>}
          <Button
            onClick={() =>
              handleRegister(
                emailRef.current?.value,
                usernameRef.current?.value,
                displaynameRef.current?.value,
                passwordRef.current?.value,
              )
            }
            className="w-full"
            disabled={isPending}
          >
            <LoadSwap isLoading={isPending} label="Registering...">
              Register
            </LoadSwap>
          </Button>
          <div className="space-x-1">
            <span className="text-muted-foreground">
              Already have an account?
            </span>
            <Link to="/login">Login</Link>
          </div>
        </CardContent>
      </Card>
    </main>
  );
}

function useRegister() {
  const navigate = Route.useNavigate();
  const { register } = useAuth();

  const [isPending, setIsPending] = useState(false);
  const [error, setError] = useState<string | null>(null);

  async function handleRegister(
    email: string | undefined,
    username: string | undefined,
    displayname: string | undefined,
    password: string | undefined,
  ) {
    if (!email || !username || !displayname || !password) {
      setError("All fields are required");
      return;
    }
    setIsPending(true);
    const res = await register(email, username, displayname, password);
    setIsPending(false);
    if (res.success) {
      navigate({ to: "/verify/$id", params: { id: res.id } });
    } else {
      setError(res.error);
    }
  }

  return { isPending, error, handleRegister };
}
