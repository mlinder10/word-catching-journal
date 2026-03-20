import type { User } from "@/types";
import { UserContext } from "./user-context";

type UserProviderProps = {
  children: React.ReactNode;
  user: User;
};

export function UserProvider({ children, user }: UserProviderProps) {
  return <UserContext.Provider value={user}>{children}</UserContext.Provider>;
}
