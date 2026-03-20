import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import App from "./app";
import "./index.css";
import { AuthProvider } from "./context/auth";
import { ThemeProvider } from "./context/theme";
import { Toaster } from "sonner";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <AuthProvider>
      <ThemeProvider>
        <App />
        <Toaster />
      </ThemeProvider>
    </AuthProvider>
  </StrictMode>,
);
