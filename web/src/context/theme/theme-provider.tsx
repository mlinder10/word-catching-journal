import useLocalStorage from "@/hooks/use-local-storage";
import { useEffect, type ReactNode } from "react";
import { ThemeContext, type Theme } from "./theme-context";

type ThemeProviderProps = {
  children: ReactNode;
};

export function ThemeProvider({ children }: ThemeProviderProps) {
  const [theme, setTheme] = useLocalStorage<Theme>("theme", "light");

  useEffect(() => {
    const html = document.getElementsByTagName("html")[0];
    switch (theme) {
      case "light":
        html.classList.remove("dark");
        html.classList.add("light");
        break;
      case "dark":
        html.classList.remove("light");
        html.classList.add("dark");
        break;
    }
  }, [theme]);

  return (
    <ThemeContext.Provider value={{ theme, setTheme }}>
      {children}
    </ThemeContext.Provider>
  );
}
