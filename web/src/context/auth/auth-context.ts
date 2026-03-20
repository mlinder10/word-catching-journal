import type { Response } from "@/api";
import type { User } from "@/types";
import { createContext } from "react";

export type AuthContextType = {
  isAuthenticated: boolean;
  isLoading: boolean;
  user: User | null;
  login: (email: string, password: string) => Promise<string | undefined>;
  register: (
    email: string,
    username: string,
    displayname: string,
    password: string,
  ) => Promise<Response<{ id: string }>>;
  logout: () => Promise<void>;
  validateToken: () => Promise<string | undefined>;
};

export const AuthContext = createContext<AuthContextType>({
  isAuthenticated: false,
  isLoading: false,
  user: null,
  login: async () => undefined,
  register: async () => ({ success: false, error: "internal server error" }),
  logout: async () => undefined,
  validateToken: async () => undefined,
});
