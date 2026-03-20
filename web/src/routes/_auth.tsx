import { createFileRoute, Outlet, redirect } from "@tanstack/react-router";
import icon from "@/assets/icon.png";

export const Route = createFileRoute("/_auth")({
  beforeLoad: async ({ context }) => {
    if (context.auth.isAuthenticated) {
      throw redirect({ to: "/app" });
    }
  },
  component: AuthLayout,
});

function AuthLayout() {
  return (
    <div className="flex flex-col h-screen">
      <header className="p-8">
        <div className="flex items-center gap-2">
          <img src={icon} className="w-8 aspect-square" />
          <h1 className="font-semibold text-2xl">Word Catching Journal</h1>
        </div>
      </header>
      <Outlet />
    </div>
  );
}
