import { createContext, type Dispatch, type SetStateAction } from "react";

export const THEMES = ["light", "dark"] as const;
export type Theme = (typeof THEMES)[number];

type ThemeContextType = {
  theme: Theme;
  setTheme: Dispatch<SetStateAction<Theme>>;
};

export const ThemeContext = createContext<ThemeContextType>({
  theme: "light",
  setTheme: () => {},
});
