import { useTheme } from "@/context/theme";
import { cn } from "@/lib/utils";
import { MoonStar, SunMedium } from "lucide-react";

const HEIGHT = "h-9";

type ThemeToggleProps = {
  className?: string;
};

export default function ThemeToggle({ className }: ThemeToggleProps) {
  const { theme, setTheme } = useTheme();

  function toggleTheme() {
    setTheme((prev) => (prev === "light" ? "dark" : "light"));
  }

  return (
    <div className={cn("flex justify-between items-center px-1", className)}>
      <div>
        <p className="font-semibold text-muted-foreground text-xs">Theme</p>
        <p className="font-semibold">{theme === "light" ? "Light" : "Dark"}</p>
      </div>
      <div
        className={cn(
          "relative bg-secondary border-2 rounded-full w-fit aspect-2/1 transition-colors duration-500 cursor-pointer",
          "flex items-center",
          HEIGHT,
          theme === "light" ? "bg-secondary" : "bg-primary",
        )}
        onClick={toggleTheme}
      >
        <div className="flex flex-1 justify-center">
          <MoonStar className={cn("text-muted-foreground")} />
        </div>
        <div className="flex flex-1 justify-center">
          <SunMedium className={cn("text-muted-foreground")} />
        </div>
        <div
          className={cn(
            "top-0 absolute bg-muted rounded-full h-full aspect-square transition-all duration-500 ease-in-out",
            theme === "light" ? "left-0" : "left-[53%]",
          )}
        />
      </div>
    </div>
  );
}
