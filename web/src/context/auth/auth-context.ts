import type { APIError, Response } from "@/api";
import type { User } from "@/types";
import { createContext } from "react";

export type AuthContextType = {
  isAuthenticated: boolean;
  isLoading: boolean;
  user: User | null;
  login: (email: string, password: string) => Promise<APIError | undefined>;
  register: (
    email: string,
    username: string,
    displayname: string,
    password: string,
  ) => Promise<Response<{ id: string }>>;
  logout: () => Promise<void>;
  validateToken: () => Promise<APIError | undefined>;
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
