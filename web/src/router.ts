import { createRouter } from "@tanstack/react-router";
import { routeTree } from "./routeTree.gen";

export const router = createRouter({
  routeTree,
  context: {
    auth: {
      user: null,
      isAuthenticated: false,
      isLoading: false,
      login: async () => undefined,
      register: async () => undefined,
      logout: async () => undefined,
      validateToken: async () => undefined,
    },
  },
});

declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}
