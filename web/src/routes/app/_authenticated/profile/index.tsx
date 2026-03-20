import PageTitle from "@/components/custom-ui/page-title";
import ProfileHeaderSection from "@/components/pages/profile/profile-header";
import ProfilePostsGrid from "@/components/pages/profile/profile-posts-grid";
import { Button } from "@/components/ui/button";
import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";
import { useAuth } from "@/context/auth";
import { useUser } from "@/context/user";
import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { Menu } from "lucide-react";
import { useRef } from "react";

export const Route = createFileRoute("/app/_authenticated/profile/")({
  component: ProfilePage,
});

function ProfilePage() {
  const { id } = useUser();
  const scrollRef = useRef<HTMLDivElement>(null);

  return (
    <main className="flex flex-col h-full">
      <PageTitle title="Profile" ActionComponent={<MenuSheet />} />
      <div className="flex flex-col flex-1 overflow-y-auto" ref={scrollRef}>
        <ProfileHeaderSection userId={id} />
        <ProfilePostsGrid userId={id} scrollRef={scrollRef} />
      </div>
    </main>
  );
}

function MenuSheet() {
  const navigate = useNavigate();
  const { logout } = useAuth();

  async function handleLogout() {
    await logout();
    navigate({ to: "/login" });
  }

  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button variant="outline">
          <Menu />
        </Button>
      </SheetTrigger>
      <SheetContent>
        <SheetHeader>
          {/* TODO: rename */}
          <SheetTitle>Profile</SheetTitle>
        </SheetHeader>
        <div className="px-4">
          <Button className="mt-auto w-full" onClick={handleLogout}>
            Logout
          </Button>
        </div>
      </SheetContent>
    </Sheet>
  );
}
