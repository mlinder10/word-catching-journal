import PageTitle from "@/components/custom-ui/page-title";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/app/_authenticated/profile/edit")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <main className="flex flex-col h-full">
      <PageTitle title="Edit Profile" />
    </main>
  );
}
