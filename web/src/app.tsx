import { useAuth } from "./context/auth";
import { RouterProvider } from "@tanstack/react-router";
import { router } from "./router";

export default function App() {
  const auth = useAuth();
  return <RouterProvider router={router} context={{ auth }} />;
}
