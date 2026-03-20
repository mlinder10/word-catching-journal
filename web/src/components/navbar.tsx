import { Link, useLocation } from "@tanstack/react-router";
import { Button } from "./ui/button";
import { Calendar, ChevronLeft, Plus, Search } from "lucide-react";
import { useState, type ComponentType } from "react";
import { cn } from "@/lib/utils";
import { useUser } from "@/context/user";
import ThemeToggle from "./custom-ui/theme-toggle";
import ProfilePicture from "./custom-ui/profile-picture";

export default function Navbar() {
  const user = useUser();
  const [isOpen, setIsOpen] = useState(true);

  return (
    <div
      className={cn(
        "flex flex-col items-center gap-4 p-4 border-r",
        isOpen && "sm:w-72 sm:items-start",
      )}
    >
      <div className="hidden sm:flex justify-between items-center sm:w-full">
        {isOpen && <h1>Word Catching Journal</h1>}
        <Button variant="ghost" onClick={() => setIsOpen((prev) => !prev)}>
          <ChevronLeft className={cn(!isOpen && "rotate-180")} />
        </Button>
      </div>
      <ul className={cn("flex flex-col w-full", !isOpen && "items-center")}>
        <NavLink to="/app" title="Feed" isOpen={isOpen} Icon={Calendar} />
        <NavLink
          to="/app/define"
          title="Define"
          isOpen={isOpen}
          Icon={Search}
        />
        <NavLink to="/app/post" title="Post" isOpen={isOpen} Icon={Plus} />
      </ul>

      <div className="flex flex-col gap-4 mt-auto w-full">
        {isOpen && <ThemeToggle className="hidden sm:flex" />}
        <Button variant="outline" asChild>
          <Link to="/app/profile">
            <ProfilePicture
              username={user.username}
              imageUrl={user.imageUrl}
              color={user.color}
              size={16}
            />
            {isOpen && (
              <span className="hidden sm:inline">{user.username}</span>
            )}
          </Link>
        </Button>
      </div>
    </div>
  );
}

type NavLinkProps = {
  to: string;
  title: string;
  isOpen: boolean;
  Icon: ComponentType;
};

function NavLink({ to, title, isOpen, Icon }: NavLinkProps) {
  const location = useLocation();
  const isActive = location.pathname === to;

  return (
    <li>
      <Button
        asChild
        variant="ghost"
        className={cn(
          "justify-start w-full text-muted-foreground",
          isActive && "text-foreground font-semibold",
        )}
      >
        <Link to={to}>
          <Icon />
          {isOpen && <span className="hidden sm:inline">{title}</span>}
        </Link>
      </Button>
    </li>
  );
}
