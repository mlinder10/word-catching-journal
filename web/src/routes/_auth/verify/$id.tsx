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
import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { useEffect, useRef, useState } from "react";

export const Route = createFileRoute("/_auth/verify/$id")({
  component: VerifyEmailPage,
});

const DELAY_IN_SECONDS = 60;

function VerifyEmailPage() {
  const { id } = Route.useParams();
  const navigate = useNavigate();
  const codeRef = useRef<HTMLInputElement>(null);
  const [error, setError] = useState<string | null>(null);
  const [wait, setWait] = useState(DELAY_IN_SECONDS);

  async function handleVerify() {
    if (!codeRef.current?.value) {
      setError("Code is required");
      return;
    }

    const res = await api.verifyEmail(codeRef.current.value);
    if (res.success) {
      navigate({ to: "/app" });
    } else {
      setError("Invalid code");
    }
  }

  async function handleResend() {
    try {
      await api.resendVerification(id);
    } catch {
      setError("Failed to resend code");
    }
  }

  useEffect(() => {
    setWait(DELAY_IN_SECONDS);
    const interval = setInterval(() => {
      setWait((prev) => prev - 1);
    }, 1000);
    return () => clearInterval(interval);
  }, []);

  return (
    <main className="place-items-center grid h-screen">
      <Card className="min-w-md">
        <CardHeader>
          <CardTitle>Verify Your Email</CardTitle>
          <CardDescription>
            A 6 digit code has been sent to your email. Please check your spam
            folder if you don&apos;t see it.
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="code">Code</Label>
            <div className="flex gap-2">
              <Input id="code" type="text" ref={codeRef} />
              <Button onClick={handleVerify}>Verify</Button>
            </div>
            {error && <p className="text-destructive">{error}</p>}
          </div>
          <Separator />
          <Button
            variant="outline"
            className="w-full"
            onClick={handleResend}
            disabled={wait > 0}
          >
            {wait > 0 ? `Resend in ${wait}s` : "Resend"}
          </Button>
        </CardContent>
      </Card>
    </main>
  );
}
