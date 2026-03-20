import PageTitle from "@/components/custom-ui/page-title";
import ProfileHeaderSection from "@/components/pages/profile/profile-header";
import ProfilePostsGrid from "@/components/pages/profile/profile-posts-grid";
import { createFileRoute, redirect } from "@tanstack/react-router";
import { useRef } from "react";

export const Route = createFileRoute("/app/_authenticated/profile/$profileId")({
  beforeLoad: ({ context, params }) => {
    if (params.profileId === context.auth.user?.id) {
      throw redirect({ to: "/app/profile" });
    }
  },
  component: ProfilePage,
});

function ProfilePage() {
  const { profileId } = Route.useParams();
  const scrollRef = useRef<HTMLDivElement>(null);

  return (
    <main className="flex flex-col h-full">
      <PageTitle title="Profile" />
      <div className="flex flex-col flex-1 overflow-y-auto" ref={scrollRef}>
        <ProfileHeaderSection userId={profileId} />
        <ProfilePostsGrid userId={profileId} scrollRef={scrollRef} />
      </div>
    </main>
  );
}
