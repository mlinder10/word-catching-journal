import type { User } from "@/types";
import { createContext } from "react";

export const UserContext = createContext<User>({} as User);
