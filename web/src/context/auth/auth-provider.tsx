import api from "@/api";
import type { User } from "@/types";
import { type ReactNode, useState } from "react";
import { AuthContext } from "./auth-context";

type AuthProviderProps = {
  children: ReactNode;
};

export function AuthProvider({ children }: AuthProviderProps) {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  async function login(email: string, password: string) {
    const res = await api.login({ email, password });
    if (res.success) setUser(res);
    else return res.error;
  }

  async function register(
    email: string,
    username: string,
    displayname: string,
    password: string,
  ) {
    return await api.register({
      email,
      username,
      displayname,
      password,
    });
  }

  async function logout() {
    await api.logout();
    setUser(null);
  }

  async function validateToken() {
    setIsLoading(true);
    const res = await api.validateToken();
    setIsLoading(false);
    if (res.success) setUser(res);
    else return res.error;
  }

  return (
    <AuthContext.Provider
      value={{
        isAuthenticated: !!user,
        isLoading,
        user,
        login,
        register,
        logout,
        validateToken,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}
