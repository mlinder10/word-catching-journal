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

export const Route = createFileRoute("/_auth/reset-password/")({
  component: ResetPasswordPage,
});

function ResetPasswordPage() {
  const emailRef = useRef<HTMLInputElement>(null);
  const [isPending, setIsPending] = useState(false);
  const [error, setError] = useState<string | null>(null);

  async function sendCode() {
    if (!emailRef.current) {
      setError("Email is required");
      return;
    }
    setError(null);
    setIsPending(true);
    const res = await api.requestResetPassword(emailRef.current.value);
    setIsPending(false);
    if (res.success) {
      toast.success("Email sent", { richColors: true });
    } else {
      setError(res.error);
    }
  }

  return (
    <main className="place-items-center grid h-screen">
      <Card className="min-w-md">
        <CardHeader>
          <CardTitle>Reset Password</CardTitle>
          <CardDescription>
            A link will be sent to your email to reset your password if you have
            an account
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="email">Email</Label>
            <Input id="email" type="email" ref={emailRef} />
            {error && <p className="text-destructive">{error}</p>}
          </div>
          <Button onClick={sendCode} disabled={isPending} className="w-full">
            Send Link
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
