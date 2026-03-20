import Navbar from "@/components/navbar";
import { useAuth } from "@/context/auth";
import { UserProvider } from "@/context/user";
import { createFileRoute, Outlet, redirect } from "@tanstack/react-router";

export const Route = createFileRoute("/app/_authenticated")({
  beforeLoad: async ({ context, location }) => {
    if (!context.auth.isAuthenticated) {
      const err = await context.auth.validateToken();
      if (err) {
        throw redirect({
          to: "/login",
          search: {
            redirect: location.href,
          },
        });
      }
    }
  },
  component: AuthenticatedLayout,
});

function AuthenticatedLayout() {
  const { user } = useAuth();
  if (!user) return null; // should never happen

  return (
    <UserProvider user={user}>
      <div className="flex h-screen">
        <Navbar />
        <div className="flex-1 h-full">
          <Outlet />
        </div>
      </div>
    </UserProvider>
  );
}
