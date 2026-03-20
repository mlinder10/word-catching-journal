import api from "@/api";
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
import { Separator } from "@/components/ui/separator";
import { createFileRoute, Link } from "@tanstack/react-router";
import { useRef, useState } from "react";
import { toast } from "sonner";

export const Route = createFileRoute("/_auth/reset-password/$code")({
  component: ResetPasswordCodePage,
});

function ResetPasswordCodePage() {
  const { code } = Route.useParams();
  const passwordRef = useRef<HTMLInputElement>(null);
  const [isPending, setIsPending] = useState(false);
  const [error, setError] = useState<string | null>(null);

  async function resetPassword() {
    if (!passwordRef.current) {
      setError("Password is required");
      return;
    }
    setError(null);
    setIsPending(true);
    const res = await api.resetPassword(code, passwordRef.current.value);
    setIsPending(false);
    if (res.success) {
      toast.success("Password reset", { richColors: true });
    } else {
      setError(res.error);
    }
  }

  return (
    <main className="place-items-center grid h-screen">
      <Card className="min-w-md">
        <CardHeader>
          <CardTitle>Reset Password</CardTitle>
          <CardDescription>Enter your new password</CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="password">Password</Label>
            <Input id="password" type="password" ref={passwordRef} />
            {error && <p className="text-destructive">{error}</p>}
          </div>
          <Button
            onClick={resetPassword}
            disabled={isPending}
            className="w-full"
          >
            Reset Password
          </Button>
          <Separator />
          <Button asChild variant="outline" className="w-full">
            <Link to="/login">Login</Link>
          </Button>
        </CardContent>
      </Card>
    </main>
  );
}
