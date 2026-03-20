import type { AuthContextType } from "@/context/auth";
import { queryClient } from "@/lib/query";
import { QueryClientProvider } from "@tanstack/react-query";
import { createRootRouteWithContext, Outlet } from "@tanstack/react-router";
// import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
// import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";

type RootRouteContext = {
  auth: AuthContextType;
};

export const Route = createRootRouteWithContext<RootRouteContext>()({
  component: RootLayout,
});

function RootLayout() {
  return (
    <QueryClientProvider client={queryClient}>
      <Outlet />
      {/* <TanStackRouterDevtools />
      <ReactQueryDevtools /> */}
    </QueryClientProvider>
  );
}
