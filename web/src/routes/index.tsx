import { createFileRoute, Link, redirect } from "@tanstack/react-router";
import {
  BookOpen,
  CheckCircle2,
  Sparkles,
  Users,
  ArrowRight,
} from "lucide-react";
import ThemeToggle from "@/components/custom-ui/theme-toggle";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import logo from "../assets/icon.png";

export const Route = createFileRoute("/")({
  component: LandingPage,
  beforeLoad: async ({ context }) => {
    // If the user already has a valid session, send them straight to the app.
    if (context.auth.isAuthenticated) {
      throw redirect({ to: "/app" });
    }

    const err = await context.auth.validateToken();
    if (!err) {
      throw redirect({ to: "/app" });
    }
  },
});

function LandingPage() {
  return (
    <main className="flex flex-col min-h-screen">
      <header className="border-b w-full">
        <div className="flex justify-between items-center mx-auto px-4 py-3 max-w-6xl">
          <div className="flex items-center gap-3">
            <div className="flex justify-center items-center bg-primary/10 border border-primary/20 rounded-lg size-12">
              <img
                src={logo}
                alt="Word Catching Journal"
                className="w-8 aspect-square"
              />
              {/* <Sparkles className="size-5 text-primary" /> */}
            </div>
            <div className="leading-tight">
              <p className="font-semibold">Word Catching Journal</p>
              <p className="text-muted-foreground text-xs">
                Catch words. Keep momentum.
              </p>
            </div>
          </div>

          <div className="flex items-center gap-2">
            <ThemeToggle className="hidden sm:flex gap-4" />
            <Button variant="ghost" asChild>
              <Link to="/login">Log in</Link>
            </Button>
            <Button asChild>
              <Link to="/register">Get started</Link>
            </Button>
          </div>
        </div>
      </header>

      <PageContent />

      <footer className="py-6 border-t text-muted-foreground text-sm">
        <div className="flex sm:flex-row flex-col sm:justify-between sm:items-center gap-2 mx-auto px-4 max-w-6xl">
          <p>© {new Date().getFullYear()} Word Catching Journal</p>
          <div className="flex items-center gap-4">
            <Link to="/login" className="hover:text-foreground">
              Log in
            </Link>
            <Link to="/register" className="hover:text-foreground">
              Register
            </Link>
          </div>
        </div>
      </footer>
    </main>
  );
}

function PageContent() {
  return (
    <div className="flex-1">
      <section className="mx-auto px-4 py-12 sm:py-16 max-w-6xl">
        <div className="items-center gap-10 grid lg:grid-cols-2">
          <div className="space-y-6">
            <div className="inline-flex items-center gap-2 bg-card px-3 py-1 border rounded-full text-sm">
              <CheckCircle2 className="size-4 text-secondary" />
              <span className="text-muted-foreground">
                Define words you love and share posts with others.
              </span>
            </div>

            <h2 className="font-bold text-4xl sm:text-5xl leading-tight">
              Turn new words into{" "}
              <span className="text-primary">a personal journal</span>.
            </h2>

            <p className="text-muted-foreground text-lg leading-relaxed">
              Word Catching Journal helps you capture definitions, build
              streaks, and revisit what you learned—one word at a time.
            </p>

            <div className="flex sm:flex-row flex-col sm:items-center gap-3">
              <Button asChild size="lg">
                <Link to="/register">
                  Start your journal <ArrowRight className="ml-2 size-4" />
                </Link>
              </Button>
              <Button asChild size="lg" variant="outline">
                <Link to="/login">I already have an account</Link>
              </Button>
            </div>

            <ul className="gap-3 grid sm:grid-cols-2 pt-2">
              <li className="flex items-start gap-2">
                <CheckCircle2 className="mt-1 size-4 text-secondary" />
                <p>Capture definitions as you read.</p>
              </li>
              <li className="flex items-start gap-2">
                <CheckCircle2 className="mt-1 size-4 text-secondary" />
                <p>Post and follow writers you like.</p>
              </li>
            </ul>
          </div>

          <Card className="relative overflow-hidden">
            <CardHeader>
              <CardTitle>How it works</CardTitle>
              <CardDescription>
                Everything starts with a single captured word.
              </CardDescription>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="flex gap-3">
                <div className="flex justify-center items-center bg-primary/10 border border-primary/20 rounded-lg size-10">
                  <BookOpen className="size-5 text-primary" />
                </div>
                <div>
                  <p className="font-semibold">1. Define</p>
                  <p className="text-muted-foreground text-sm">
                    Add the word and your definition.
                  </p>
                </div>
              </div>
              <div className="flex gap-3">
                <div className="flex justify-center items-center bg-secondary/10 border border-secondary/20 rounded-lg size-10">
                  <Sparkles className="size-5 text-secondary" />
                </div>
                <div>
                  <p className="font-semibold">2. Catch a moment</p>
                  <p className="text-muted-foreground text-sm">
                    Turn it into a post when you&apos;re ready.
                  </p>
                </div>
              </div>
              <div className="flex gap-3">
                <div className="flex justify-center items-center bg-primary/10 border border-primary/20 rounded-lg size-10">
                  <Users className="size-5 text-primary" />
                </div>
                <div>
                  <p className="font-semibold">3. Build momentum</p>
                  <p className="text-muted-foreground text-sm">
                    Follow others and revisit what you learn.
                  </p>
                </div>
              </div>
            </CardContent>
            <div className="-top-10 -right-10 absolute bg-primary/15 blur-2xl rounded-full size-40 pointer-events-none" />
            <div className="-bottom-10 -left-10 absolute bg-secondary/10 blur-2xl rounded-full size-40 pointer-events-none" />
          </Card>
        </div>
      </section>

      <section className="mx-auto px-4 pb-14 max-w-6xl">
        <div className="gap-4 grid md:grid-cols-3">
          <Card>
            <CardHeader>
              <CardTitle>Capture definitions</CardTitle>
              <CardDescription>
                Store words with context so they stick.
              </CardDescription>
            </CardHeader>
            <CardContent className="text-muted-foreground">
              Keep notes organized and easy to find later.
            </CardContent>
          </Card>
          <Card>
            <CardHeader>
              <CardTitle>Post & share</CardTitle>
              <CardDescription>
                Share your learning with your feed.
              </CardDescription>
            </CardHeader>
            <CardContent className="text-muted-foreground">
              Turn definitions into short posts that others can follow.
            </CardContent>
          </Card>
          <Card>
            <CardHeader>
              <CardTitle>Follow writers</CardTitle>
              <CardDescription>
                Stay inspired by what others catch.
              </CardDescription>
            </CardHeader>
            <CardContent className="text-muted-foreground">
              Build a feed of definitions and new perspectives.
            </CardContent>
          </Card>
        </div>
      </section>
    </div>
  );
}
